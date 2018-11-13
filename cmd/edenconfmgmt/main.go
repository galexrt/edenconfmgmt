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
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	configs_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha"
	eventreactors_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/eventreactors/v1alpha"
	events_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	jobs_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/jobs/v1alpha"
	nodes_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha"
	tasks_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/tasks/v1alpha"
	templatemacros_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/templatemacros/v1alpha"
	variables_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha"
	"github.com/galexrt/edenconfmgmt/pkg/auth"
	"github.com/galexrt/edenconfmgmt/pkg/common"
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
	flagProductionMode    = "production"
	flagListenAddressGRPC = "listen-address-grpc"
	flagListenAddressHTTP = "listen-address-http"
	flagNeighbors         = "neighbors"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "edenconfmgmt",
		Short: "Configuration management with automatic clustering, events and stuff.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger = common.GetLogger()
		},
		RunE: Run,
	}
	// Logger for logging.
	logger *zap.Logger
	// Prometheus metrics registry.
	promReg = prometheus.NewRegistry()
	// grpc prometheus server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()
)

func init() {
	promReg.MustRegister(grpcMetrics)

	rootCmd.PersistentFlags().Bool(flagProductionMode, true, "production mode (logger will use different output format)")
	rootCmd.PersistentFlags().String(flagListenAddressGRPC, ":1337", "grpc listen address")
	rootCmd.PersistentFlags().String(flagListenAddressHTTP, ":1338", "http listen address")
	rootCmd.PersistentFlags().StringSlice(flagNeighbors, []string{}, "list of other neighbors (one neighbor per flag)")
	viper.BindPFlag(flagProductionMode, rootCmd.PersistentFlags().Lookup(flagProductionMode))
	viper.BindPFlag(flagListenAddressGRPC, rootCmd.PersistentFlags().Lookup(flagListenAddressGRPC))
	viper.BindPFlag(flagListenAddressHTTP, rootCmd.PersistentFlags().Lookup(flagListenAddressHTTP))
	viper.SetDefault(flagProductionMode, true)
	viper.SetDefault(flagListenAddressGRPC, ":1337")
	viper.SetDefault(flagListenAddressHTTP, ":1338")
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
		logger.Fatal("failed to execute root cmd", zap.Error(err))
		os.Exit(1)
	}
	os.Exit(0)
}

// Run runs all routines to start the work of edenconfmgmt application.
func Run(cmd *cobra.Command, args []string) error {
	defer logger.Sync()
	logger.Info("starting", zap.String("command", os.Args[0]), zap.String("versionInfo", version.Info()))

	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// TODO Implement TLS
	lis, err := net.Listen("tcp", viper.GetString(flagListenAddressGRPC))
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
	authProvider, err := auth.New()
	if err != nil {
		logger.Fatal("failed to init auth", zap.Error(err))
	}

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpcMetrics.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			authProvider.StreamInterceptor,
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpcMetrics.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			authProvider.UnaryInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}
	grpcServer := grpc.NewServer(opts...)
	grpc_zap.ReplaceGrpcLogger(logger)

	// Register APIs to grpc server
	registerGRPCAPIs(grpcServer)

	// Initialize Prometheus GRPC metrics.
	grpcMetrics.InitializeMetrics(grpcServer)

	// Create a HTTP server for prometheus.
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(promReg, promhttp.HandlerOpts{}),
		Addr:    viper.GetString(flagListenAddressHTTP),
	}

	wg := sync.WaitGroup{}

	// TODO Create a "mesh" proxy.
	// Would be cool if each daemon can tell other servers where the masters are or even proxy to them.

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
	// Configs
	configServer := configs_v1alpha.New()
	configs_v1alpha.RegisterConfigsServer(srv, configServer)
	// EventReactors
	eventReactorsServer := eventreactors_v1alpha.New()
	eventreactors_v1alpha.RegisterEventReactorsServer(srv, eventReactorsServer)
	// Events
	eventsServer := events_v1alpha.New()
	events_v1alpha.RegisterEventsServer(srv, eventsServer)
	// Jobs
	jobsServer := jobs_v1alpha.New()
	jobs_v1alpha.RegisterJobsServer(srv, jobsServer)
	// Nodes
	nodesServer := nodes_v1alpha.New()
	nodes_v1alpha.RegisterNodesServer(srv, nodesServer)
	// Tasks
	tasksServer := tasks_v1alpha.New()
	tasks_v1alpha.RegisterTasksServer(srv, tasksServer)
	// TemplateMacros
	templateMacrosServer := templatemacros_v1alpha.New()
	templatemacros_v1alpha.RegisterTemplateMacrosServer(srv, templateMacrosServer)
	// Variables
	variablesServer := variables_v1alpha.New()
	variables_v1alpha.RegisterVariablesServer(srv, variablesServer)
}
