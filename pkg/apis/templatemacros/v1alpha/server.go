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

// NewServer returns a TemplateMacrosService
func New() TemplateMacrosServer {
	return &TemplateMacrosService{}
}

// Get get a TemplateMacro.
func (n *TemplateMacrosService) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return &GetResponse{}, nil
}

// List list TemplateMacros.
func (n *TemplateMacrosService) List(context.Context, *ListRequest) (*ListResponse, error) {
	return &ListResponse{}, nil
}

// Add add a TemplateMacro.
func (n *TemplateMacrosService) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return &AddResponse{}, nil
}

// Update update a TemplateMacro.
func (n *TemplateMacrosService) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return &UpdateResponse{}, nil
}

// Delete delete a TemplateMacro.
func (n *TemplateMacrosService) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return &DeleteResponse{}, nil
}

// Watch Watch TemplateMacros.
func (n *TemplateMacrosService) Watch(*WatchRequest, TemplateMacros_WatchServer) error {
	return nil
}
