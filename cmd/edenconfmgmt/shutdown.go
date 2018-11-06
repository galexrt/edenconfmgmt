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
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	nodes_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha"
	"github.com/galexrt/edenconfmgmt/pkg/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	sigCounter  = 0
	shutdownCmd = &cobra.Command{
		Use:   "shutdown",
		Short: "Trigger shutdown of edenconfmgmt daemon",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			common.GetLogger()
		},
		RunE: Shutdown,
	}
)

// Shutdown command
func Shutdown(cmd *cobra.Command, args []string) error {
	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			<-sigCh
			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
			logger.Info("signal received, ignoring ...")
			if sigCounter > 0 {
				close(stopCh)
			}
			sigCounter++
		}
	}()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	cc, err := grpc.Dial("127.0.0.1:1337", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer cc.Close()
	nodesClient := nodes_v1alpha.NewNodesServiceClient(cc)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	getReq := &nodes_v1alpha.GetRequest{
		GetOptions: &core_v1alpha.GetOptions{
			Name: hostname,
		},
	}
	getResponse, err := nodesClient.Get(ctx, getReq)
	fmt.Printf("Nodes.Get(): %+v - %+v\n", getResponse, err)
	// TODO mark node as (shut)down(ed) with reason (e.g., shutdown of server, service restart)
	//getResponse.Node.Spec

	return nil
}
