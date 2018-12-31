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
	"github.com/galexrt/edenconfmgmt/pkg/cache"
)

// Memory "in-memory" cache adapter implementation
type Memory struct {
	cache.CacheStoreAdapter
	cache map[string]string
}

// NewMemory return new Memory adapter.
func NewMemory() *Memory {
	return &Memory{
		cache: map[string]string{},
	}
}

// Get get a value for a key.
func (adp *Memory) Get(key string) (string, error) {
	return adp.cache[key], nil
}

// IsSet return bool state if a key exists.
func (adp *Memory) IsSet(key string) (bool, error) {
	_, ok := adp.cache[key]
	return ok, nil
}

// Put put a key value pair.
func (adp *Memory) Put(key string, value string) error {
	adp.cache[key] = value
	return nil
}

// Delete delete a key value pair.
func (adp *Memory) Delete(key string) error {
	delete(adp.cache, key)
	// Because `delete()` does not return anything, we always assume delete is successful.
	return nil
}
