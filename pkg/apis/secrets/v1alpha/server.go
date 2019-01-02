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
)

// SecretsService handler for config events.
type SecretsService struct {
	SecretsServer
}

// New returns a SecretsService
func New() SecretsServer {
	return &SecretsService{}
}

// Get get a Secret.
func (n *SecretsService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return &GetResponse{}, nil
}

// List list Secrets.
func (n *SecretsService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return &ListResponse{}, nil
}

// Add add a Secret.
func (n *SecretsService) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return &AddResponse{}, nil
}

// Update update a Secret.
func (n *SecretsService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return &UpdateResponse{}, nil
}

// Delete delete a Secret.
func (n *SecretsService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return &DeleteResponse{}, nil
}

// Watch Watch Secrets.
func (n *SecretsService) Watch(req *WatchRequest, watch Secrets_WatchServer) error {
	return nil
}
