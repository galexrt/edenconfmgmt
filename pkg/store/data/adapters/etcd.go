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
	"time"

	"github.com/galexrt/edenrun/pkg/store/data"
	"go.etcd.io/etcd/clientv3"
)

// ETCD implementation of Handler interface for ETCD.
type ETCD struct {
	cli *clientv3.Client
	data.Store
}

// ETCDOptions options for ETCD connection.
// Just the flags of `ETCDCTL_API=3 etcdctl` (etcdtctl version `3.3.8`, API version `3.3`) copied.
type ETCDOptions struct {
	CACert                string
	Cert                  string
	CommandTimeout        time.Duration
	Debug                 bool
	DialTimeout           time.Duration
	DiscoverySRV          string
	Endpoints             []string
	InsecureDiscovery     bool
	InsecureSkipTLSVerify bool
	InsecureTransport     bool
	KeepaliveTime         time.Duration
	KeepaliveTimeout      time.Duration
	Key                   string
	User                  string
}

func init() {
	adapters["etcd"] = NewETCD
}

// NewETCD create new ETCD
func NewETCD(ctx context.Context) (data.Store, error) {
	etcd := &ETCD{}
	user := ""
	password := ""
	if user != "" {
		splitUserInfo := strings.Split(user, ":")
		if len(splitUserInfo) != 2 {
			return nil, fmt.Errorf("ETCD: user flag does not contain a password")
		}
		user = splitUserInfo[0]
		password = splitUserInfo[1]
	}
	// TODO look at etcdctl for better client and/or transport config "generation"
	// because currently most flags defined here are not used..
	cfg := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		Username:    user,
		Password:    password,
		DialTimeout: 2 * time.Second,
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	etcd.cli = cli

	// TODO Should a etcd.Status() call be made to one or more given endpoints?
	// See https://godoc.org/go.etcd.io/etcd/clientv3#Maintenance
	//etcd.cli.Status(ctx, endpoint)

	return etcd, err
}

// Get return a specific key.
func (st *ETCD) Get(ctx context.Context, key string) ([]byte, error) {
	resp, err := st.cli.Get(ctx, key)
	if err != nil {
		return []byte{}, err
	}
	for _, kv := range resp.Kvs {
		return kv.Value, nil
	}
	return []byte{}, nil
}

// GetRecursive return a list of keys with values at the given key arg.
func (st *ETCD) GetRecursive(ctx context.Context, key string) (map[string][]byte, error) {
	var results map[string][]byte
	opts := []clientv3.OpOption{clientv3.WithPrefix()}
	resp, err := st.cli.Get(ctx, key, opts...)
	if err != nil {
		return results, err
	}
	for _, kv := range resp.Kvs {
		results[string(kv.Key)] = kv.Value
	}
	return results, nil
}

// Put set a key to a specific value.
func (st *ETCD) Put(ctx context.Context, key string, value []byte) error {
	_, err := st.cli.Put(ctx, key, string(value))
	return err
}

// Delete delete a key value pair.
func (st *ETCD) Delete(ctx context.Context, key string) error {
	var opts []clientv3.OpOption
	if key[len(key)-1:] == "/" {
		opts = append(opts, clientv3.WithPrefix())
	}
	_, err := st.cli.Delete(ctx, key, opts...)
	return err
}

// Watch watch a key or directory for creation, changes and deletion.
func (st *ETCD) Watch(ctx context.Context, key string) (clientv3.WatchChan, error) {
	return st.cli.Watch(ctx, key), nil
}

// WatchRecursively watch a directory for creation, changes and deletion.
// If the `key` ends in `/` it will be a recursive watch.
func (st *ETCD) WatchRecursively(ctx context.Context, key string) (clientv3.WatchChan, error) {
	return st.cli.Watch(ctx, key, []clientv3.OpOption{clientv3.WithPrefix()}...), nil
}

// Close closes the store and cancels all watches (if supported).
func (st *ETCD) Close() error {
	return st.cli.Close()
}
