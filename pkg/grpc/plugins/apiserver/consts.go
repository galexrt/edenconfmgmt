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

package apiserver

// Method "name" for the calls that are "supported" by the apiserver grpc plugin.
type Method int

const (
	// MethodGet method "name" for `Get()`
	MethodGet Method = iota + 1
	// MethodList method "name" for `List()`
	MethodList
	// MethodCreate method "name" for `Create()`
	MethodCreate
	// MethodUpdate method "name" for `Update()`
	MethodUpdate
	// MethodDelete method "name" for `Delete()`
	MethodDelete
	// MethodWatch method "name" for `Watch()`
	MethodWatch
)
