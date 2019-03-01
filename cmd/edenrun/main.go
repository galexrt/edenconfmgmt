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
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/coreos/go-systemd/daemon"
	"github.com/galexrt/edenrun/pkg/service"
	"github.com/galexrt/edenrun/pkg/store/cache"
	"github.com/galexrt/edenrun/pkg/store/object"
	"github.com/galexrt/edenrun/pkg/tracer"
	"github.com/galexrt/edenrun/pkg/utils"

	_ "net/http/pprof"

	events_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/core/events/v1alpha"
	nodes_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/core/nodes/v1alpha"
	configs_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/daemons/configs/v1alpha"
	clustervariables_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/data/clustervariables/v1alpha"
	secrets_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/data/secrets/v1alpha"
	variables_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/data/variables/v1alpha"
	cronjobs_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/jobs/cronjobs/v1alpha"
	playbooks_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/runny/playbooks/v1alpha"
	beacons_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/beacons/v1alpha"
	triggers_v1alpha "github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/triggers/v1alpha"

	"github.com/galexrt/edenrun/pkg/auth"
	store_data_adapters "github.com/galexrt/edenrun/pkg/store/data/adapters"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

const (
	flagConfigFile = "config"
	flagLogLevel   = "log-level"
)

var (
	logLevel = zap.NewAtomicLevel()

	// Logger for logging.
	logger *zap.Logger
	// Prometheus metrics registry.
	promReg = prometheus.NewRegistry()
	// grpc prometheus server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "edenrun",
		Short: "Configuration management with automatic clustering, events and stuff.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var level zapcore.Level
			if err := (&level).UnmarshalText([]byte(viper.GetString(flagLogLevel))); err != nil {
				panic(err)
			}
			logLevel.SetLevel(level)

			encoderCfg := zap.NewDevelopmentEncoderConfig()

			logger = zap.New(zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderCfg),
				zapcore.Lock(os.Stdout),
				logLevel,
			))
		},
		RunE: Run,
	}
)

func init() {
	promReg.MustRegister(grpcMetrics)

	rootCmd.PersistentFlags().String(flagConfigFile, "config.yaml", "config file path")
	rootCmd.MarkPersistentFlagRequired(flagConfigFile)
}

func main() {
	Execute()
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

// Run runs all routines to start the work of edenrun application.
func Run(cmd *cobra.Command, args []string) error {
	// Run Plan:
	// 1. Create logger.
	// 2. Read config file.
	// 3. Start up master or node parts, depending on config.

	defer logger.Sync()
	logger.Info("starting", zap.String("command", os.Args[0]), zap.String("versionInfo", version.Info()))

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	logger.Debug("signal handler installed")
	stopCh := make(chan struct{})

	svcCtrl := service.New(logger.Named("service"))
	go svcCtrl.Start(stopCh)

	wg := sync.WaitGroup{}

	if err := tracer.Init(); err != nil {
		logger.Fatal("failed to init tracer", zap.Error(err))
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := tracer.Start(svcCtrl.GetStopCh("tracer")); err != nil {
			logger.Error("error during tracer stop", zap.Error(err))
		}
	}()

	informer, objectStore := getStore()

	// Auth
	authProvider, err := auth.New()
	if err != nil {
		logger.Fatal("failed to init auth", zap.Error(err))
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{
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
	}...)
	grpc_zap.ReplaceGrpcLogger(logger)

	registerAPIs(grpcServer, informer, objectStore)

	// Initialize Prometheus GRPC metrics.
	grpcMetrics.InitializeMetrics(grpcServer)

	// Create a HTTP server for prometheus.
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(promReg, promhttp.HandlerOpts{}),
		Addr:    "0.0.0.0:1338",
	}

	// TODO Implement TLS
	lis, err := net.Listen("tcp", "0.0.0.0:1337")
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	closer := utils.ChannelCloser{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := informer.Start(stopCh); err != nil {
			logger.Error("failed to start/run informer", zap.Error(err))
			closer.Close(stopCh)
		}
	}()

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

	daemon.SdNotify(true, daemon.SdNotifyReady)

	select {
	case <-sigCh:
		logger.Info("signal received, shutting down ...")
	case <-stopCh:
		logger.Info("stop channel closed, shutting down ...")
	}

	closer.Close(stopCh)

	daemon.SdNotify(true, daemon.SdNotifyStopping)

	// Shutdown grpc server
	grpcServer.GracefulStop()

	// Shutdown http server
	ctx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("http server graceful shutdown failed", zap.Error(err))
	}

	// Wait for all routines to have exited
	wg.Wait()
	return nil
}

func getStore() (*object.Informer, *object.Store) {
	logger.Info("getting data store")
	dataStoreCtx, dataStoreCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer dataStoreCancel()
	dataStoreHandlerName := "etcd"
	dataStoreAdapter, err := store_data_adapters.Get(dataStoreCtx, dataStoreHandlerName)
	if err != nil {
		logger.Fatal("failed to create store handler", zap.String("storehandler", dataStoreHandlerName), zap.Error(err))
	}

	logger.Info("getting cache store")
	cacheStoreCtx, cacheStoreCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cacheStoreCancel()
	cacheStoreHandlerName := "memory"
	cacheStoreAdapter, err := store_data_adapters.Get(cacheStoreCtx, cacheStoreHandlerName)
	if err != nil {
		logger.Fatal("failed to create store handler", zap.String("storehandler", cacheStoreHandlerName), zap.Error(err))
	}
	cacheStore := cache.New(dataStoreAdapter, cacheStoreAdapter, logger)

	informer := object.NewInformer(cacheStore, logger)
	objectStore := object.NewStore(cacheStore, informer, logger)

	return informer, objectStore
}

func registerAPIs(srv *grpc.Server, informer *object.Informer, objectStore *object.Store) {
	// Beacons
	informer.Register(beacons_v1alpha.APIPath)
	beaconsServer := beacons_v1alpha.NewBeaconsService(objectStore.Prefixed(beacons_v1alpha.APIPath))
	beacons_v1alpha.RegisterBeaconsServer(srv, beaconsServer)
	// ClusterVariables
	informer.Register(clustervariables_v1alpha.APIPath)
	clusterVariablesServer := clustervariables_v1alpha.NewClusterVariablesService(objectStore.Prefixed(clustervariables_v1alpha.APIPath))
	clustervariables_v1alpha.RegisterClusterVariablesServer(srv, clusterVariablesServer)
	// Configs (Configs)
	informer.Register(configs_v1alpha.APIPath)
	configsServer := configs_v1alpha.NewConfigsService(objectStore.Prefixed(configs_v1alpha.APIPath))
	configs_v1alpha.RegisterConfigsServer(srv, configsServer)
	// CronJobs
	informer.Register(cronjobs_v1alpha.APIPath)
	cronJobsServer := cronjobs_v1alpha.NewCronJobsService(objectStore.Prefixed(cronjobs_v1alpha.APIPath))
	cronjobs_v1alpha.RegisterCronJobsServer(srv, cronJobsServer)
	// Events
	informer.Register(events_v1alpha.APIPath)
	eventsServer := events_v1alpha.NewEventsService(objectStore.Prefixed(events_v1alpha.APIPath))
	events_v1alpha.RegisterEventsServer(srv, eventsServer)
	// Nodes
	informer.Register(nodes_v1alpha.APIPath)
	nodesServer := nodes_v1alpha.NewNodesService(objectStore.Prefixed(nodes_v1alpha.APIPath))
	nodes_v1alpha.RegisterNodesServer(srv, nodesServer)
	// Secrets
	informer.Register(secrets_v1alpha.APIPath)
	secretsServer := secrets_v1alpha.NewSecretsService(objectStore.Prefixed(secrets_v1alpha.APIPath))
	secrets_v1alpha.RegisterSecretsServer(srv, secretsServer)
	// PlayBooks
	informer.Register(playbooks_v1alpha.APIPath)
	playBooksServer := playbooks_v1alpha.NewPlayBooksService(objectStore.Prefixed(playbooks_v1alpha.APIPath))
	playbooks_v1alpha.RegisterPlayBooksServer(srv, playBooksServer)
	// Triggers
	informer.Register(triggers_v1alpha.APIPath)
	triggersServer := triggers_v1alpha.NewTriggersService(objectStore.Prefixed(triggers_v1alpha.APIPath))
	triggers_v1alpha.RegisterTriggersServer(srv, triggersServer)
	// Variables
	informer.Register(variables_v1alpha.APIPath)
	variablesServer := variables_v1alpha.NewVariablesService(objectStore.Prefixed(variables_v1alpha.APIPath))
	variables_v1alpha.RegisterVariablesServer(srv, variablesServer)
}
