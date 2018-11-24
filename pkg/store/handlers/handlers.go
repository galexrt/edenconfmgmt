/*
Copyright 2018 Alexander Trost. All rights reserved.

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

package handlers

import (
	"fmt"

	"github.com/galexrt/edenconfmgmt/pkg/store"
	"github.com/spf13/cobra"
)

var (
	// flagRegisters list of available store handlers flag register functions.
	flagRegisters = map[string]func(cmd *cobra.Command){}

	// handlers list of available store handlers.
	handlers = map[string]func() (store.Store, error){}
)

// RegisterFlags register flags using the function provided in flagRegisters var.
func RegisterFlags(cmd *cobra.Command) {
	for _, register := range flagRegisters {
		register(cmd)
	}
}

// Get return a newly created store handler.
func Get(handlerName string) (store.Store, error) {
	newFunc, ok := handlers[handlerName]
	if !ok {
		return nil, fmt.Errorf("no handler with name %s found in handlers list", handlerName)
	}
	return newFunc()
}
