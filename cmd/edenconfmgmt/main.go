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

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	configs_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/agentconfigs/v1alpha"
	beacons_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/beacons/v1alpha"
	clustervariables_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/clustervariables/v1alpha"
	core_v1 "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	cronjobs_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/cronjobs/v1alpha"
	events_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	nodes_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha"
	secrets_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/secrets/v1alpha"
	taskbooks_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/taskbooks/v1alpha"
	triggers_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/triggers/v1alpha"
	variables_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha"
	"github.com/galexrt/edenconfmgmt/pkg/store/cache"
	"github.com/galexrt/edenconfmgmt/pkg/utils"

	_ "net/http/pprof"

	"github.com/galexrt/edenconfmgmt/pkg/auth"
	"github.com/galexrt/edenconfmgmt/pkg/common"
	store_data_adapters "github.com/galexrt/edenconfmgmt/pkg/store/data/adapters"
	"github.com/galexrt/edenconfmgmt/pkg/store/object"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
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
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

const (
	flagProductionMode          = "production"
	flagLogLevel                = "log-level"
	flagListenAddressGRPC       = "listen-address-grpc"
	flagListenAddressHTTP       = "listen-address-http"
	flagDataStore               = "data-store"
	flagDataStoreKeyPrefix      = "data-store-key-prefix"
	flagCacheStore              = "cache-store"
	flagCacheStoreKeyPrefix     = "cache-store-key-prefix"
	flagStoreWatchCleanInterval = "store-watch-clean-interval"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "edenconfmgmt",
		Short: "Configuration management with automatic clustering, events and stuff.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// TODO Implement atomic/dynamic log level
			var level zapcore.Level
			if err := (&level).UnmarshalText([]byte(viper.GetString(flagLogLevel))); err != nil {
				return err
			}
			logger = common.GetLogger(&level)
			return nil
		},
		RunE: Run,
	}
	shutdownCmd = &cobra.Command{
		Use:   "shutdown",
		Short: "Trigger shutdown of edenconfmgmt daemon",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// TODO Implement atomic/dynamic log level
			var level zapcore.Level
			if err := (&level).UnmarshalText([]byte(viper.GetString(flagLogLevel))); err != nil {
				return err
			}
			logger = common.GetLogger(&level)
			return nil
		},
		RunE: Shutdown,
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
	rootCmd.PersistentFlags().String(flagLogLevel, "INFO", "log level to use for logging")
	rootCmd.PersistentFlags().String(flagListenAddressGRPC, ":1337", "grpc listen address")
	rootCmd.PersistentFlags().String(flagListenAddressHTTP, ":1338", "http listen address")
	// Directories
	// Store
	rootCmd.PersistentFlags().String(flagDataStore, "etcd", "datastore to use")
	rootCmd.PersistentFlags().String(flagDataStoreKeyPrefix, "/edendata", "key prefix to use in the datastore")
	rootCmd.PersistentFlags().String(flagCacheStore, "memory", "cachestore to use")
	rootCmd.PersistentFlags().String(flagCacheStoreKeyPrefix, "", "key prefix to use in the cachestore")
	rootCmd.PersistentFlags().Duration(flagStoreWatchCleanInterval, 7*time.Second, "interval to remove (and possibly close) unused store wataches")

	// Register all datatstore adapters flags for cache and data store.
	store_data_adapters.RegisterFlags(flagDataStore, rootCmd)
	store_data_adapters.RegisterFlags(flagCacheStore, rootCmd)

	// Bind all persistent flags to viper, for easy access.
	viper.BindPFlags(rootCmd.PersistentFlags())
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

	dataStoreHandlerName := strings.ToLower(viper.GetString(flagDataStore))
	dataStoreAdapter, err := store_data_adapters.Get(dataStoreHandlerName, flagDataStore+"-")
	if err != nil {
		logger.Fatal("failed to create store handler", zap.String("storehandler", dataStoreHandlerName), zap.Error(err))
	}
	cacheStoreHandlerName := strings.ToLower(viper.GetString(flagCacheStore))
	cacheStoreAdapter, err := store_data_adapters.Get(cacheStoreHandlerName, flagCacheStore+"-")
	if err != nil {
		logger.Fatal("failed to create store handler", zap.String("storehandler", cacheStoreHandlerName), zap.Error(err))
	}
	cacheStore := cache.New(dataStoreAdapter, cacheStoreAdapter, logger)
	objectStore := object.New(cacheStore, logger)

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
			grpc_validator.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpcMetrics.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			authProvider.UnaryInterceptor,
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}
	grpcServer := grpc.NewServer(opts...)
	grpc_zap.ReplaceGrpcLogger(logger)

	// Register APIs to grpc server
	registerGRPCAPIs(grpcServer, objectStore)

	// Initialize Prometheus GRPC metrics.
	grpcMetrics.InitializeMetrics(grpcServer)

	// Create a HTTP server for prometheus.
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(promReg, promhttp.HandlerOpts{}),
		Addr:    viper.GetString(flagListenAddressHTTP),
	}

	// TODO Implement TLS
	lis, err := net.Listen("tcp", viper.GetString(flagListenAddressGRPC))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	closer := utils.ChannelCloser{}
	wg := sync.WaitGroup{}

	// TODO Create a "mesh" proxy.
	// Would be cool if each daemon can tell other servers where the masters are or UnaryServerInterceptor(),even proxy to them.

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to listen and serve http server", zap.Error(err))
			closer.Close(stopCh)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			logger.Error("grpcServer.Serve() returned non-nil error on GracefulStop", zap.Error(err))
			closer.Close(stopCh)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := magicRun(stopCh, objectStore); err != nil {
			logger.Error("magic run returned error", zap.Error(err))
			closer.Close(stopCh)
		}
	}()

	select {
	case <-sigCh:
		logger.Info("signal received, shutting down ...")
		closer.Close(stopCh)
	case <-stopCh:
		logger.Info("stop channel closed, shutting down ...")
	}

	// Shutdown grpc server
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

func registerGRPCAPIs(srv *grpc.Server, objectStore *object.Store) {
	// Beacons
	beaconsServer := beacons_v1alpha.NewBeaconsService(objectStore.Prefixed(beacons_v1alpha.APIPath))
	beacons_v1alpha.RegisterBeaconsServer(srv, beaconsServer)
	// ClusterVariables
	clusterVariablesServer := clustervariables_v1alpha.NewClusterVariablesService(objectStore.Prefixed(clustervariables_v1alpha.APIPath))
	clustervariables_v1alpha.RegisterClusterVariablesServer(srv, clusterVariablesServer)
	// Configs (AgentConfigs)
	configsServer := configs_v1alpha.NewAgentConfigsService(objectStore.Prefixed(configs_v1alpha.APIPath))
	configs_v1alpha.RegisterAgentConfigsServer(srv, configsServer)
	// CronJobs
	cronJobsServer := cronjobs_v1alpha.NewCronJobsService(objectStore.Prefixed(cronjobs_v1alpha.APIPath))
	cronjobs_v1alpha.RegisterCronJobsServer(srv, cronJobsServer)
	// Events
	eventsServer := events_v1alpha.NewEventsService(objectStore.Prefixed(events_v1alpha.APIPath))
	events_v1alpha.RegisterEventsServer(srv, eventsServer)
	// Nodes
	nodesServer := nodes_v1alpha.NewNodesService(objectStore.Prefixed(nodes_v1alpha.APIPath))
	nodes_v1alpha.RegisterNodesServer(srv, nodesServer)
	// Secrets
	secretsServer := secrets_v1alpha.NewSecretsService(objectStore.Prefixed(secrets_v1alpha.APIPath))
	secrets_v1alpha.RegisterSecretsServer(srv, secretsServer)
	// TaskBooks
	taskBooksServer := taskbooks_v1alpha.NewTaskBooksService(objectStore.Prefixed(taskbooks_v1alpha.APIPath))
	taskbooks_v1alpha.RegisterTaskBooksServer(srv, taskBooksServer)
	// Triggers
	triggersServer := triggers_v1alpha.NewTriggersService(objectStore.Prefixed(triggers_v1alpha.APIPath))
	triggers_v1alpha.RegisterTriggersServer(srv, triggersServer)
	// Variables
	variablesServer := variables_v1alpha.NewVariablesService(objectStore.Prefixed(variables_v1alpha.APIPath))
	variables_v1alpha.RegisterVariablesServer(srv, variablesServer)
}

// Shutdown command
func Shutdown(cmd *cobra.Command, args []string) error {
	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sigCounter := 0
	go func() {
		for {
			<-sigCh
			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
			logger.Info("signal received, ignoring ...")
			if sigCounter > 0 {
				close(stopCh)
			}
			sigCounter++
		}
	}()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	cc, err := grpc.Dial("127.0.0.1:1337", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer cc.Close()
	nodesClient := nodes_v1alpha.NewNodesClient(cc)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	getReq := &nodes_v1alpha.GetRequest{
		Options: &core_v1.GetOptions{
			Name: hostname,
		},
	}
	getResponse, err := nodesClient.Get(ctx, getReq)
	fmt.Printf("Nodes.Get(): %+v - %+v\n", getResponse, err)
	// TODO mark node as (shut)down(ed) with reason (e.g., shutdown of server, service restart)
	//getResponse.Node.Spec

	return nil
}
