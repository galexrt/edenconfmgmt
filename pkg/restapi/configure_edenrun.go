// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/galexrt/edenrun/pkg/restapi/operations"
	"github.com/galexrt/edenrun/pkg/restapi/operations/core"
	"github.com/galexrt/edenrun/pkg/restapi/operations/nodes"
)

//go:generate swagger generate server --target ../../pkg --name Edenrun --spec ../../swagger.yaml

func configureFlags(api *operations.EdenrunAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EdenrunAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.NodesDeleteApisCoreEdenRunV1alphaNodesHandler = nodes.DeleteApisCoreEdenRunV1alphaNodesHandlerFunc(func(params nodes.DeleteApisCoreEdenRunV1alphaNodesParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.DeleteApisCoreEdenRunV1alphaNodes has not yet been implemented")
	})
	api.NodesGetApisCoreEdenRunV1alphaNodesHandler = nodes.GetApisCoreEdenRunV1alphaNodesHandlerFunc(func(params nodes.GetApisCoreEdenRunV1alphaNodesParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.GetApisCoreEdenRunV1alphaNodes has not yet been implemented")
	})
	api.NodesGetApisCoreEdenRunV1alphaWatchNodesHandler = nodes.GetApisCoreEdenRunV1alphaWatchNodesHandlerFunc(func(params nodes.GetApisCoreEdenRunV1alphaWatchNodesParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.GetApisCoreEdenRunV1alphaWatchNodes has not yet been implemented")
	})
	api.CoreGetVersionHandler = core.GetVersionHandlerFunc(func(params core.GetVersionParams) middleware.Responder {
		return middleware.NotImplemented("operation core.GetVersion has not yet been implemented")
	})
	api.NodesPatchApisCoreEdenRunV1alphaNodesHandler = nodes.PatchApisCoreEdenRunV1alphaNodesHandlerFunc(func(params nodes.PatchApisCoreEdenRunV1alphaNodesParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.PatchApisCoreEdenRunV1alphaNodes has not yet been implemented")
	})
	api.NodesPostApisCoreEdenRunV1alphaNodesHandler = nodes.PostApisCoreEdenRunV1alphaNodesHandlerFunc(func(params nodes.PostApisCoreEdenRunV1alphaNodesParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.PostApisCoreEdenRunV1alphaNodes has not yet been implemented")
	})
	api.NodesPutApisCoreEdenRunV1alphaNodesHandler = nodes.PutApisCoreEdenRunV1alphaNodesHandlerFunc(func(params nodes.PutApisCoreEdenRunV1alphaNodesParams) middleware.Responder {
		return middleware.NotImplemented("operation nodes.PutApisCoreEdenRunV1alphaNodes has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
