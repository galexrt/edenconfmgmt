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

package adapters

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/galexrt/edenconfmgmt/pkg/datastore"
	"github.com/galexrt/edenconfmgmt/pkg/datastore/watcher"
)

// Memory "in-memory" cache adapter implementation.
// `sync.Map` is used here due to the nature of parallele access to the (cache) adapter.
type Memory struct {
	cache sync.Map
	// As a map does not have an event/notification system we have a channel.
	datastore.Store
}

func init() {
	adapters["memory"] = NewMemory
}

// NewMemory return new Memory store.
func NewMemory(keyPrefix string) (datastore.Store, error) {
	return &Memory{
		cache: sync.Map{},
	}, nil
}

// Get get a value for a key.
func (st *Memory) Get(ctx context.Context, key string) (string, bool, error) {
	valueLoaded, ok := st.cache.Load(key)
	var value string
	var err error
	if ok {
		value, ok = valueLoaded.(string)
		if !ok {
			err = fmt.Errorf("failed to cast value to string (key: %s)", key)
		}
	}
	return value, ok, err
}

// GetRecursive return a list of keys with values at the given key arg.
func (st *Memory) GetRecursive(ctx context.Context, key string) (map[string]string, bool, error) {
	var results map[string]string
	var errors []error

	keys := st.searchKeysSubstr(key)
	var found bool
	if len(keys) > 0 {
		found = true
	} else {
		return results, found, nil
	}

	for _, fKey := range keys {
		loadedValue, ok := st.cache.Load(fKey)
		if !ok {
			errors = append(errors, fmt.Errorf("unable to load key %s from sync.Map", fKey))
			continue
		}
		results[fKey], ok = loadedValue.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("unable to cast value of key %s from sync.Map", fKey))
		}
	}

	var err error
	if len(errors) > 0 {
		errsConcat := "failed to load/cast value(s): "
		for _, e := range errors {
			errsConcat += e.Error()
			errsConcat += "; "
		}
		err = fmt.Errorf(errsConcat)
	}

	return results, found, err
}

func (st *Memory) searchKeysSubstr(search string) []string {
	var keys []string
	st.cache.Range(func(key, _ interface{}) bool {
		if strings.HasPrefix(search, key.(string)) {
			keys = append(keys, key.(string))
		}
		return true
	})
	return keys
}

// Put put a key value pair.
func (st *Memory) Put(ctx context.Context, key string, value string) error {
	st.cache.Store(key, value)
	return nil
}

// Delete delete a key value pair.
func (st *Memory) Delete(ctx context.Context, key string, recursively bool) error {
	keys := st.searchKeysSubstr(key)
	if len(keys) > 0 {
		return nil
	}
	for _, fKey := range keys {
		st.cache.Delete(fKey)
	}
	return nil
}

// Watch watch a key or directory for creation, changes and deletion.
func (st *Memory) Watch(ctx context.Context, key string, recursive bool) (*watcher.Informer, error) {
	return nil, nil
}

// Close adapter.
func (st *Memory) Close() error {
	return nil
}
