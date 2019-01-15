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

	"github.com/galexrt/edenconfmgmt/pkg/store/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/namespace"
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
	flagRegisters["etcd"] = func(prefix string, cmd *cobra.Command) {
		cmd.PersistentFlags().String(prefix+flagETCDCACert, "", "ETCD: verify certificates of TLS-enabled secure servers using this CA bundle")
		cmd.PersistentFlags().String(prefix+flagETCDCert, "", "ETCD: identify secure client using this TLS certificate file")
		cmd.PersistentFlags().Duration(prefix+flagETCDCommandTimeout, 5*time.Second, "ETCD: timeout for short running command (excluding dial timeout)")
		cmd.PersistentFlags().Bool(prefix+flagETCDDebug, false, "ETCD: enable client-side debug logging")
		cmd.PersistentFlags().Duration(prefix+flagETCDDialTimeout, 2*time.Second, "ETCD: dial timeout for client connections")
		cmd.PersistentFlags().String(prefix+flagETCDDiscoverySRV, "", "ETCD: domain name to query for SRV records describing cluster endpoints")
		cmd.PersistentFlags().StringSlice(prefix+flagETCDEndpoints, []string{"127.0.0.1:2379"}, "ETCD: gRPC endpoints")
		cmd.PersistentFlags().Bool(prefix+flagETCDInsecureDiscovery, true, "ETCD: accept insecure SRV records describing cluster endpoints")
		cmd.PersistentFlags().Bool(prefix+flagETCDInsecureSkipTLSVerify, false, "ETCD: skip server certificate verification")
		cmd.PersistentFlags().Bool(prefix+flagETCDInsecureTransport, true, "ETCD: disable transport security for client connections")
		cmd.PersistentFlags().Duration(prefix+flagETCDKeepaliveTime, 30*time.Second, "ETCD: keepalive time for client connections")
		cmd.PersistentFlags().Duration(prefix+flagETCDKeepaliveTimeout, 30*time.Second, "ETCD: keepalive timeout for client connections")
		cmd.PersistentFlags().String(prefix+flagETCDKey, "", "ETCD: identify secure client using this TLS key file")
		cmd.PersistentFlags().String(prefix+flagETCDUser, "", "ETCD: username[:password] for authentication (prompt if password is not supplied)")
	}
	adapters["etcd"] = NewETCD
}

// NewETCD create new ETCD
func NewETCD() (data.Store, error) {
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

	cli, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	etcd.cli = cli

	// TODO Should a etcd.Status() call be made to one or more given endpoints?
	// See https://godoc.org/go.etcd.io/etcd/clientv3#Maintenance

	return etcd, err
}

// SetKeyPrefix set the prefix to prefix all given keys with.
func (st *ETCD) SetKeyPrefix(prefix string) {
	// Use namespaced KV client
	st.cli.KV = namespace.NewKV(st.cli.KV, prefix)
	st.cli.Watcher = namespace.NewWatcher(st.cli.Watcher, prefix)
	st.cli.Lease = namespace.NewLease(st.cli.Lease, prefix)
}

// Get return a specific key.
func (st *ETCD) Get(ctx context.Context, key string) (string, bool, error) {
	resp, err := st.cli.Get(ctx, key)
	if err != nil {
		return "", false, err
	}
	for _, kv := range resp.Kvs {
		return kv.String(), true, nil
	}
	return "", false, nil
}

// GetRecursive return a list of keys with values at the given key arg.
func (st *ETCD) GetRecursive(ctx context.Context, key string) (map[string]string, bool, error) {
	var found bool
	var results map[string]string
	opts := []clientv3.OpOption{clientv3.WithPrefix()}
	resp, err := st.cli.Get(ctx, key, opts...)
	if err != nil {
		return results, false, err
	}
	if len(resp.Kvs) > 0 {
		found = true
	}
	for _, kv := range resp.Kvs {
		results[string(kv.Key)] = string(kv.Value)
	}
	return results, found, nil
}

// Put set a key to a specific value.
func (st *ETCD) Put(ctx context.Context, key string, value string) error {
	_, err := st.cli.Put(ctx, key, value)
	return err
}

// Delete delete a key value pair.
func (st *ETCD) Delete(ctx context.Context, key string, recursive bool) error {
	var opts []clientv3.OpOption
	if recursive {
		opts = append(opts, clientv3.WithPrefix())
	}
	_, err := st.cli.Delete(ctx, key, opts...)
	return err
}

// Watch watch a key or directory for creation, changes and deletion.
func (st *ETCD) Watch(ctx context.Context, key string, recursive bool) (clientv3.WatchChan, error) {
	watch := st.cli.Watch(ctx, key)
	return watch, nil
}

// Close closes the store and cancels all watches (if supported).
func (st *ETCD) Close() error {
	return st.cli.Close()
}
