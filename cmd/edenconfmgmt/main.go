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
	"github.com/galexrt/edenconfmgmt/pkg/auth"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaeger_prometheus "github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
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
	logger *zap.Logger
)

func init() {
	rootCmd.PersistentFlags().String(listenAddressGRPC, ":1337", "grpc listen address")
	rootCmd.PersistentFlags().String(listenAddressHTTP, ":1338", "http listen address")
	//rootCmd.PersistentFlags().StringSliceVarP(&cmdOpts.neighbors, "neighbors", "n", []string{}, "comma separated list of other neighbors")
	viper.BindPFlag(listenAddressGRPC, rootCmd.PersistentFlags().Lookup(listenAddressGRPC))
	viper.BindPFlag(listenAddressHTTP, rootCmd.PersistentFlags().Lookup(listenAddressHTTP))
	viper.SetDefault(listenAddressGRPC, ":1337")
	viper.SetDefault(listenAddressHTTP, ":1338")
}

func main() {
	var err error
	//logger, err = zap.NewProduction()
	logger, err = zap.NewDevelopment()
	if err != nil {
		fmt.Printf("failed to create zap logger. %+v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()
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
		logger.Fatal("failed to execute root cmd", zap.Error(err))
		os.Exit(1)
	}
	os.Exit(0)
}

// Run runs all routines to start the work of edenconfmgmt application.
func Run(cmd *cobra.Command, args []string) error {
	logger.Info("starting", zap.String("command", os.Args[0]), zap.String("versionInfo", version.Info()))

	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	lis, err := net.Listen("tcp", viper.GetString(listenAddressGRPC))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	var tracer opentracing.Tracer
	if false {
		var closer io.Closer
		tracer, closer, err = initJaeger("edenconfmgmt")
		if err != nil {
			logger.Fatal("failed to init jaeger", zap.Error(err))
		}
		defer closer.Close()
	} else {
		tracer = opentracing.NoopTracer{}
	}

	// grpc_opentracing automatically uses the global tracer so we set it.
	opentracing.SetGlobalTracer(tracer)

	// Auth
	authProvider := auth.New()

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			authProvider.StreamInterceptor,
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			authProvider.UnaryInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}
	grpcServer := grpc.NewServer(opts...)
	grpc_zap.ReplaceGrpcLogger(logger)

	// Register APIs to grpc server
	registerGRPCAPIs(grpcServer)

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
			logger.Error("failed to listen and serve http server", zap.Error(err))
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = grpcServer.Serve(lis); err != nil {
			logger.Error("grpcServer.Serve() returned non-nil error on GracefulStop", zap.Error(err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := magicRun(stopCh); err != nil {
			logger.Error("magic run returned error", zap.Error(err))
		}
	}()

	<-sigCh
	logger.Info("signal received, shutting down ...")
	close(stopCh)

	// First shutdown grpc server
	grpcServer.GracefulStop()

	// Shutdown http server
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("http server graceful shutdown failed", zap.Error(err))
	}

	wg.Wait()
	return nil
}

func registerGRPCAPIs(srv *grpc.Server) {
	nodesServiceServer := nodes_v1alpha.NewServer()
	nodes_v1alpha.RegisterNodesServiceServer(srv, nodesServiceServer)
}
