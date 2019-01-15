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

package informer

import (
	"context"
)

// InformerType the type a Informer can be.
type InformerType string

const (
	// InformerTypeUnknown TODO Add comments
	InformerTypeUnknown InformerType = "Unknown"
	// InformerTypeObject TODO Add comments
	InformerTypeObject InformerType = "Object"
	// InformerTypeDirectory TODO Add comments
	InformerTypeDirectory InformerType = "Directory"
)

// InformerMiddleware middleware function to be used with Informer
type InformerMiddleware func(*ResultState) (bool, error)

// Informer holds information and middleware to run on `ResultState` coming from a watch channel (SharedInformer manages the watch).
type Informer struct {
	input       chan *ResultState
	output      chan *ResultState
	Type        InformerType
	Path        string
	middleware  []InformerMiddleware
	sharedClose SharedInformerClose
}

// NewInformer return a new Informer.
func NewInformer(kp *SharedInformer, input chan *ResultState, infType InformerType, path string, middleware ...InformerMiddleware) *Informer {
	return &Informer{
		input:      input,
		output:     make(chan *ResultState),
		Type:       infType,
		Path:       path,
		middleware: middleware,
	}
}

// Start start watch for Informer and use context.Context to cancel it.
func (inf *Informer) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case in := <-inf.input:
			runMiddlewares(inf.middleware, in)
			inf.output <- in
		}
	}
}

// StartWithStopCh start watch for Informer and use stopCh to cancel it.
func (inf *Informer) StartWithStopCh(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case in := <-inf.input:
			runMiddlewares(inf.middleware, in)
			inf.output <- in
		}
	}
}

// WatchChan return the watch channel which can/must be used to receive the `ResultState`s of the watch.
// Start() must have been called before this otherwise there will never be output.
func (inf *Informer) WatchChan() <-chan *ResultState {
	return inf.output
}

// Close the Informer's watch (channmel).
func (inf *Informer) close() {
	inf.sharedClose(inf.Path)
	close(inf.input)
	close(inf.output)
}
