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

package tracer

import (
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
)

// Init inits the Tracer using the provided configuration
func Init() (opentracing.Tracer, io.Closer, error) {
	var tracer opentracing.Tracer
	var err error
	var closer io.Closer
	if false {
		tracer, closer, err = initJaeger("edenrun")
		if err != nil {
			return nil, nil, fmt.Errorf("failed to init jaeger. %+v", err)
		}
	} else {
		tracer = opentracing.NoopTracer{}
	}

	return tracer, closer, nil
}
