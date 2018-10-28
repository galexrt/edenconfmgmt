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

	core_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
)

// NodesService handler for node events.
type NodesService struct {
	NodesServiceServer
}

// NewServer returns a NodesService
func NewServer() NodesServiceServer {
	return &NodesService{}
}

// Version return core_v1alpha.VersionResponse.
func (n *NodesService) Version(context.Context, *core_v1alpha.VersionRequest) (*core_v1alpha.VersionResponse, error) {
	return &core_v1alpha.VersionResponse{
		Version: "1",
	}, nil
}

// AddNode add a node.
func (n *NodesService) AddNode(context.Context, *Node) (*AddNodeResponse, error) {
	return &AddNodeResponse{}, nil
}
