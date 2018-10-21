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
)

// NodesService
type NodesService struct {
	NodesServiceServer
}

// NewServer returns a NodesService
func NewServer() NodesServiceServer {
	return &NodesService{}
}

// Version
func (n *NodesService) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return &VersionResponse{
		Version: "1",
	}, nil
}

// AddNode
func (n *NodesService) AddNode(context.Context, *Node) (*AddNodeResponse, error) {
	return &AddNodeResponse{}, nil
}
