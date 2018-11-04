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

package main

import (
	"github.com/galexrt/edenconfmgmt/pkg/common"
	"github.com/spf13/cobra"
)

var (
	shutdownCmd = &cobra.Command{
		Use:   "shutdown",
		Short: "Trigger shutdown",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			common.GetLogger()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)
