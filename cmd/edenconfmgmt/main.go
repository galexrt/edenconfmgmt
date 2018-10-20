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
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

type cmdLineOpts struct {
	neighbors []string
}

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "edenconfmgmt",
		Short: "Configuration management with automatic clustering, events and stuff.",
		RunE:  Run,
	}

	opts cmdLineOpts

	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
)

func init() {
	zerolog.TimeFieldFormat = ""
	//rootCmd.PersistentFlags().StringSliceVarP(&opts.neighbors, "neighbors", "n", []string{}, "comma separated list of other neighbors")
}

func main() {
	Execute()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal().Err(err)
		os.Exit(1)
	}
	os.Exit(0)
}

// Run runs all routines to start the work of edenconfmgmt application.
func Run(cmd *cobra.Command, args []string) error {
	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}

	/*wg.Add(1)
	go func() {
		defer wg.Done()
		if err := nbs.Run(stopCh); err != nil {
			logger.Error(err)
		}
	}()*/

	<-sigCh
	close(stopCh)
	wg.Wait()
	return nil
}
