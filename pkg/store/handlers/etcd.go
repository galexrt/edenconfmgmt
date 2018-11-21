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
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/galexrt/edenconfmgmt/pkg/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/client"
)

const (
	flagETCDCACert                = "etcd-cacert"
	flagETCDCert                  = "etcd-cert"
	flagETCDCommandTimeout        = "etcd-cacert"
	flagETCDDebug                 = "etcd-cacert"
	flagETCDDialTimeout           = "etcd-cacert"
	flagETCDDiscoverySRV          = "etcd-cacert"
	flagETCDEndpoints             = "etcd-cacert"
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
	client client.Client
	kapi   client.KeysAPI
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
	watcher client.Watcher
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
}

// New create new ETCD store.
func New() (*ETCD, error) {
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
	cfg := client.Config{
		Endpoints: viper.GetStringSlice(flagETCDEndpoints),
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				KeepAlive: viper.GetDuration(flagETCDKeepaliveTime),
				Timeout:   viper.GetDuration(flagETCDKeepaliveTimeout),
			}).Dial,
			TLSHandshakeTimeout: viper.GetDuration(flagETCDDialTimeout),
		},
		Username: user,
		Password: password,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
		SelectionMode:           client.EndpointSelectionRandom,
	}
	var err error
	if etcd.client, err = client.New(cfg); err != nil {
		return nil, err
	}
	etcd.kapi = client.NewKeysAPI(etcd.client)
	return etcd, nil
}

// Get return a specific key.
func (st *ETCD) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := st.kapi.Get(ctx, key, &client.GetOptions{Quorum: true})
	if err != nil {
		return "", err
	}
	return resp.Node.Value, nil
}

// Set set a key to a specific value.
func (st *ETCD) Set(key string, value string) error {
	return st.SetTTL(key, value, 0*time.Second)
}

// SetTTL set a key to a specific value with a TTL.
func (st *ETCD) SetTTL(key string, value string, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := st.kapi.Set(ctx, key, value, &client.SetOptions{TTL: ttl})
	return err
}

// Delete delete a key value pair.
func (st *ETCD) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := st.kapi.Delete(ctx, key, &client.DeleteOptions{})
	return err
}

// Watch watch a key or directory for creation, changes and deletion.
func (st *ETCD) Watch(key string, dir bool) (*ETCDWatcher, error) {
	watcher := NewETCDWatcher(st.kapi.Watcher(key, &client.WatcherOptions{Recursive: dir}))
	return watcher, nil
}

// NewETCDWatcher create a new ETCDWatcher from a etcd client.Watcher.
func NewETCDWatcher(watcher client.Watcher) *ETCDWatcher {
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
