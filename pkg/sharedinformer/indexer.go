/*
Copyright 2018 Alexander Trost. All rights reserved.

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

package sharedinformer

import (
	"context"
	"path/filepath"
	"sync"
	"time"

	"github.com/galexrt/edenconfmgmt/pkg/common"
	"github.com/galexrt/edenconfmgmt/pkg/datastore"
	"github.com/galexrt/edenconfmgmt/pkg/datastore/handlers"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

var logger = common.GetLogger(nil)

// Indexer thing
type Indexer struct {
	options             *Options
	getWatchFunc        func(ctx context.Context) (clientv3.WatchChan, error)
	nameInformer        map[string]InformerMap
	nameInformerMutex   sync.Mutex
	labelsInformer      map[string]InformerMap
	labelsInformerMutex sync.Mutex
	wg                  sync.WaitGroup
}

// Options Indexer options
type Options struct {
	ParallelEventProcessingLimit int64
}

// Informer custom type for one-way channel returning Result.
type Informer struct {
	Watch  chan *Result
	timeID int64
	search string
}

// InformerMap kind of only here to save some characters.
type InformerMap map[int64]*Informer

// Result result returned by the channel
type Result struct {
	Closed bool
	Error  error
	State  datastore.ResponseState
	Key    string
	Name   string
	Object []byte
}

// New return a new Indexer
func New(getWatchFunc func(context.Context) (clientv3.WatchChan, error)) *Indexer {
	return &Indexer{
		options: &Options{
			ParallelEventProcessingLimit: 128,
		},
		getWatchFunc:   getWatchFunc,
		nameInformer:   map[string]InformerMap{},
		labelsInformer: map[string]InformerMap{},
		wg:             sync.WaitGroup{},
	}
}

// Run start the etcd watch informer loop
func (ind *Indexer) Run(ctx context.Context, stopCh chan struct{}) error {
	logger.Debug("indexer started")
	// TODO make this confiruable nad/or react better to none existing keys (e.g., `Created: false` loops forever...)
	failLimit := 10
	for failCount := 0; failCount < failLimit; failCount++ {
		select {
		case <-stopCh:
			return nil
		default:
		}
		// TODO Restart etcd watch by running getWatchFunc again
		watcher, err := ind.getWatchFunc(ctx)
		if err != nil {
			return err
		}
		logger.Debug("etcd watch: got a watcher from getWatchFunc()")
	watchLoop:
		for {
			var result clientv3.WatchResponse
			select {
			case <-stopCh:
				return nil
			case result = <-watcher:
			}
			logger.Debug("etcd watch: result", zap.Any("result", result))
			if result.Canceled {
				logger.Warn("etcd watch has been canceled", zap.Any("etcd-header", result.Header))
				// TODO Should we really restart the watch in this case?
				break watchLoop
			}
			if len(result.Events) == 0 {
				logger.Warn("etcd watch: no events in etcd watch response! WHY? We don't know.. Let's wait a second and continue")
				time.Sleep(1 * time.Second)
				continue
			}
			for _, event := range result.Events {
				ind.wg.Add(1)
				go func(ev *clientv3.Event) {
					defer ind.wg.Done()
					ind.processETCDEvent(ev)
				}(event)
			}
			ind.wg.Wait()
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func (ind *Indexer) processETCDEvent(event *clientv3.Event) {
	logger.Debug("etcd watch: result event", zap.Any("event", event))
	// {"event": {"kv":{"key":"L3Rlc3QxMjMvbm9wZS90ZXN0","create_revision":5,"mod_revision":42,"version":38,"value":"bm9wZQ=="}}}
	key := string(event.Kv.Key)
	name := filepath.Base(key)

	logger.Debug("event name and key", zap.String("name", name), zap.String("key", key))

	// As we only watch APIs here
	informerList, ok := ind.nameInformer[name]
	if !ok {
		logger.Debug("no informer for name", zap.String("name", name), zap.String("key", key))
		return
	}
	state := handlers.EtcdEventStateToHandlerState(event.Type, event.IsCreate(), event.IsModify())

	resp := &Result{
		Error:  nil,
		State:  state,
		Key:    key,
		Name:   name,
		Object: event.Kv.Value,
	}

	for _, informer := range informerList {
		informer.Watch <- resp
	}
}

// GetNameInformer return an Informer watching on specific named object containing a channel which one can "listen" to.
func (ind *Indexer) GetNameInformer(name string) *Informer {
	if _, ok := ind.nameInformer[name]; !ok {
		ind.nameInformer[name] = InformerMap{}
	}
	watcher := make(chan *Result)
	id := time.Now().UnixNano()
	ok := true
	for ok {
		if _, ok = ind.nameInformer[name][id]; ok {
			id = id + 1
		}
	}

	informer := &Informer{
		timeID: id,
		Watch:  watcher,
		search: name,
	}

	ind.nameInformerMutex.Lock()
	ind.nameInformer[name][id] = informer
	ind.nameInformerMutex.Unlock()

	return informer
}

// GetLabelsInformer return an Informer watching on specificly labelled object containing a channel which one can "listen" to.
func (ind *Indexer) GetLabelsInformer(labels string) *Informer {
	if _, ok := ind.labelsInformer[labels]; !ok {
		ind.labelsInformer[labels] = InformerMap{}
	}
	watcher := make(chan *Result)
	id := time.Now().UnixNano()
	ok := true
	for ok {
		if _, ok = ind.labelsInformer[labels][id]; ok {
			id = id + 1337
		}
	}

	informer := &Informer{
		timeID: id,
		Watch:  watcher,
		search: labels,
	}

	ind.labelsInformerMutex.Lock()
	ind.labelsInformer[labels][id] = informer
	ind.labelsInformerMutex.Unlock()

	return informer
}

// CloseLabelsInformer close a labelsInformer watch channel.
func (ind *Indexer) CloseLabelsInformer(informer *Informer) {
	if _, ok := ind.labelsInformer[informer.search]; ok {
		if _, ok := ind.labelsInformer[informer.search][informer.timeID]; ok {
			ind.labelsInformerMutex.Lock()
			delete(ind.labelsInformer[informer.search], informer.timeID)
			ind.labelsInformerMutex.Unlock()
		}
	}
}
