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

	"github.com/galexrt/edenconfmgmt/pkg/datastore"
	"github.com/galexrt/edenconfmgmt/pkg/datastore/watcher"
)

// Store using the dataStore and a second cacheStore for caching.
type Store struct {
	dataStore   datastore.Store
	cacheStore  datastore.Store
	watchKeeper watcher.Keeper
	datastore.Store
}

// New return new Store.
func New(dataStore datastore.Store, cacheStore datastore.Store) *Store {
	return &Store{
		dataStore:  dataStore,
		cacheStore: cacheStore,
	}
}

// Get get a value for a key.
func (st *Store) Get(ctx context.Context, key string) (string, bool, error) {
	result, found, err := st.cacheStore.Get(ctx, key)
	if err != nil {
		return result, found, err
	}
	if !found {
		result, found, err = st.dataStore.Get(ctx, key)
		if err != nil {
			return result, found, err
		}
	}
	err = st.cacheStore.Put(ctx, key, result)
	return result, found, err
}

// GetRecursive return a list of keys with values at the given key arg.
func (st *Store) GetRecursive(ctx context.Context, key string) (map[string]string, bool, error) {
	results, found, err := st.cacheStore.GetRecursive(ctx, key)
	if err != nil {
		return results, found, err
	}
	if !found {
		results, found, err = st.dataStore.GetRecursive(ctx, key)
	}
	var errors []error
	for k, v := range results {
		if err = st.cacheStore.Put(ctx, k, v); err != nil {
			// TODO Handle error (e.g., build concated error)
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		errsConcat := "failed to put GetRecursive() result into cachestore: "
		for _, e := range errors {
			errsConcat += e.Error()
			errsConcat += "; "
		}
		err = fmt.Errorf(errsConcat)
	}
	return results, found, err
}

// Put put a key value pair.
func (st *Store) Put(ctx context.Context, key string, value string) error {
	if err := st.dataStore.Put(ctx, key, value); err != nil {
		return err
	}
	return st.cacheStore.Put(ctx, key, value)
}

// Delete delete a key value pair.
func (st *Store) Delete(ctx context.Context, key string, recursively bool) error {
	if err := st.dataStore.Delete(ctx, key, recursively); err != nil {
		return err
	}
	return st.cacheStore.Delete(ctx, key, recursively)
}

// Watch watch a key or directory for creation, changes and deletion.
func (st *Store) Watch(ctx context.Context, key string, recursive bool) (*watcher.Informer, error) {
	// TODO Register watch in system and make it possible to close using the Store.Close() function below
	return nil, nil
}

// Close adapter.
func (st *Store) Close() error {
	var err error
	var errors []error
	errors = append(errors, st.dataStore.Close())
	errors = append(errors, st.cacheStore.Close())
	if len(errors) > 0 {
		errsConcat := "failed to close datastore and cachestore: "
		for _, e := range errors {
			if e != nil {
				continue
			}
			errsConcat += e.Error()
			errsConcat += "; "
		}
		err = fmt.Errorf(errsConcat)
	}
	return err
}
