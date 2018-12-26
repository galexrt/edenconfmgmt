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

package datastore

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

// ResponseState response state type.
type ResponseState string

const (
	// ResponseStateCreated created response state.
	ResponseStateCreated ResponseState = "Created"
	// ResponseStateUpdated updated response state.
	ResponseStateUpdated ResponseState = "Updated"
	// ResponseStateDeleted deleted response state.
	ResponseStateDeleted ResponseState = "Deleted"
	// ResponseStateUnknown unknown response state.
	ResponseStateUnknown ResponseState = "Unknown"
)

// Store Interface to abstract get, put, del, watch actions to etcd datastore.
// TODO does it make sense to have different datastores?
type Store interface {
	// SetPrefix set the prefix to prefix all given keys with.
	SetKeyPrefix(prefix string)

	// Get return a specific key (range).
	Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error)
	// Put set a key to a specific value.
	Put(ctx context.Context, key string, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error)
	// PutTTL set a key to a specific value with a TTL.
	PutTTL(ctx context.Context, key string, value string, ttl int64, opts ...clientv3.OpOption) (*clientv3.PutResponse, error)
	// Delete delete a key value pair.
	Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error)
	// Watch watch a key or directory for creation, changes and deletion.
	Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan
	// Close closes the store and cancels all watch requests with it.
	Close() error
}
