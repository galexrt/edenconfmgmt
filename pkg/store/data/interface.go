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

package data

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

// Store Interface to abstract get, put, del, watch actions to etcd datastore.
type Store interface {
	// Get return a specific key value pair.
	Get(ctx context.Context, key string) ([]byte, error)
	// GetRecursive return a list of keys with values at the given key arg.
	GetRecursive(ctx context.Context, key string) (map[string][]byte, error)
	// Put set a key to a specific value.
	Put(ctx context.Context, key string, value []byte) error
	// Delete delete a key value pair.
	Delete(ctx context.Context, key string) error
	// Watch watch a key for creation, changes and deletion.
	Watch(ctx context.Context, key string) (clientv3.WatchChan, error)
	// WatchRecursively watch a directory for creation, changes and deletion.
	WatchRecursively(ctx context.Context, key string) (clientv3.WatchChan, error)
	// Close closes the store and cancels all watches (if supported).
	Close() error

	// Status returns the "status" of the store (e.g., if it is healthy)
	Status() *Status
}

// Status store status structure.
type Status struct {
	// Healthy true for healthy, false for unhealthy.
	Healthy bool
}
