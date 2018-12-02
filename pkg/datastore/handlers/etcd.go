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
	"strings"
	"time"

	"github.com/galexrt/edenconfmgmt/pkg/datastore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/namespace"
	"go.etcd.io/etcd/mvcc/mvccpb"
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

// ETCD implementation of datastore.Store interface for ETCD.
type ETCD struct {
	client *clientv3.Client
	datastore.Store
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

// NewETCD create new ETCD
func NewETCD(keyPrefix string) (datastore.Store, error) {
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

	// Use namespaced KV client
	cli.KV = namespace.NewKV(cli.KV, keyPrefix)
	cli.Watcher = namespace.NewWatcher(cli.Watcher, keyPrefix)
	cli.Lease = namespace.NewLease(cli.Lease, keyPrefix)

	etcd.client = cli

	// TODO Should a etcd.Status() call be made to one or more given endpoints?
	// See https://godoc.org/go.etcd.io/etcd/clientv3#Maintenance

	return etcd, err
}

// Get return a specific key.
func (st *ETCD) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return st.client.Get(ctx, key, opts...)
}

// Put set a key to a specific value.
func (st *ETCD) Put(ctx context.Context, key string, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return st.client.Put(ctx, key, value, opts...)
}

// PutTTL set a key to a specific value with a TTL.
func (st *ETCD) PutTTL(ctx context.Context, key string, value string, ttl int64, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	resp, err := st.client.Grant(ctx, ttl)
	if err != nil {
		return nil, err
	}
	opts = append([]clientv3.OpOption{clientv3.WithLease(resp.ID)}, opts...)
	return st.client.Put(ctx, key, value, opts...)
}

// Delete delete a key value pair.
func (st *ETCD) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	return st.client.Delete(ctx, key, opts...)
}

// Watch watch a key or directory for creation, changes and deletion.
func (st *ETCD) Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	return st.client.Watch(ctx, key, opts...)
}

// EtcdEventStateToHandlerState convert Etcd client v3 event type to our own states.
func EtcdEventStateToHandlerState(eventType mvccpb.Event_EventType, isCreate bool, isModify bool) string {
	if isCreate {
		return datastore.ResponseStateCreated
	}
	if isModify {
		return datastore.ResponseStateUpdated
	}
	if eventType == mvccpb.DELETE {
		return datastore.ResponseStateDeleted
	}
	return datastore.ResponseStateUnknown
}
