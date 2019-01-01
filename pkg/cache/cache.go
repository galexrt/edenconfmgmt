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
	"github.com/galexrt/edenconfmgmt/pkg/cache/adapters"
	"github.com/galexrt/edenconfmgmt/pkg/datastore"
)

// Cache cache store
type Cache struct {
	datastore.Store
	store datastore.Store
	cache adapters.CacheStoreAdapter
}

// New return a new Cache store
func New(store datastore.Store) *Cache {
	adapter := adapters.NewMemory()
	return &Cache{
		store: store,
		cache: &adapter,
	}
}
