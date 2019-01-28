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

import "fmt"

var (
	// ErrNotFound error that is returned when an object is not found.
	ErrNotFound = fmt.Errorf("object not found")
	// ErrAlreadyExists error that an object with that name already exists.
	ErrAlreadyExists = fmt.Errorf("object with name already exists")
	// ErrNoSelector no name nor label selector had been given.
	ErrNoSelector = fmt.Errorf("no name nor label selector set")
	// ErrNoLabelsInformer no labels were in the selector when given to the `object.Informer`.
	ErrNoLabelsInformer = fmt.Errorf("no labels in selector (object.Informer)")
)
