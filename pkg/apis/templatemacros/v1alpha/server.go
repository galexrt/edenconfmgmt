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

// TemplateMacrosService handler for config events.
type TemplateMacrosService struct {
	TemplateMacrosServer
}

// New returns a TemplateMacrosService
func New() TemplateMacrosServer {
	return &TemplateMacrosService{}
}

// Get get a TemplateMacro.
func (n *TemplateMacrosService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return &GetResponse{}, nil
}

// List list TemplateMacros.
func (n *TemplateMacrosService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return &ListResponse{}, nil
}

// Add add a TemplateMacro.
func (n *TemplateMacrosService) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return &AddResponse{}, nil
}

// Update update a TemplateMacro.
func (n *TemplateMacrosService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return &UpdateResponse{}, nil
}

// Delete delete a TemplateMacro.
func (n *TemplateMacrosService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return &DeleteResponse{}, nil
}

// Watch Watch TemplateMacros.
func (n *TemplateMacrosService) Watch(req *WatchRequest, watch TemplateMacros_WatchServer) error {
	return nil
}
