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

package v1alpha

import (
	"context"
	"time"

	jsoniter "github.com/json-iterator/go"

	core_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	events_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// NodesService handler for config events.
type NodesService struct {
	NodesServer
}

// New returns a NodesService
func New() NodesServer {
	return &NodesService{}
}

// Get get a Node.
func (n *NodesService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	resp := &GetResponse{}

	return resp, nil
}

// List list Nodes.
func (n *NodesService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	resp := &ListResponse{}

	return resp, nil
}

// Add add a Node.
func (n *NodesService) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	req.Node.SetDefaults()
	resp := &AddResponse{
		Node: req.Node,
	}

	return resp, nil
}

// Update update a Node.
func (n *NodesService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	req.Node.SetDefaults()
	resp := &UpdateResponse{}

	return resp, nil
}

// Delete delete a Node.
func (n *NodesService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	resp := &DeleteResponse{}

	return resp, nil
}

// Watch Watch Nodes.
func (n *NodesService) Watch(req *WatchRequest, watch Nodes_WatchServer) error {
	for {
		if err := Nodes_WatchServer.Send(watch, &WatchResponse{
			Node: &Node{
				Metadata: &core_v1alpha.ObjectMetadata{
					Name: "node",
				},
			},
			Event: &events_v1alpha.Event{
				Metadata: &core_v1alpha.ObjectMetadata{
					Name: "node",
				},
				Spec: &events_v1alpha.EventSpec{
					Summary: "Update",
				},
			},
		}); err != nil {
			return err
		}
		<-time.After(5 * time.Second)
	}
}
