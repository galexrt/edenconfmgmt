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
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	nodes_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaeger_prometheus "github.com/uber/jaeger-lib/metrics/prometheus"
	"google.golang.org/grpc"
)

const (
	listenAddressGRPC = "listen-address-grpc"
	listenAddressHTTP = "listen-address-http"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "edenconfmgmt",
		Short: "Configuration management with automatic clustering, events and stuff.",
		RunE:  Run,
	}
	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = ""
	rootCmd.PersistentFlags().String(listenAddressGRPC, ":1337", "grpc listen address")
	rootCmd.PersistentFlags().String(listenAddressHTTP, ":1338", "http listen address")
	//rootCmd.PersistentFlags().StringSliceVarP(&cmdOpts.neighbors, "neighbors", "n", []string{}, "comma separated list of other neighbors")
	viper.BindPFlag(listenAddressGRPC, rootCmd.PersistentFlags().Lookup(listenAddressGRPC))
	viper.BindPFlag(listenAddressHTTP, rootCmd.PersistentFlags().Lookup(listenAddressHTTP))
	viper.SetDefault(listenAddressGRPC, ":1337")
	viper.SetDefault(listenAddressHTTP, ":1338")
}

func main() {
	Execute()
}

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaeger(service string) (opentracing.Tracer, io.Closer, error) {
	metricsFactory := jaeger_prometheus.New()

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 0.01,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	return cfg.New(service, config.Logger(jaeger.StdLogger), config.Metrics(metricsFactory))
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
	logger.Info().Msgf("starting %s (%s)", os.Args[0], version.Info())

	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	lis, err := net.Listen("tcp", viper.GetString(listenAddressGRPC))
	if err != nil {
		logger.Fatal().Err(fmt.Errorf("failed to listen: %v", err))
	}

	var (
		tracer opentracing.Tracer
		closer io.Closer
	)
	if false {
		tracer, closer, err = initJaeger("edenconfmgmt")
		if err != nil {
			logger.Fatal().Err(err)
		}
	} else {
		tracer = opentracing.NoopTracer{}
	}

	opentracing.SetGlobalTracer(tracer)

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())),
		grpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(opentracing.GlobalTracer())),
	}

	grpcServer := grpc.NewServer(opts...)

	// Register APIs to grpc server
	registerAPIs(grpcServer)

	reg := prometheus.NewRegistry()
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	reg.MustRegister(grpcMetrics)
	grpcMetrics.InitializeMetrics(grpcServer)

	// Create a HTTP server for prometheus.
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		Addr:    viper.GetString(listenAddressHTTP),
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = httpServer.ListenAndServe(); err != nil {
			logger.Fatal().Err(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = grpcServer.Serve(lis); err != nil {
			logger.Fatal().Err(fmt.Errorf("Serve() returned non-nil error on GracefulStop: %v", err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := magicRun(stopCh); err != nil {
			logger.Error().Err(err)
		}
	}()

	<-sigCh
	logger.Info().Msg("signal received, shutting down ...")
	close(stopCh)

	// First shutdown grpc server
	grpcServer.GracefulStop()

	// Shutdown http server
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)

	wg.Wait()
	closer.Close()
	return nil
}

func registerAPIs(srv *grpc.Server) {
	nodesServiceServer := nodes_v1alpha.NewServer()
	nodes_v1alpha.RegisterNodesServiceServer(srv, nodesServiceServer)
}
