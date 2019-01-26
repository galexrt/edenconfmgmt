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

package v1alpha

// SetDefaults set defaults on object (called through grpc plugin magic code generation).
func (o *Node) SetDefaults() error {
	// TODO Set defaults for object.
	o.Spec.Network = &Network{
		Fqdn: "YOLO",
	}
	return nil
}
