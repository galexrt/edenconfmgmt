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
	"github.com/galexrt/edenconfmgmt/pkg/utils/api"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"
)

// MetadataUnmarshalerAndMarshaler GetMetadata func proto.Unmarshaler and proto.Marshaler in one interface
type MetadataUnmarshalerAndMarshaler interface {
	GetMetadata() *core_v1.ObjectMetadata
	proto.Marshaler
	proto.Unmarshaler
}

// Store object to hold objects and info about these objects.
type Store struct {
	prefix string
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
		prefix: "",
	}
}

// Prefixed return a prefixed Store instance.
func (s *Store) Prefixed(prefix string) *Store {
	return &Store{
		cache:  s.cache,
		logger: s.logger.Named("pkg/store/object:Store"),
		prefix: prefix,
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
		resp, err := s.cache.Get(ctx, api.GetObjectPath(s.prefix, opts.GetNamespace(), opts.GetName()))
		if err != nil {
			return nil, err
		}
		if len(resp) == 0 {
			return nil, ErrNotFound
		}
		return resp, nil
	} else if opts.LabelSelectors != nil {
		// TODO Implement label selector
	}
	return nil, fmt.Errorf("no name nor label selector set")
}

// List
func (s *Store) List(ctx context.Context, opts *core_v1.ListOptions) ([][]byte, error) {
	if len(opts.LabelSelectors) == 0 {
		resp, err := s.cache.List(ctx, api.GetObjectPath(s.prefix, opts.GetNamespace(), opts.GetName()))
		if err != nil {
			return nil, err
		}
		if len(resp) == 0 {
			return nil, ErrNotFound
		}
		ret := make([][]byte, len(resp))
		for k := range resp {
			ret = append(ret, resp[k])
		}
		return ret, nil
	} else {
		// TODO Implement label selector
	}
	return nil, nil
}

// Create
func (s *Store) Create(ctx context.Context, obj MetadataUnmarshalerAndMarshaler, opts *core_v1.CreateOptions) error {
	// TODO First get and then put if it does not exist.
	_, err := s.Get(ctx, &core_v1.GetOptions{
		Name:      obj.GetMetadata().GetName(),
		Namespace: obj.GetMetadata().GetNamespace(),
	})
	if err != nil && err != ErrNotFound {
		return err
	}
	out, err := obj.Marshal()
	if err != nil {
		return nil
	}
	return s.cache.Put(ctx, api.GetObjectPath(s.prefix, obj.GetMetadata().GetNamespace(), obj.GetMetadata().GetName()), out)
}

// Update
func (s *Store) Update(ctx context.Context, obj MetadataUnmarshalerAndMarshaler, opts *core_v1.UpdateOptions) error {
	out, err := obj.Marshal()
	if err != nil {
		return nil
	}
	return s.cache.Put(ctx, api.GetObjectPath(s.prefix, obj.GetMetadata().GetNamespace(), obj.GetMetadata().GetName()), out)
}

// Delete
func (s *Store) Delete(ctx context.Context, obj MetadataUnmarshalerAndMarshaler, opts *core_v1.DeleteOptions) error {
	return s.cache.Delete(ctx, api.GetObjectPath(s.prefix, obj.GetMetadata().GetNamespace(), obj.GetMetadata().GetName()))
}

// Watch
func (s *Store) Watch(ctx context.Context, opts *core_v1.WatchOptions) (chan *cache.InformerResult, error) {
	return s.cache.Watch(ctx, api.GetObjectPath(s.prefix, opts.GetNamespace(), opts.GetName()))
}
