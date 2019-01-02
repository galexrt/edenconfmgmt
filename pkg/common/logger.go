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

package common

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	flagProductionMode = "production"
)

var logger *zap.Logger

// GetLogger init logger depending on production mode.
func GetLogger(l *zapcore.Level) *zap.Logger {
	if logger != nil {
		return logger
	}
	var err error
	if viper.GetBool(flagProductionMode) {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		fmt.Printf("failed to create zap logger. %+v\n", err)
		os.Exit(1)
	}
	return logger
}
