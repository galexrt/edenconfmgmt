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

package cache

// ResultState response state type.
type ResultState string

const (
	// ResultStateCreated created response state.
	ResultStateCreated ResultState = "Created"
	// ResultStateUpdated updated response state.
	ResultStateUpdated ResultState = "Updated"
	// ResultStateDeleted deleted response state.
	ResultStateDeleted ResultState = "Deleted"
	// ResultStateUnknown unknown response state.
	ResultStateUnknown ResultState = "Unknown"
)

// InformerResult result returned by the channel
type InformerResult struct {
	Closed  bool
	Errors  []error
	State   ResultState
	Key     string
	Value   []byte
	Version int64
}
