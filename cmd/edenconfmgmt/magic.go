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
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger.Info("---> store.PUT NOW TEST")
	obj := &nodes_v1alpha.Node{
		Metadata: &core_v1.ObjectMetadata{
			Name: hostname,
			Labels: map[string]string{
				"app":  "node",
				"nope": "123",
			},
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
	if err != nil && err != object.ErrAlreadyExists {
		logger.Error("---> store.CREATE ERROR", zap.Error(err))
		time.Sleep(5 * time.Second)
		return err
	}
	logger.Info("---> store.CREATE Done")

	go func() {
		for {
			ctx1, cancel1 := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel1()
			_, err = nodesClient.Update(ctx1, &nodes_v1alpha.UpdateRequest{
				Options: &core_v1.UpdateOptions{},
				Node: &nodes_v1alpha.Node{
					Metadata: &core_v1.ObjectMetadata{
						Name: hostname + fmt.Sprintf("%d", rand.Int()),
						Labels: map[string]string{
							"app":  "node",
							"nope": "123",
						},
					},
					Spec: &nodes_v1alpha.Spec{
						Network: &nodes_v1alpha.Network{
							Fqdn: time.Now().String(),
						},
					},
				}})
			if err != nil {
				logger.Info("---> store.GET ERROR", zap.Error(err))
				time.Sleep(5 * time.Second)
				return
			}
			time.Sleep(5 * time.Second)
		}
	}()

	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()
	watcher, err := nodesClient.Watch(ctx2, &nodes_v1alpha.WatchRequest{
		Options: &core_v1.WatchOptions{
			LabelSelectors: []string{
				"app=node",
				"nope=123",
			},
		}})
	logger.Info("---> WATCH ERROR", zap.Error(err))

	for {
		res, err := watcher.Recv()
		logger.Info("---> WATCH RESULT", zap.Any("result", res), zap.Error(err))
		if err != nil && err != grpc.ErrServerStopped {
			return err
		}
		spew.Dump(res)
	}
}
