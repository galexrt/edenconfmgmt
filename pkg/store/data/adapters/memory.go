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

	"github.com/galexrt/edenrun/pkg/store/data"
	"go.etcd.io/etcd/clientv3"
)

// Memory "in-memory" cache adapter implementation.
// `sync.Map` is used here due to the nature of parallele access to the (cache) adapter.
type Memory struct {
	cache sync.Map
	// As a map does not have an event/notification system we have a channel.
	data.Store
}

func init() {
	adapters["memory"] = NewMemory
}

// NewMemory return new Memory store.
func NewMemory(ctx context.Context) (data.Store, error) {
	return &Memory{
		cache: sync.Map{},
	}, nil
}

// Get get a value for a key.
func (st *Memory) Get(ctx context.Context, key string) ([]byte, error) {
	valueLoaded, ok := st.cache.Load(key)
	var value []byte
	var err error
	if ok {
		value, ok = valueLoaded.([]byte)
		if !ok {
			err = fmt.Errorf("failed to cast value to []byte (key: %s)", key)
		}
	}
	return value, err
}

// GetRecursive return a list of keys with values at the given key arg.
func (st *Memory) GetRecursive(ctx context.Context, key string) (map[string][]byte, error) {
	var results map[string][]byte
	var errors []error

	keys := st.searchKeysSubstr(key)
	if len(keys) == 0 {
		return results, nil
	}

	for _, fKey := range keys {
		loadedValue, ok := st.cache.Load(fKey)
		if !ok {
			errors = append(errors, fmt.Errorf("unable to load key %s from sync.Map", fKey))
			continue
		}
		results[fKey], ok = loadedValue.([]byte)
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

	return results, err
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
func (st *Memory) Put(ctx context.Context, key string, value []byte) error {
	st.cache.Store(key, value)
	return nil
}

// Delete delete a key value pair.
func (st *Memory) Delete(ctx context.Context, key string) error {
	if key[len(key)-1:] == "/" {
		keys := st.searchKeysSubstr(key)
		if len(keys) > 0 {
			return nil
		}
		for _, sKey := range keys {
			st.cache.Delete(sKey)
		}
	} else {
		st.cache.Delete(key)
	}
	return nil
}

// Watch watch a key for creation, changes and deletion.
func (st *Memory) Watch(ctx context.Context, key string) (clientv3.WatchChan, error) {
	return nil, nil
}

// WatchRecursively watch a directory for creation, changes and deletion.
func (st *Memory) WatchRecursively(ctx context.Context, key string) (clientv3.WatchChan, error) {
	return nil, nil
}

// Close adapter.
func (st *Memory) Close() error {
	return nil
}
