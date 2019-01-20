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

package adapters

import (
	"fmt"

	"github.com/galexrt/edenconfmgmt/pkg/store/data"
	"github.com/spf13/cobra"
)

var (
	// flagRegisters list of available data adapters flag register functions.
	flagRegisters = map[string]func(prefix string, cmd *cobra.Command){}

	// adapters list of available data adapters.
	adapters = map[string]func(flagPrefix string) (data.Store, error){}
)

// RegisterFlags register flags using the function provided in flagRegisters var.
func RegisterFlags(prefix string, cmd *cobra.Command) {
	if prefix != "" {
		prefix = prefix + "-"
	}
	for _, register := range flagRegisters {
		register(prefix, cmd)
	}
}

// Get return a newly created data handler.
func Get(handlerName string, flagPrefix string) (data.Store, error) {
	newFunc, ok := adapters[handlerName]
	if !ok {
		return nil, fmt.Errorf("no handler with name %s found in adapters list", handlerName)
	}
	return newFunc(flagPrefix)
}
