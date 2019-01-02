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

package utilsapi

import (
	core_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
)

// ErrorToErrorResponse translates a Golang error to an error core object
func ErrorToErrorResponse(err error) *core_v1alpha.Error {
	message := ""
	if err != nil {
		message = err.Error()
	}
	// TODO Try to create an useful code for the response.
	return &core_v1alpha.Error{
		Code:    -1,
		Message: message,
	}
}
