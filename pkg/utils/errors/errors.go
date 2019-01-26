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

package errors

import (
	"fmt"
	"strings"
)

// Concat concat multiple errors into one error (separated by semicolon).
func Concat(errors ...error) error {
	if len(errors) > 0 {
		errsConcat := "failed to close datastore and cachestore: "
		onlyNullErrs := true
		for _, e := range errors {
			if e == nil {
				continue
			}
			onlyNullErrs = false
			errsConcat += e.Error()
			errsConcat += "; "
		}
		if onlyNullErrs {
			return nil
		}
		return fmt.Errorf(strings.TrimRight(errsConcat, "; "))
	}
	return nil
}
