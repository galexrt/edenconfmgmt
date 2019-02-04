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
	"sync"

	storecommon "github.com/galexrt/edenconfmgmt/pkg/store/common"
	"github.com/galexrt/edenconfmgmt/pkg/store/data"
	"github.com/galexrt/edenconfmgmt/pkg/utils/errors"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

// Store using the dataStore and a second cacheStore for caching.
type Store struct {
	dataStore  data.Store
	cacheStore data.Store
	informer   *Informer
	prefix     string
	wg         sync.WaitGroup
	logger     *zap.Logger
}

// New return new Store.
func New(dataStore data.Store, cacheStore data.Store, logger *zap.Logger) *Store {
	inf := NewInformer(dataStore, cacheStore, logger)
	return &Store{
		dataStore:  dataStore,
		cacheStore: cacheStore,
		informer:   inf,
		logger:     logger.Named("pkg/store/cache:Store"),
	}
}

// Start start the logic behind the cache store (utilizes st.dataStore and st.cacheStore).
func (st *Store) Start(stopCh chan struct{}) error {
	var errs []error

	st.wg.Add(1)
	go func() {
		defer st.wg.Done()
		errs = append(errs, st.informer.Start(stopCh))
	}()

	for {
		select {
		case <-stopCh:
			st.wg.Wait()
			return errors.Concat(errs...)
		}
	}
}

// Get get a value for a key.
func (st *Store) Get(ctx context.Context, key string) ([]byte, error) {
	result, err := st.cacheStore.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		result, err = st.dataStore.Get(ctx, key)
		if err != nil {
			return result, err
		}
	}
	err = st.cacheStore.Put(ctx, key, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// List return a list of key value pairs where the key is the directory.
func (st *Store) List(ctx context.Context, key string) (map[string][]byte, error) {
	result, err := st.cacheStore.GetRecursive(ctx, key)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		result, err = st.dataStore.GetRecursive(ctx, key)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	for rk := range result {
		err = st.cacheStore.Put(ctx, "", result[rk])
		if err != nil {
			return nil, err
		}
	}
	return result, err
}

// Put creates or updates if exists, a key value pair.
func (st *Store) Put(ctx context.Context, key string, value []byte) error {
	if err := st.dataStore.Put(ctx, key, value); err != nil {
		return err
	}
	return st.cacheStore.Put(ctx, key, value)
}

// Delete delete a key value pair.
func (st *Store) Delete(ctx context.Context, key string) error {
	if err := st.dataStore.Delete(ctx, key); err != nil {
		return err
	}
	return st.cacheStore.Delete(ctx, key)
}

// Watch watch a key for creation, changes and deletion.
func (st *Store) Watch(ctx context.Context, key string) (chan *storecommon.InformerResult, error) {
	var watch clientv3.WatchChan
	if !st.informer.DataStoreChExists(key) {
		var err error
		watch, err = st.dataStore.Watch(ctx, key)
		if err != nil {
			return nil, err
		}
	}
	return st.informer.Watch(ctx, key, watch)
}

// WatchRecursively watch a directory for creation, changes and deletion.
func (st *Store) WatchRecursively(ctx context.Context, key string) (chan *storecommon.InformerResult, error) {
	var watch clientv3.WatchChan
	if !st.informer.DataStoreChExists(key) {
		var err error
		watch, err = st.dataStore.WatchRecursively(ctx, key)
		if err != nil {
			return nil, err
		}
	}
	return st.informer.Watch(ctx, key, watch)
}

// Close data and cache store adapter.
func (st *Store) Close() error {
	return errors.Concat(st.dataStore.Close(), st.cacheStore.Close())
}
