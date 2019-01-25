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

package object

import (
	"context"

	core_v1 "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	"github.com/galexrt/edenconfmgmt/pkg/store/cache"
)

// Store object to hold objects and info about these objects.
type Store struct {
	store *cache.Store
	apis  map[string]*apiInfo
}

type apiInfo struct {
	name   map[string]string
	labels map[string]string
}

// New return a new object Store
func New(store *cache.Store) *Store {
	return &Store{
		store: store,
	}
}

// Get
func (s *Store) Get(ctx context.Context, opts *core_v1.GetOptions) (interface{}, error) {
	//namespace := opts.Namespace
	if opts.Name != "" {
		//
	} else if opts.LabelSelectors != nil {

	}
	return nil, nil
}

// List
func (s *Store) List(ctx context.Context, opts *core_v1.ListOptions) ([]interface{}, error) {
	return nil, nil
}

// Create
func (s *Store) Create(ctx context.Context, opts *core_v1.CreateOptions) (interface{}, error) {
	return nil, nil
}

// Update
func (s *Store) Update(ctx context.Context, opts *core_v1.UpdateOptions) (interface{}, error) {
	return nil, nil
}

// Delete
func (s *Store) Delete(ctx context.Context, opts *core_v1.DeleteOptions) error {
	return nil
}

// Watch
func (s *Store) Watch(ctx context.Context, opts *core_v1.WatchOptions) (chan *InformerResult, error) {
	return nil, nil
}
