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

	core_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	nodes_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func magicRun(stopCh chan struct{}) error {
	// Watch for node changes of itself
	hostname, err := os.Hostname()
	if err != nil {
		logger.Error("failed to get hostname", zap.Error(err))
		return err
	}

grpcConn:
	for {
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		grpcClient, err := grpc.Dial("127.0.0.1:1337", opts...)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		defer grpcClient.Close()
	grpcNodesClient:
		for {
			nodesClient := nodes_v1alpha.NewNodesClient(grpcClient)

			in := &nodes_v1alpha.WatchRequest{
				WatchOptions: &core_v1alpha.WatchOptions{
					Name: hostname,
				},
			}
			nodesWatcher, err := nodesClient.Watch(context.Background(), in)
			if err != nil {
				logger.Error("failed to watch nodes", zap.Error(err))
				return err
			}

			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				select {
				case <-stopCh:
					// TODO Move this to own code which is just for maintaining GRPC connection
					grpcClient.Close()
					logger.Error("grpc client closing conncetion failed ", zap.Error(err))
				}
			}()

		nodesWatch:
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
				fmt.Printf("TEST: %+v\n", nodeResp)
			}
		}
	}
	return nil
}
