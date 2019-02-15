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
	core_v1 "github.com/galexrt/edenrun/pkg/apis/eden.run/core/meta/v1"
	"github.com/gogo/protobuf/proto"
)

// MetadataUnmarshalerAndMarshaler GetMetadata() proto.Unmarshaler and proto.Marshaler in an interface.
type MetadataUnmarshalerAndMarshaler interface {
	GetMetadata() *core_v1.ObjectMetadata
	proto.Marshaler
	proto.Unmarshaler
}

// GetNameAndNamespaceOptions GetName() and GetNamespace() in an interface.
type GetNameAndNamespaceOptions interface {
	GetName() string
	GetNamespace() string
}
