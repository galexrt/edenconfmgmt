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
	"github.com/spf13/cobra"
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
		cmd.PersistentFlags().Duration(flagETCDKeepaliveTime, 2*time.Second, "ETCD: keepalive time for client connections")
		cmd.PersistentFlags().Duration(flagETCDKeepaliveTimeout, 6*time.Second, "ETCD: keepalive timeout for client connections")
		cmd.PersistentFlags().String(flagETCDKey, "", "ETCD: identify secure client using this TLS key file")
		cmd.PersistentFlags().String(flagETCDUser, "", "ETCD: username[:password] for authentication (prompt if password is not supplied)")
	}
}

// New create new ETCD store.
func New() (*ETCD, error) {
	// TODO use flag values to generate config for etcd client
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
