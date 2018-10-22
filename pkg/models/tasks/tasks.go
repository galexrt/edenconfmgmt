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

package tasks

import (
	"github.com/galexrt/edenconfmgmt/pkg/models/common"
)

// Task
type Task struct {
	RunOptions *RunOptions `yaml:"runOptions,omitempty"`
	Steps      []Step      `yaml:"steps"`
}

// RunOptions
type RunOptions struct {
	Sync      bool       `yaml:"sync,omitempty"`
	Serialize *Serialize `yaml:"serialize,omitempty"`
	Target    *Target    `yaml:"target,omitempty"`
}

// Target
type Target struct {
	Hosts []string `yaml:"hosts"`
	Limit int64    `yaml:"limit"`
}

// Serialize
type Serialize struct {
	Count               int64 `yaml:"count"`
	FailAfterAllTargets bool  `yaml:"failAfterAllTargets"`
}

// Step
type Step map[string]Action

// Action
type Action struct {
	RunOptions *RunOptions              `yaml:"runOptions"`
	Conditions *common.Conditions       `yaml:"conditions"`
	Action     string                   `yaml:"action"`
	Parameters []map[string]interface{} `yaml:"parameters"`
}
