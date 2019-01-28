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

package main

import (
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"

	core_v1 "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	nodes_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha"
	"github.com/galexrt/edenconfmgmt/pkg/store/object"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func magicRun(stopCh chan struct{}, store *object.Store) error {
	logger.Info("magicRun started")
	// Watch for node changes of itself

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	grpcClient, err := grpc.Dial("127.0.0.1:1337", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer grpcClient.Close()
	nodesClient := nodes_v1alpha.NewNodesClient(grpcClient)

	go func() {
		<-stopCh
		grpcClient.Close()
	}()

	// TODO add logic to refresh hostname every once in a while
	hostname, err := os.Hostname()
	if err != nil {
		logger.Error("failed to get hostname", zap.Error(err))
		return err
	}

	go func() {
		for {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			logger.Info("---> store.PUT NOW TEST")
			obj := &nodes_v1alpha.Node{
				Metadata: &core_v1.ObjectMetadata{
					Name: hostname,
				},
				Spec: &nodes_v1alpha.Spec{
					Network: &nodes_v1alpha.Network{
						Fqdn: hostname + ".example.com",
					},
				},
			}
			_, err = nodesClient.Create(ctx, &nodes_v1alpha.CreateRequest{
				Node: obj,
			})
			if err != nil {
				logger.Error("---> store.CREATE ERROR", zap.Error(err))
				time.Sleep(5 * time.Second)
				return
			}
			logger.Info("---> store.CREATE Done")
			resp, err := nodesClient.Get(ctx, &nodes_v1alpha.GetRequest{
				Options: &core_v1.GetOptions{
					Name: hostname,
				}})
			if err != nil {
				logger.Info("---> store.GET", zap.Error(err))
				time.Sleep(5 * time.Second)
				return
			}
			logger.Info("---> store.GET TEST RESULT", zap.Any("result", resp))
			time.Sleep(5 * time.Second)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	watcher, err := nodesClient.Watch(ctx, &nodes_v1alpha.WatchRequest{
		Options: &core_v1.WatchOptions{
			Name: hostname,
		}})
	logger.Info("---> WATCH ERROR", zap.Error(err))

	for {
		res, err := watcher.Recv()
		logger.Info("---> WATCH RESULT", zap.Any("result", res), zap.Error(err))
		if err != nil && err != grpc.ErrServerStopped {
			return err
		}
	}

	wg := &sync.WaitGroup{}

	for {
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		// TODO Move this logic to it's own package and init it in the cmd part using the neighbors flag
		grpcClient, err := grpc.Dial("127.0.0.1:1337", opts...)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
			return err
		}
		defer grpcClient.Close()

		for {
			nodesClient := nodes_v1alpha.NewNodesClient(grpcClient)

			wg.Add(1)
			go func() {
				defer wg.Done()
				select {
				case <-stopCh:
					// TODO Move this to own code which is just for maintaining GRPC connection
					err := grpcClient.Close()
					if err != nil {
						logger.Error("grpc client closing conncetion failed ", zap.Error(err))
					} else {
						logger.Info("grpc client closed")
					}
				}
			}()

			/*maxTry := 3
			for try := 1; try <= maxTry; try++ {
				if _, err := registerNode(nodesClient, hostname); err != nil {
					logger.Error(fmt.Sprintf("failed to register node (try %d of %d)", try, maxTry), zap.Error(err))
					if try >= maxTry {
						logger.Error(fmt.Sprintf("failed to register node after %d tries", maxTry))
						return err
					}
					continue
				}
				break
			}*/

			in := &nodes_v1alpha.WatchRequest{
				Options: &core_v1.WatchOptions{
					Name: hostname,
				},
			}
			nodesWatcher, err := nodesClient.Watch(context.Background(), in)
			if err != nil && err != grpc.ErrClientConnClosing {
				logger.Error("failed to watch nodes", zap.Error(err))
				return err
			}

			for {
				nodeResp, err := nodesWatcher.Recv()
				if err == io.EOF {
					logger.Warn("watch closed, restarting watch")
					break
				}
				if err != nil {
					logger.Error("failed to watch nodes", zap.Error(err))
					return err
				}
				logger.Info("TEST", zap.Any("nodeResp", nodeResp))
			}
		}
	}
}
