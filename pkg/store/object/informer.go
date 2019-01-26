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
	"sync"

	"github.com/galexrt/edenconfmgmt/pkg/store/cache"
	"github.com/galexrt/edenconfmgmt/pkg/store/data"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

// InformerResult result returned by the channel.
type InformerResult struct {
	Closed  bool
	Errors  []error
	State   cache.ResultState
	Key     string
	Value   interface{}
	Version int64
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

// Informer keeps track of watches and the receivers of those watches.
type Informer struct {
	channel    map[string]*chanContainer
	receivers  map[string]*receiverList
	dataStore  data.Store
	cacheStore data.Store
	wg         sync.WaitGroup
	logger     *zap.Logger
}

// NewInformer returns a new Informer.
func NewInformer(dataStore data.Store, cacheStore data.Store, logger *zap.Logger) *Informer {
	return &Informer{
		channel:    map[string]*chanContainer{},
		receivers:  map[string]*receiverList{},
		dataStore:  dataStore,
		cacheStore: cacheStore,
		logger:     logger.Named("pkg/store/object:Informer"),
	}
}

// TODO "Copy" pkg/store/cache" functions
