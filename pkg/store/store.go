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

package store

import (
	"time"
)

// Store Interface to abstract get, put, del, watch actions to a store. Where a
// store might be an etcd, consul or even in-memory.
type Store interface {
	// Get return a specific key.
	Get(key string) (string, error)
	// Put set a key to a specific value.
	Set(key string, value string) error
	// PutTTL set a key to a specific value with a TTL.
	SetTTL(key string, value string, ttl time.Duration) error
	// Delete delete a key value pair.
	Delete(key string) error
	// Watch watch a key or directory for creation, changes and deletion.
	Watch(key string, dir bool) (*Watcher, error)
}
