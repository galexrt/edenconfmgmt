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

package apis

import (
	"github.com/galexrt/edenrun/pkg/store/object"
	"google.golang.org/grpc"
	// TODO Add APIs here and add API Path const to apis trigger.go, runny/runny.go and so on.
)

var (
	APIs = map[string]func(srv *grpc.Server, objectStore *object.Store, informer *object.Informer){}
)

// Register
func Register(srv *grpc.Server, objectStore *object.Store, informer *object.Informer) {
	for path, api := range APIs {
		api(srv, objectStore.Prefixed(path), informer)
	}
}
