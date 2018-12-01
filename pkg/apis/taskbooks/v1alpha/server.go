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

// TaskBookBooksService handler for config events.
type TaskBookBooksService struct {
	TaskBookBooksServer
}

// New returns a TaskBookBooksService
func New() TaskBookBooksServer {
	return &TaskBookBooksService{}
}

// Get get a TaskBook.
func (n *TaskBookBooksService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return &GetResponse{}, nil
}

// List list TaskBookBooks.
func (n *TaskBookBooksService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return &ListResponse{}, nil
}

// Add add a TaskBook.
func (n *TaskBookBooksService) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return &AddResponse{}, nil
}

// Update update a TaskBook.
func (n *TaskBookBooksService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return &UpdateResponse{}, nil
}

// Delete delete a TaskBook.
func (n *TaskBookBooksService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return &DeleteResponse{}, nil
}

// Watch Watch TaskBookBooks.
func (n *TaskBookBooksService) Watch(req *WatchRequest, watch TaskBookBooks_WatchServer) error {
	return nil
}
