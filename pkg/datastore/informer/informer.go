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

// State response state type.
type State string

const (
	// StateCreated created response state.
	StateCreated State = "Created"
	// StateUpdated updated response state.
	StateUpdated State = "Updated"
	// StateDeleted deleted response state.
	StateDeleted State = "Deleted"
	// StateUnknown unknown response state.
	StateUnknown State = "Unknown"
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

// Result result returned by the channel
type Result struct {
	Closed  bool
	Errors  []error
	State   State
	Name    string
	Value   string
	Version int64
}

// InformerMiddleware middleware function to be used with Informer
type InformerMiddleware func(*Result) (bool, error)

// Informer holds information and middleware to run on `Result` coming from a watch channel (SharedInformer manages the watch).
type Informer struct {
	input       chan *Result
	output      chan *Result
	Type        InformerType
	Path        string
	middleware  []InformerMiddleware
	sharedClose SharedInformerClose
}

// NewInformer return a new Informer.
func NewInformer(kp *SharedInformer, input chan *Result, infType InformerType, path string, middleware ...InformerMiddleware) *Informer {
	return &Informer{
		input:      input,
		output:     make(chan *Result),
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

// WatchChan return the watch channel which can/must be used to receive the `Result`s of the watch.
// Start() must have been called before this otherwise there will never be output.
func (inf *Informer) WatchChan() <-chan *Result {
	return inf.output
}

// Close the Informer's watch (channmel).
func (inf *Informer) close() {
	inf.sharedClose(inf.Path)
	close(inf.input)
	close(inf.output)
}
