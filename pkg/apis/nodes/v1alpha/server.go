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

package v1alpha

import (
	"context"
	"fmt"
	"time"

	"github.com/json-iterator/go"

	core_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	events_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	"github.com/galexrt/edenconfmgmt/pkg/datastore"
	utilsapi "github.com/galexrt/edenconfmgmt/pkg/utils/api"
	"go.etcd.io/etcd/clientv3"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// NodesService handler for config events.
type NodesService struct {
	store datastore.Store
	NodesServer
}

// New returns a NodesService
func New(dataStore datastore.Store) NodesServer {
	return &NodesService{
		store: dataStore,
	}
}

// Get get a Node.
func (n *NodesService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	resp := &GetResponse{}
	storeResp, err := n.store.Get(ctx, utilsapi.ObjectPath(DataStorePath, req.GetOptions.Name), clientv3.WithPrefix())
	if err != nil {
		resp.Error = utilsapi.ErrorToErrorResponse(err)
		return resp, nil
	}
	var (
		metadata *core_v1alpha.ObjectMetadata
		spec     *Spec
	)
	for _, kv := range storeResp.Kvs {
		var part interface{}

		switch string(kv.Key) {
		case "metadata":
			part = &core_v1alpha.ObjectMetadata{}
			metadata = part.(*core_v1alpha.ObjectMetadata)
		case "spec":
			part = &Spec{}
			spec = part.(*Spec)
		}
		if err = json.Unmarshal(kv.Value, part); err != nil {
			resp.Error = utilsapi.ErrorToErrorResponse(fmt.Errorf(""))
			return resp, nil
		}
	}
	resp.Node = &Node{
		Metadata: metadata,
		Spec:     spec,
	}
	return resp, nil
}

// List list Nodes.
func (n *NodesService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	resp := &ListResponse{}

	return resp, nil
}

// Add add a Node.
func (n *NodesService) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	node := req.Node
	resp := &AddResponse{
		Node: node,
	}
	node.SetDefaults()
	if err := node.Validate(); err != nil {
		resp.Error = utilsapi.ErrorToErrorResponse(err)
		return resp, nil
	}
	nodeString, err := json.Marshal(node)
	if err != nil {
		resp.Error = utilsapi.ErrorToErrorResponse(err)
		return resp, nil
	}
	_, err = n.store.Put(ctx, utilsapi.ObjectPath(DataStorePath, resp.Node.Metadata.Name), string(nodeString))
	resp.Error = utilsapi.ErrorToErrorResponse(err)
	return resp, nil
}

// Update update a Node.
func (n *NodesService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
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
