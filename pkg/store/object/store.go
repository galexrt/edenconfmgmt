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
	"sync"

	core_v1 "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	"github.com/galexrt/edenconfmgmt/pkg/store/cache"
	storecommon "github.com/galexrt/edenconfmgmt/pkg/store/common"
	"github.com/galexrt/edenconfmgmt/pkg/utils/api"
	"go.uber.org/zap"
)

// Store object to hold objects and info about these objects.
type Store struct {
	prefix   string
	cache    *cache.Store
	informer *Informer
	wg       sync.WaitGroup
	logger   *zap.Logger
}

type apiInfo struct {
	cache  map[string][]byte
	names  map[string]string
	labels map[string]string
}

// New return a new object Store
func New(store *cache.Store, informer *Informer, logger *zap.Logger) *Store {
	return &Store{
		cache:    store,
		informer: informer,
		logger:   logger.Named("pkg/store/object:Store"),
		prefix:   "",
	}
}

// Prefixed return a prefixed Store instance.
func (s *Store) Prefixed(prefix string) *Store {
	return &Store{
		cache:    s.cache,
		informer: s.informer,
		logger:   s.logger.Named("pkg/store/object:Store"),
		prefix:   prefix,
	}
}

// Start start the object store logic loops (compcation, cleanup).
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

// Get get an object.
func (s *Store) Get(ctx context.Context, opts *core_v1.GetOptions) ([]byte, error) {
	namespace := opts.Namespace
	_ = namespace
	if opts.Name != "" {
		resp, err := s.cache.Get(ctx, s.getObjectPathFromOpts(opts))
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
	return nil, ErrNoSelector
}

// List list objects.
func (s *Store) List(ctx context.Context, opts *core_v1.ListOptions) ([][]byte, error) {
	if len(opts.LabelSelectors) == 0 {
		resp, err := s.cache.List(ctx, s.getObjectPathFromOpts(opts))
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
	}
	// TODO Implement label selector
	return nil, nil
}

// Create create object, errors if one already exists.
func (s *Store) Create(ctx context.Context, obj MetadataUnmarshalerAndMarshaler, opts *core_v1.CreateOptions) error {
	res, err := s.Get(ctx, &core_v1.GetOptions{
		Name:      obj.GetMetadata().GetName(),
		Namespace: obj.GetMetadata().GetNamespace(),
	})
	if err != nil && err != ErrNotFound {
		if len(res) > 0 {
			return ErrAlreadyExists
		}
		return err
	}
	out, err := obj.Marshal()
	if err != nil {
		return nil
	}
	return s.cache.Put(ctx, s.getObjectPathFromObjAndOpts(obj, opts), out)
}

// Update update an object.
func (s *Store) Update(ctx context.Context, obj MetadataUnmarshalerAndMarshaler, opts *core_v1.UpdateOptions) error {
	out, err := obj.Marshal()
	if err != nil {
		return nil
	}
	return s.cache.Put(ctx, s.getObjectPathFromObjAndOpts(obj, opts), out)
}

// Delete delete an object.
func (s *Store) Delete(ctx context.Context, obj MetadataUnmarshalerAndMarshaler, opts *core_v1.DeleteOptions) error {
	return s.cache.Delete(ctx, s.getObjectPathFromObjAndOpts(obj, opts))
}

// Watch watch an object.
func (s *Store) Watch(ctx context.Context, opts *core_v1.WatchOptions) (chan *storecommon.InformerResult, error) {
	if opts.GetName() != "" {
		return s.informer.WatchOnName(ctx, s.getObjectPathFromOpts(opts))
	}
	return s.informer.WatchOnLabels(ctx, opts.LabelSelectors)
}

func (s *Store) getObjectPathFromOpts(opts GetNameAndNamespaceOptions) string {
	return api.GetObjectPath(s.prefix, opts.GetNamespace(), opts.GetName())
}

func (s *Store) getObjectPathFromObjAndOpts(obj MetadataUnmarshalerAndMarshaler, opts GetNameAndNamespaceOptions) string {
	var name, namespace string
	if obj.GetMetadata() != nil {
		namespace = obj.GetMetadata().GetNamespace()
		name = obj.GetMetadata().GetName()
	} else {
		namespace = opts.GetNamespace()
		name = opts.GetName()
	}
	return api.GetObjectPath(s.prefix, namespace, name)
}
