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
	"strings"
	"sync"

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

// Informer keeps track of watches and the receivers of those watches.
type Informer struct {
	channel         map[string]*chanContainer
	receiversName   map[string]*receiverList
	receiversLabels map[string]*receiverList
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
		receiversLabels: map[string]*receiverList{},
		dataStore:       dataStore,
		logger:          logger.Named("pkg/store/object:Informer"),
		globalContext:   ctx,
	}
}

// Start
func (inf *Informer) Start(stopCh chan struct{}) error {
	select {
	case <-stopCh:
		return nil
	}
}

// Register register a watch path.
func (inf *Informer) Register(apiBasePath string) error {
	watchKey := path.Join("/", apiBasePath)
	dataStoreCh, err := inf.dataStore.Watch(inf.globalContext, watchKey)
	if err != nil {
		return err
	}
	if _, ok := inf.channel[watchKey]; !ok {
		if dataStoreCh == nil {
			return fmt.Errorf("no dataStoreCh given and none in the dataStoreCh list ")
		}
		inf.channel[watchKey] = &chanContainer{
			watch: dataStoreCh,
		}
		go inf.watch(inf.globalContext, watchKey)
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

// getReceiverLabels
func (inf *Informer) getReceiverLabels(key string) []*receiverList {
	receivers := []*receiverList{}

	return receivers
}

// getReceiverNames
// Example `inf.receivers` keys:
// * /registry/myapi/
// * /registry/otherapi/
// When asked for `key: /registry/otherapi/my-object`, the second key's
// value would get returned.
func (inf *Informer) getReceiverNames(key string) []*receiverList {
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
	for _, label := range labels {
		if _, ok := inf.receiversLabels[label]; !ok {
			inf.receiversLabels[label] = &receiverList{
				list: []chan *storecommon.InformerResult{},
			}
		}
		inf.receiversLabels[label].Lock()
		inf.receiversLabels[label].list = append(inf.receiversLabels[label].list, ch)
		inf.receiversLabels[label].Unlock()
	}
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
					inf.objectCreatedOrChanged(resp)
				case storecommon.ResultStateDeleted:
					inf.objectDelete(resp)
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
	obj := &struct {
		proto.Message
		MetadataUnmarshalerAndMarshaler
	}{}
	err := proto.Unmarshal(result.Value, obj)
	if err != nil {
		return err
	}
	fmt.Printf("OBJECT CREATED OR CHANGED: %+v\n", obj)

	// TODO Get all receiivers for name and labels, then send InformerResult
	//inf.sendInformerResult(result, receivers)
	return nil
}

func (inf *Informer) objectDelete(result *storecommon.InformerResult) {
	// TODO Delete labels and so on for given object.

}
