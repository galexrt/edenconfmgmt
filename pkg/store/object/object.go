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
	"fmt"
	"sync"

	core_v1 "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	"github.com/galexrt/edenconfmgmt/pkg/store/cache"
	"go.uber.org/zap"
)

// Store object to hold objects and info about these objects.
type Store struct {
	cache  *cache.Store
	wg     sync.WaitGroup
	logger *zap.Logger
	apis   map[string]*apiInfo
}

type apiInfo struct {
	cache  map[string][]byte
	names  map[string]string
	labels map[string]string
}

// New return a new object Store
func New(store *cache.Store, logger *zap.Logger) *Store {
	return &Store{
		cache:  store,
		logger: logger.Named("pkg/store/object:Store"),
	}
}

// Init
func (s *Store) Init() error {
	// TODO Should we "init" the object caches by:
	// a) getting all objects from the s.cache
	// b) start a watch for each api registered
	for k, v := range s.apis {
		err := s.warmupAPICache(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) warmupAPICache(api string, info *apiInfo) error {
	_, err := s.cache.List(context.Background(), "nope")

	return err
}

// Start
func (s *Store) Start(stopCh chan struct{}) error {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.cache.Start(stopCh)
		s.cache.Close()
	}()
	s.wg.Done()
	return nil
}

// Register register an API
func (s *Store) Register(name string) error {
	s.apis[name] = &apiInfo{}
	return nil
}

// Get
func (s *Store) Get(ctx context.Context, opts *core_v1.GetOptions) ([]byte, error) {
	namespace := opts.Namespace
	_ = namespace
	if opts.Name != "" {
		resp, err := s.cache.Get(ctx, "/"+opts.Name)
		if err != nil {
			return nil, err
		}
		return resp, nil
	} else if opts.LabelSelectors != nil {
		//
	}
	return nil, fmt.Errorf("no name nor label selector set")
}

// List
func (s *Store) List(ctx context.Context, opts *core_v1.ListOptions) ([][]byte, error) {
	return nil, nil
}

// Create
func (s *Store) Create(ctx context.Context, obj []byte, opts *core_v1.CreateOptions) error {
	return s.cache.Put(ctx, "/my-cool-node", obj)
}

// Update
func (s *Store) Update(ctx context.Context, obj []byte, opts *core_v1.UpdateOptions) error {
	return nil
}

// Delete
func (s *Store) Delete(ctx context.Context, opts *core_v1.DeleteOptions) error {
	return nil
}

// Watch
func (s *Store) Watch(ctx context.Context, opts *core_v1.WatchOptions) (chan *InformerResult, error) {
	return nil, nil
}
