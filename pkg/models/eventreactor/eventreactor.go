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

package eventreactor

import (
	"time"

	"github.com/galexrt/edenconfmgmt/pkg/models/common"
	"github.com/galexrt/edenconfmgmt/pkg/models/jobs"
)

// EventReactors list of EventReactor.
type EventReactors map[string]EventReactor

// EventReactor Reaction to certain events.
type EventReactor struct {
	Reaction       Reaction          `yaml:"reaction"`
	TriggerLimiter TriggerLimiter    `yaml:"triggerLimiter"`
	Conditions     common.Conditions `yaml:"conditions"`
}

// Reaction currently only jobs can be triggered when an event occurs.
type Reaction struct {
	Jobs jobs.Jobs `yaml:"jobs"`
	// TODO Allow other events to be triggered her.
	//Events events.Events `yaml:"events"`
}

// TriggerLimiter trigger limiter options.
type TriggerLimiter struct {
	MinOccurences int64         `yaml:"minOccurences"`
	ReReactDelay  time.Duration `yaml:"reReactDelay"`
	WaitForMore   WaitForMore   `yaml:"waitForMore"`
}

// WaitForMore trigger limiter WaitForMore options.
type WaitForMore struct {
	Timeout            time.Duration `yaml:"timeout"`
	ResetTimeoutOnMore bool          `yaml:"resetTimeoutOnMore"`
}
