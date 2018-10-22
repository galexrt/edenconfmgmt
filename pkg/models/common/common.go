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

package common

import (
	"time"
)

// Base base file structure. This covers the "Variables Base File"
type Base map[string][]File

// File structure for "including" one or more files
type File struct {
	Name  string   `yaml:"name"`
	Paths []string `yaml:"paths"`
}

// Conditions condition which are available to be used.
type Conditions struct {
	When   *Condition `yaml:"when"`
	Sucess *Condition `yaml:"success"`
}

// Condition contents of a single condition.
type Condition struct {
	Condition string `yaml:"condition"`
	Retry     *Retry `yaml:"retry"`
}

// Retry holds interval and limit to be able to retry an action as the user wants/needs it.
type Retry struct {
	Interval time.Duration `yaml:"interval"`
	Limit    int64         `yaml:"limit"`
}
