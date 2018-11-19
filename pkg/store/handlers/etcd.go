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

package handlers

import (
	"time"

	"github.com/galexrt/edenconfmgmt/pkg/store"
	flag "github.com/spf13/pflag"
)

// ETCD implementation of store.Store interface for ETCD.
type ETCD struct {
	store.Store
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
	flagRegisters["etcd"] = func(flagSet *flag.FlagSet) error {
		// TODO add all flags for etcd options here.
		return nil
	}
}

// New create new ETCD store.
func New() (*ETCD, error) {
	return &ETCD{}, nil
}

// Get return a specific key.
func (st *ETCD) Get(key string) (string, error) {
	return "", nil
}

// Put set a key to a specific value.
func (st *ETCD) Put(key string, value string) error {
	return nil
}

// PutTTL set a key to a specific value with a TTL.
func (st *ETCD) PutTTL(key string, value string, ttlSeconds int64) error {
	return nil
}

// Delete delete a key value pair.
func (st *ETCD) Delete(key string) error {
	return nil
}

// Watch watch a key for creation, changes and deletion.
func (st *ETCD) Watch(key string) (string, error) {
	return "", nil
}

// WatchDir watch a directory for creation, changes and deletion of keys.
func (st *ETCD) WatchDir(dir string) error {
	return nil
}
