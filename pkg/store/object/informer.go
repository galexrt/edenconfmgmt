/*
Copyright 2019 Alexander Trost. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package object

import (
	"context"
	"fmt"
	"path"
	"sort"
	"strings"
	"sync"

	core_v1 "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	"github.com/galexrt/edenconfmgmt/pkg/store/cache"
	storecommon "github.com/galexrt/edenconfmgmt/pkg/store/common"
	"github.com/galexrt/edenconfmgmt/pkg/utils/errors"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"
)

type chanContainer struct {
	watch chan *storecommon.InformerResult
	link  *chanContainer
}

type receiverList struct {
	list       []chan *storecommon.InformerResult
	usageCount uint64
	sync.RWMutex
}

type labelReceivers struct {
	labels    map[string]*labelReceivers
	receivers *receiverList
}

// Informer keeps track of watches and the receivers of those watches.
type Informer struct {
	channel         map[string]*chanContainer
	receiversName   map[string]*receiverList
	receiversLabels map[string]*labelReceivers
	dataStore       *cache.Store
	wg              sync.WaitGroup
	logger          *zap.Logger
	globalContext   context.Context
}

// NewInformer returns a new Informer.
func NewInformer(ctx context.Context, dataStore *cache.Store, logger *zap.Logger) *Informer {
	return &Informer{
		channel:         map[string]*chanContainer{},
		receiversName:   map[string]*receiverList{},
		receiversLabels: map[string]*labelReceivers{},
		dataStore:       dataStore,
		logger:          logger.Named("pkg/store/object:Informer"),
		globalContext:   ctx,
	}
}

// Start start cleanup/compaction logics.
func (inf *Informer) Start(stopCh chan struct{}) error {
	select {
	case <-stopCh:
		return nil
	}
}

// Register register a watch path.
func (inf *Informer) Register(apiBasePath string) error {
	watchKey := path.Join("/", apiBasePath, "/")
	dataStoreCh, err := inf.dataStore.WatchRecursively(inf.globalContext, watchKey)
	if err != nil {
		return err
	}
	if _, ok := inf.channel[watchKey]; !ok {
		if dataStoreCh == nil {
			return fmt.Errorf("no dataStoreCh given and none in the dataStoreCh list")
		}
		inf.channel[watchKey] = &chanContainer{
			watch: dataStoreCh,
		}
		go inf.watch(inf.globalContext, watchKey)
	}
	if _, ok := inf.receiversLabels[watchKey]; !ok {
		inf.receiversLabels[watchKey] = &labelReceivers{
			labels: map[string]*labelReceivers{},
			receivers: &receiverList{
				list: []chan *storecommon.InformerResult{},
			},
		}
	}
	return nil
}

// DataStoreChExists return if a dataStoreCh exists in the system for the given key.
// (Can also mean that there is a "higher" level dataStoreCh in the list of `inf.channel`)
func (inf *Informer) DataStoreChExists(key string) bool {
	if st := inf.getDataStoreCh(path.Join("/", key)); st != nil {
		return true
	}
	return false
}

// getDataStoreCh
// Example `inf.channel` keys:
// * /registry/myapi/
// * /registry/otherapi/
// When asked for `key: /registry/otherapi/my-object`, the second key's
// value would get returned.
func (inf *Informer) getDataStoreCh(key string) *chanContainer {
	currentPath := "/"
	components := strings.Split(key, "/")
	for _, c := range components {
		currentPath = path.Join(currentPath, c)
		if _, ok := inf.channel[currentPath]; ok {
			return inf.channel[currentPath]
		}
	}
	return nil
}

// getReceiversForName
// Example `inf.receivers` keys:
// * /registry/myapi/
// * /registry/otherapi/
// When asked for `key: /registry/otherapi/my-object`, the second key's
// value would get returned.
func (inf *Informer) getReceiversForName(key string) []*receiverList {
	receivers := []*receiverList{}

	for recvPath := range inf.receiversName {
		if recvPath == key {
			receivers = append(receivers, inf.receiversName[recvPath])
		} else {
			components := strings.Split(key, "/")
			for k := len(components); k > 0; k-- {
				currentPath := strings.Join(components[0:k], "/")
				if recvPath == currentPath || strings.TrimRight(recvPath, "/") == currentPath {
					receivers = append(receivers, inf.receiversName[recvPath])
					break
				}
			}
		}
	}

	return receivers
}

// getReceiversForLabels
func (inf *Informer) getReceiversForLabels(labels map[string]string) []*receiverList {
	receivers := []*receiverList{}

	var keys []string
	for k := range labels {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// TODO Build api name
	apiName := "/nodes/v1alpha"
	recvLabelList, ok := inf.receiversLabels[apiName]
	if !ok {
		return nil
	}

	for _, k := range keys {
		label := strings.Join([]string{k, labels[k]}, "=")
		if _, ok = recvLabelList.labels[label]; !ok {
			return nil
		}
		recvLabelList = recvLabelList.labels[label]
		receivers = append(receivers, recvLabelList.receivers)
	}

	return receivers
}

// WatchOnName return a Watch for a given path (`key`).
func (inf *Informer) WatchOnName(ctx context.Context, key string) (chan *storecommon.InformerResult, error) {
	normKey := path.Join("/", key)
	if _, ok := inf.receiversName[normKey]; !ok {
		inf.receiversName[normKey] = &receiverList{
			list: []chan *storecommon.InformerResult{},
		}
	}
	ch := make(chan *storecommon.InformerResult)
	inf.receiversName[normKey].Lock()
	inf.receiversName[normKey].list = append(inf.receiversName[normKey].list, ch)
	inf.receiversName[normKey].Unlock()
	return ch, nil
}

// WatchOnLabels return a watch for a given list of labels.
func (inf *Informer) WatchOnLabels(ctx context.Context, labels []string) (chan *storecommon.InformerResult, error) {
	if len(labels) == 0 {
		return nil, ErrNoLabelsInformer
	}
	ch := make(chan *storecommon.InformerResult)

	// TODO Build api name
	apiName := "/nodes/v1alpha"
	recvLabelList, ok := inf.receiversLabels[apiName]
	if !ok {
		return nil, fmt.Errorf("something went wrong, no api found in our receivers list")
	}

	resultLabelList := recvLabelList
	for _, label := range labels {
		if _, ok = resultLabelList.labels[label]; !ok {
			resultLabelList.labels[label] = &labelReceivers{
				labels: map[string]*labelReceivers{},
				receivers: &receiverList{
					list: []chan *storecommon.InformerResult{},
				},
			}
		}
		resultLabelList = resultLabelList.labels[label]
	}
	resultLabelList.receivers.list = append(resultLabelList.receivers.list, ch)

	return ch, nil
}

// watch
func (inf *Informer) watch(ctx context.Context, key string) {
	dataStoreCh := inf.getDataStoreCh(key)
	for {
		select {
		case <-ctx.Done():
			return
		case resp := <-dataStoreCh.watch:
			inf.wg.Add(1)
			go func(resp *storecommon.InformerResult) {
				defer inf.wg.Done()
				switch resp.State {
				case storecommon.ResultStateCreated:
					fallthrough
				case storecommon.ResultStateUpdated:
					err := inf.objectCreatedOrChanged(resp)
					if err != nil {
						inf.logger.Error("error in informer notification for created or updated event", zap.Error(err))
					}
				case storecommon.ResultStateDeleted:
					err := inf.objectDelete(resp)
					if err != nil {
						inf.logger.Error("error in informer notification for deletion event", zap.Error(err))
					}
				default:
					inf.logger.Warn("got dataStore event with no or unknown ResultState", zap.String("key", key), zap.Int64("keyVersion", resp.Version), zap.Any("resultState", resp.State))
				}
			}(resp)
			inf.wg.Wait()
			if resp.Closed {
				inf.logger.Warn("dataStore event has been canceled", zap.Error(errors.Concat(resp.Errors...)))
				return
			}
		}
	}
}

func (inf *Informer) sendInformerResult(result *storecommon.InformerResult, receivers []*receiverList) {
	for _, rcvs := range receivers {
		rcvs.RLock()
		for _, rcv := range rcvs.list {
			rcv <- result
		}
		rcvs.RUnlock()
	}
}

func (inf *Informer) objectCreatedOrChanged(result *storecommon.InformerResult) error {
	obj := &core_v1.GenericObject{}
	err := proto.Unmarshal(result.Value, obj)
	if err != nil {
		return err
	}

	if obj.GetMetadata() == nil {
		return nil
	}
	// TODO Get all receiivers for name and labels, then send InformerResult
	wg := sync.WaitGroup{}
	name := obj.GetMetadata().GetName()
	wg.Add(1)
	go func() {
		defer wg.Done()
		// TODO build path for name
		inf.sendInformerResult(result, inf.getReceiversForName(name))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		inf.sendInformerResult(result, inf.getReceiversForLabels(obj.GetMetadata().GetLabels()))
	}()

	wg.Wait()
	return nil
}

func (inf *Informer) objectDelete(result *storecommon.InformerResult) error {
	obj := &core_v1.GenericObject{}
	err := proto.Unmarshal(result.Value, obj)
	if err != nil {
		return err
	}
	// TODO Inform about an object deletion.

	return nil
}
