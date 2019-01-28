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

package api

import (
	"path"
)

const (
	dataStoreNamespacesPath = "/%s/%s/%s"
)

// GetObjectPath return the correct object path for the given values.
func GetObjectPath(dataStoreAPIPath string, namespace string, objectName string) string {
	return NamespacedObjectPath(dataStoreAPIPath, namespace, objectName)
}

// ObjectPath return data store path for a non namespaced object (DEPRECATED).
func ObjectPath(dataStoreAPIPath string, objectName string) string {
	return NamespacedObjectPath(dataStoreAPIPath, "", objectName)
}

// NamespacedObjectPath return data store  path for a namespaced object.
func NamespacedObjectPath(dataStoreAPIPath string, namespace string, objectName string) string {
	return path.Join("/", dataStoreAPIPath, namespace, objectName)
}
