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

package adapters

import (
	"fmt"
	"sync"
)

// Memory "in-memory" cache adapter implementation.
// `sync.Map` is used here due to the nature of parallele access to the (cache) adapter.
type Memory struct {
	CacheStoreAdapter
	cache sync.Map
}

// NewMemory return new Memory adapter.
func NewMemory() Memory {
	return Memory{
		cache: sync.Map{},
	}
}

// Get get a value for a key.
func (adp *Memory) Get(key string) (string, bool, error) {
	valueLoaded, ok := adp.cache.Load(key)
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

// IsSet return bool state if a key exists.
func (adp *Memory) IsSet(key string) (bool, error) {
	_, ok := adp.cache.Load(key)
	return ok, nil
}

// Put put a key value pair.
func (adp *Memory) Put(key string, value string) error {
	adp.cache.Store(key, value)
	return nil
}

// Delete delete a key value pair.
func (adp *Memory) Delete(key string) error {
	return adp.Delete(key)
}

// Close adapter.
func (adp *Memory) Close() error {
	return nil
}
