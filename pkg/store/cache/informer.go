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

package cache

import (
	"context"
	"fmt"
	"path"
	"strings"
	"sync"

	"github.com/galexrt/edenrun/pkg/store/data"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"go.uber.org/zap"
)

// Informer keeps track of watches and the receivers of those watches.
type Informer struct {
	channel    map[string]*chanContainer
	receivers  map[string]*receiverList
	dataStore  data.Store
	cacheStore data.Store
	wg         sync.WaitGroup
	logger     *zap.Logger
}

type chanContainer struct {
	watch clientv3.WatchChan
	link  *chanContainer
}

type receiverList struct {
	list       []chan *InformerResult
	usageCount uint64
	sync.RWMutex
}

// NewInformer returns a new Informer.
func NewInformer(dataStore data.Store, cacheStore data.Store, logger *zap.Logger) *Informer {
	return &Informer{
		channel:    map[string]*chanContainer{},
		receivers:  map[string]*receiverList{},
		dataStore:  dataStore,
		cacheStore: cacheStore,
		logger:     logger.Named("pkg/store/cache:Informer"),
	}
}

// Start start the watch compaction logic.
func (inf *Informer) Start(stopCh chan struct{}) error {
	var err error
	for {
		select {
		// TODO use `inf.dataStore` to get a more "compact" watch (e.g., replace two single key watches with one directory watch)
		case <-stopCh:
			return err
		}
	}
}

// DataStoreChExists return if a dataStoreCh exists in the system.
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

// getReceiverChs
// Example `inf.receivers` keys:
// * /registry/myapi/
// * /registry/otherapi/
// When asked for `key: /registry/otherapi/my-object`, the second key's
// value would get returned.
func (inf *Informer) getReceiverChs(key string) []*receiverList {
	receivers := []*receiverList{}

	for recvPath := range inf.receivers {
		if recvPath == key {
			receivers = append(receivers, inf.receivers[recvPath])
		} else {
			components := strings.Split(key, "/")
			for k := len(components); k > 0; k-- {
				currentPath := strings.Join(components[0:k], "/")
				if recvPath == currentPath || strings.TrimRight(recvPath, "/") == currentPath {
					receivers = append(receivers, inf.receivers[recvPath])
					break
				}
			}
		}
	}

	return receivers
}

// Watch return a Watch for a given path (`key`).
func (inf *Informer) Watch(ctx context.Context, key string, dataStoreCh clientv3.WatchChan) (chan *InformerResult, error) {
	normKey := path.Join("/", key)
	if _, ok := inf.channel[normKey]; !ok {
		if dataStoreCh == nil {
			return nil, fmt.Errorf("no dataStoreCh given and none in the dataStoreCh list ")
		}
		inf.channel[normKey] = &chanContainer{
			watch: dataStoreCh,
		}
		go inf.watch(ctx, normKey)
	}
	if _, ok := inf.receivers[normKey]; !ok {
		inf.receivers[normKey] = &receiverList{
			list: []chan *InformerResult{},
		}
	}
	ch := make(chan *InformerResult)
	inf.receivers[normKey].Lock()
	inf.receivers[normKey].list = append(inf.receivers[normKey].list, ch)
	inf.receivers[normKey].Unlock()
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
			var errs []error
			if err := resp.Err(); err != nil {
				errs = append(errs, err)
			}
			for _, event := range resp.Events {
				key := string(event.Kv.Key)
				state := convertETCDtoResultState(event)
				value := event.Kv.Value
				version := event.Kv.Version
				inf.wg.Add(1)
				go func(key string, state ResultState, value []byte) {
					defer inf.wg.Done()
					switch state {
					case ResultStateCreated:
						fallthrough
					case ResultStateUpdated:
						inf.cacheStore.Put(ctx, key, value)
					case ResultStateDeleted:
						inf.cacheStore.Delete(ctx, key)
					default:
						inf.logger.Warn("got dataStore event with no or unknown ResultState", zap.String("key", key), zap.Int64("keyVersion", version), zap.Any("resultState", state))
					}
				}(key, state, value)
				inf.wg.Add(1)
				go func() {
					defer inf.wg.Done()
					inf.sendInformerResult(errs, key, resp.Canceled, state, value, version)
				}()
			}
			if resp.Canceled {
				inf.logger.Warn("dataStore event has been canceled", zap.Error(resp.Err()))
				return
			}
		}
	}
}

func (inf *Informer) sendInformerResult(errs []error, key string, canceled bool, state ResultState, value []byte, version int64) {
	result := &InformerResult{
		Errors:  errs,
		Closed:  canceled,
		Key:     key,
		State:   state,
		Value:   value,
		Version: version,
	}
	// TODO should we lock the list here?
	receivers := inf.getReceiverChs(key)
	for _, rcvs := range receivers {
		rcvs.RLock()
		for _, rcv := range rcvs.list {
			rcv <- result
		}
		rcvs.RUnlock()
	}
}

func convertETCDtoResultState(event *clientv3.Event) ResultState {
	if event.Type == mvccpb.DELETE {
		return ResultStateDeleted
	}
	if event.Type == mvccpb.PUT {
		if event.Kv.CreateRevision == event.Kv.Version {
			return ResultStateUpdated
		}
		return ResultStateCreated
	}
	return ResultStateUnknown
}
