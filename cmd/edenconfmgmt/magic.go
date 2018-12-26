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

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	core_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	nodes_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha"
	"github.com/galexrt/edenconfmgmt/pkg/datastore"
	"github.com/galexrt/edenconfmgmt/pkg/cache/sharedinformer"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func magicRun(stopCh chan struct{}, dataStore datastore.Store) error {
	// Watch for node changes of itself

	pCtx, cancel1 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel1()
	node := &nodes_v1alpha.Node{
		Metadata: &core_v1alpha.ObjectMetadata{
			Name: "TEST",
		},
	}
	out, err := json.Marshal(node)
	if err != nil {
		return err
	}
	_, err = dataStore.Put(pCtx, "/test123/my-object", string(out))
	if err != nil {
		return err
	}

	getWatchFunc := func(ctx context.Context) (clientv3.WatchChan, error) {
		return dataStore.Watch(ctx, "/test123", clientv3.WithPrefix()), nil
	}
	indexer := sharedinformer.New(getWatchFunc)

	ctx, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	go indexer.Run(ctx, stopCh)

	inf := indexer.GetNameInformer("my-object")

	for {
		select {
		case result := <-inf.Watch:
			var object nodes_v1alpha.Node
			err = json.Unmarshal(result.Object, &object)
			if err != nil {
				return err
			}
			break
		case <-stopCh:
			return nil
		}
	}

	return nil

	// TODO add logic to refresh hostname every once in a while
	hostname, err := os.Hostname()
	if err != nil {
		logger.Error("failed to get hostname", zap.Error(err))
		return err
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

			maxTry := 3
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
			}

			// TODO add node "I am alive" loop
			go keepAliveNode(stopCh)

			in := &nodes_v1alpha.WatchRequest{
				WatchOptions: &core_v1alpha.WatchOptions{
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

func registerNode(nodesClient nodes_v1alpha.NodesClient, hostname string) (*nodes_v1alpha.AddResponse, error) {
	// TODO
	addRequest := &nodes_v1alpha.AddRequest{
		Node: &nodes_v1alpha.Node{
			Metadata: &core_v1alpha.ObjectMetadata{
				Kind:       nodes_v1alpha.Kind,
				ApiVersion: nodes_v1alpha.APIVersion,
				Name:       hostname,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return nodesClient.Add(ctx, addRequest)
}

func keepAliveNode(stopCh chan struct{}) {
	// TODO
}
