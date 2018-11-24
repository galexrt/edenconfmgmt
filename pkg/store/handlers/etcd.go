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
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/galexrt/edenconfmgmt/pkg/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
)

const (
	flagETCDCACert                = "etcd-cacert"
	flagETCDCert                  = "etcd-cert"
	flagETCDCommandTimeout        = "etcd-command-timeout"
	flagETCDDebug                 = "etcd-debug"
	flagETCDDialTimeout           = "etcd-dial-timeout"
	flagETCDDiscoverySRV          = "etcd-discovery-srv"
	flagETCDEndpoints             = "etcd-endpoints"
	flagETCDInsecureDiscovery     = "etcd-insecure-discovery"
	flagETCDInsecureSkipTLSVerify = "etcd-insecure-skip-tls-verify"
	flagETCDInsecureTransport     = "etcd-insecure-transport"
	flagETCDKeepaliveTime         = "etcd-keepalive-time"
	flagETCDKeepaliveTimeout      = "etcd-keepalive-timeout"
	flagETCDKey                   = "etcd-key"
	flagETCDUser                  = "etcd-user"
)

// ETCD implementation of store.Store interface for ETCD.
type ETCD struct {
	prefix string
	client *clientv3.Client
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

// ETCDWatcher watcher for ETCD implementing the store.Watcher interface.
type ETCDWatcher struct {
	watcher clientv3.Watcher
	store.Watcher
}

func init() {
	flagRegisters["etcd"] = func(cmd *cobra.Command) {
		cmd.PersistentFlags().String(flagETCDCACert, "", "ETCD: verify certificates of TLS-enabled secure servers using this CA bundle")
		cmd.PersistentFlags().String(flagETCDCert, "", "ETCD: identify secure client using this TLS certificate file")
		cmd.PersistentFlags().Duration(flagETCDCommandTimeout, 5*time.Second, "ETCD: timeout for short running command (excluding dial timeout)")
		cmd.PersistentFlags().Bool(flagETCDDebug, false, "ETCD: enable client-side debug logging")
		cmd.PersistentFlags().Duration(flagETCDDialTimeout, 2*time.Second, "ETCD: dial timeout for client connections")
		cmd.PersistentFlags().String(flagETCDDiscoverySRV, "", "ETCD: domain name to query for SRV records describing cluster endpoints")
		cmd.PersistentFlags().StringSlice(flagETCDEndpoints, []string{}, "ETCD: gRPC endpoints")
		cmd.PersistentFlags().Bool(flagETCDInsecureDiscovery, true, "ETCD: accept insecure SRV records describing cluster endpoints")
		cmd.PersistentFlags().Bool(flagETCDInsecureSkipTLSVerify, false, "ETCD: skip server certificate verification")
		cmd.PersistentFlags().Bool(flagETCDInsecureTransport, true, "ETCD: disable transport security for client connections")
		cmd.PersistentFlags().Duration(flagETCDKeepaliveTime, 30*time.Second, "ETCD: keepalive time for client connections")
		cmd.PersistentFlags().Duration(flagETCDKeepaliveTimeout, 30*time.Second, "ETCD: keepalive timeout for client connections")
		cmd.PersistentFlags().String(flagETCDKey, "", "ETCD: identify secure client using this TLS key file")
		cmd.PersistentFlags().String(flagETCDUser, "", "ETCD: username[:password] for authentication (prompt if password is not supplied)")
	}
	handlers["etcd"] = NewETCD
}

// NewETCD create new ETCD store.
func NewETCD() (store.Store, error) {
	etcd := &ETCD{}
	user := viper.GetString(flagETCDUser)
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
		Endpoints:   viper.GetStringSlice(flagETCDEndpoints),
		Username:    user,
		Password:    password,
		DialTimeout: viper.GetDuration(flagETCDDialTimeout),
	}

	var err error
	if etcd.client, err = clientv3.New(cfg); err != nil {
		return nil, err
	}

	return etcd, err
}

// SetPrefix set the prefix to prefix all given keys with.
func (st *ETCD) SetPrefix(prefix string) {
	st.prefix = prefix
}

// Get return a specific key.
func (st *ETCD) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := st.client.Get(ctx, path.Join(st.prefix, key))
	if err != nil {
		return "", err
	}
	return string(resp.Kvs[0].Value), nil
}

// Set set a key to a specific value.
func (st *ETCD) Set(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := st.client.Put(ctx, path.Join(st.prefix, key), value)
	return err
}

// SetTTL set a key to a specific value with a TTL.
func (st *ETCD) SetTTL(key string, value string, ttl int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := st.client.Grant(ctx, ttl)
	if err != nil {
		return err
	}
	_, err = st.client.Put(ctx, path.Join(st.prefix, key), value, clientv3.WithLease(resp.ID))
	return err
}

// Delete delete a key value pair.
func (st *ETCD) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := st.client.Delete(ctx, path.Join(st.prefix, key))
	return err
}

// Watch watch a key or directory for creation, changes and deletion.
func (st *ETCD) Watch(key string, dir bool) (store.Watcher, error) {
	/*
		watcher := NewETCDWatcher(st.client.Watcher(path.Join(st.prefix, key), &clientv3.WatcherOptions{Recursive: dir}))
		return watcher, nil
	*/
	return nil, nil
}

/*
// NewETCDWatcher create a new ETCDWatcher from a etcd clientv3.Watcher.
func NewETCDWatcher(watcher clientv3.Watcher) *ETCDWatcher {
	return &ETCDWatcher{
		watcher: watcher,
	}
}

// Watch return channel which "streams" events as they come.
func (w *ETCDWatcher) Watch(ctx context.Context) (chan *store.Response, error) {
	ch := make(chan *store.Response)
	go func() {
		for {
			resp, err := w.watcher.Next(ctx)
			watchResp := &store.Response{
				Error: err,
			}
			if resp.Node != nil {
				watchResp.Key = resp.Node.Key
				watchResp.Value = resp.Node.Value
			}
			ch <- watchResp
			if err != nil {
				return
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}()
	return ch, nil
}
*/
