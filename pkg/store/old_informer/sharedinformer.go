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

package informer

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

/* =============================================================================
TODO
* SharedInformer will create the `chan *Result`
* SharedInformer will "compact" watches on the same (sub-)path(s) automatically. Example:
	* **2** watches: 1. `/example/apiname/kindA123/objectabc`, 2. `/example/apiname/nopethisisatest/xyz123123`
	* => Results in: **ONE** recursive watch on `/example/apiname`.
	* Advantages: Less watches on the datastore.
	* Disadvantage: This will probably be a bottle neck at some point as incoming Results will
		for the first time be single threadely accepted (future: add (global) worker count?).
============================================================================= */

type SharedInformerClose func(path string)

// SharedInformer registration and "time" keeper over watches.
type SharedInformer struct {
	watches                 map[string]chan *Result
	watchUsageCounter       map[string]*uint64
	wg                      sync.WaitGroup
	StoreWatchCleanInterval time.Duration
}

// StartWithStopCh start SharedInformer and use context.Context to cancel it.
func (inf *SharedInformer) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			inf.close()
			return
			// TODO make configurable again
		case <-time.After(inf.StoreWatchCleanInterval):
			inf.cleanStoreWatches()
		}
	}
}

// StartWithStopCh start watch for SharedInformer and use stopCh to cancel it.
func (inf *SharedInformer) StartWithStopCh(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			inf.close()
			return
			// TODO make configurable again
		case <-time.After(7 * time.Second):
			inf.cleanStoreWatches()
		}
	}
}

// Watch returns a new watch channel.
func (inf *SharedInformer) Watch(input chan *Result, path string, middleware ...InformerMiddleware) *Informer {
	if _, ok := inf.watches[path]; !ok {
		inf.watches[path] = input
		// TODO Add new goroutine for this watch
		inf.addStoreWatch(path)
	}

	_ = atomic.AddUint64(inf.watchUsageCounter[path], 1)

	infType := InformerTypeObject
	if path[len(path)-1:] == "/" {
		infType = InformerTypeDirectory
	}

	return &Informer{
		input:       make(chan *Result),
		output:      make(chan *Result),
		Type:        infType,
		Path:        path,
		middleware:  middleware,
		sharedClose: inf.Close,
	}
}

func (inf *SharedInformer) addStoreWatch(path string) {
	// TODO
}

func (inf *SharedInformer) cleanStoreWatches() {
	// TODO remove store watches for watchUsageCounter == 0
}

func (inf *SharedInformer) close() {
	for _, ch := range inf.watches {
		close(ch)
	}
}

func (inf *SharedInformer) Close(path string) {
	atomic.AddUint64(inf.watchUsageCounter[path], ^uint64(0))
}
