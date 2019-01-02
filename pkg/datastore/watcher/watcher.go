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

package watcher

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

// Result result returned by the channel
type Result struct {
	Closed bool
	Errors []error
	State  State
	Name   string
	Value  string
}

// Keeper registration and "time" keeper over watches.
type Keeper struct {
}

// Register register a new watch channel.
func (kp *Keeper) Register() {
	return
}

/* =============================================================================
TODO
* Keeper will create the chan *Result
* Keeper will "compact" watches on the same (sub-)path(s) automatically. Example:
	* **2** watches: 1. `/example/apiname/kindA123/objectabc`, 2. `/example/apiname/nopethisisatest/xyz123123`
	* => Results in: **ONE** recursive watch on `/example/apiname`.
	* Advantages: Less watches on the datastore.
	* Disadvantage: This will probably be a bottle neck at some point as incoming Results will
		for the first time be single threadely accepted (future: add (global) worker count?).
============================================================================= */

// Informer holds information and middleware to run on `Result` coming from a watch channel (Keeper manages the watch).
type Informer struct {
	input      chan *Result
	output     chan *Result
	Type       InformerType
	Path       string
	Middleware []InformerMiddleware
}

// NewInformer return a new Informer.
func NewInformer(input chan *Result, infType InformerType, path string, middleware ...InformerMiddleware) *Informer {
	return &Informer{
		input:      input,
		output:     make(chan *Result),
		Type:       infType,
		Path:       path,
		Middleware: middleware,
	}
}

// InformerMiddleware middleware function to be used with Informer
type InformerMiddleware func(*Result) (bool, error)

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

// Start start watch on input channel and push them to the output channel.
func (inf *Informer) Start(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case in := <-inf.input:
			inf.runInformerMiddleware(in)
			inf.output <- in
		}
	}
}

func (inf *Informer) runInformerMiddleware(result *Result) {
	for _, mware := range inf.Middleware {
		done, err := mware(result)
		if err != nil {
			result.Errors = append(result.Errors, err)
		}
		if done {
			return
		}
	}
}

// WatchChan return the watch channel which can/must be used to receive the `Result`s of the watch.
func (inf *Informer) WatchChan() <-chan *Result {
	return inf.output
}

// Close the Informer's watch (channmel).
func (inf *Informer) Close() error {
	return nil
}
