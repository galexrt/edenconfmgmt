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

package cache

import (
	"github.com/galexrt/edenconfmgmt/pkg/datastore"
)

// Cache cache store
type Cache struct {
	datastore.Store
	store datastore.Store
	cache map[string]string
}

// CacheStoreAdapter
type CacheStoreAdapter interface {
	// Get get a value for a key.
	Get(key string) (string, error)
	// IsSet return bool state if a key exists.
	IsSet(key string) (bool, error)
	// Put put a key value pair.
	Put(key string, value string) error
	// Delete delete a key value pair.
	Delete(key string) error
}

// New return a new Cache store
func New(store datastore.Store) *Cache {
	return &Cache{
		store: store,
		cache: map[string]string{},
	}
}
