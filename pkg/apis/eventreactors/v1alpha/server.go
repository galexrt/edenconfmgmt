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

// EventReactorsService handler for config events.
type EventReactorsService struct {
	EventReactorsServiceServer
}

// NewServer returns a EventReactorsService
func NewServer() EventReactorsServiceServer {
	return &EventReactorsService{}
}

// Get get a EventReactor.
func (n *EventReactorsService) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return &GetResponse{}, nil
}

// List list EventReactors.
func (n *EventReactorsService) List(context.Context, *ListRequest) (*ListResponse, error) {
	return &ListResponse{}, nil
}

// Add add a EventReactor.
func (n *EventReactorsService) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return &AddResponse{}, nil
}

// Update update a EventReactor.
func (n *EventReactorsService) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return &UpdateResponse{}, nil
}

// Delete delete a EventReactor.
func (n *EventReactorsService) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return &DeleteResponse{}, nil
}

// Watch Watch EventReactors.
func (n *EventReactorsService) Watch(*WatchRequest, EventReactorsService_WatchServer) error {
	return nil
}
