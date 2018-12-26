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
    "go.etcd.io/etcd/clientv3"
    "context"
)

type Cache struct {
    datastore.Store
    cache map[string]interface{}
}

// SetPrefix set the prefix to prefix all given keys with.
func (c *Cache) SetKeyPrefix(prefix string) {

}

// Get return a specific key (range).
func (c *Cache) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
    return nil, nil
}
// Put set a key to a specific value.
func (c *Cache) Put(ctx context.Context, key string, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
    return nil, nil
}
// PutTTL set a key to a specific value with a TTL.
func (c *Cache) PutTTL(ctx context.Context, key string, value string, ttl int64, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
    return nil, nil
}
// Delete delete a key value pair.
func (c *Cache) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
    return nil, nil
}
// Watch watch a key or directory for creation, changes and deletion.
func (c *Cache) Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan {
    return nil
}
// Close closes the store and cancels all watch requests with it.
func (c *Cache) Close() error {
    return nil
}
