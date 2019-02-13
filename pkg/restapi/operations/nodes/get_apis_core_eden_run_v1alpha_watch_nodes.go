// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetApisCoreEdenRunV1alphaWatchNodesHandlerFunc turns a function with the right signature into a get apis core eden run v1alpha watch nodes handler
type GetApisCoreEdenRunV1alphaWatchNodesHandlerFunc func(GetApisCoreEdenRunV1alphaWatchNodesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApisCoreEdenRunV1alphaWatchNodesHandlerFunc) Handle(params GetApisCoreEdenRunV1alphaWatchNodesParams) middleware.Responder {
	return fn(params)
}

// GetApisCoreEdenRunV1alphaWatchNodesHandler interface for that can handle valid get apis core eden run v1alpha watch nodes params
type GetApisCoreEdenRunV1alphaWatchNodesHandler interface {
	Handle(GetApisCoreEdenRunV1alphaWatchNodesParams) middleware.Responder
}

// NewGetApisCoreEdenRunV1alphaWatchNodes creates a new http.Handler for the get apis core eden run v1alpha watch nodes operation
func NewGetApisCoreEdenRunV1alphaWatchNodes(ctx *middleware.Context, handler GetApisCoreEdenRunV1alphaWatchNodesHandler) *GetApisCoreEdenRunV1alphaWatchNodes {
	return &GetApisCoreEdenRunV1alphaWatchNodes{Context: ctx, Handler: handler}
}

/*GetApisCoreEdenRunV1alphaWatchNodes swagger:route GET /apis/core.eden.run/v1alpha/watch/nodes core.eden.run/v1alpha/nodes getApisCoreEdenRunV1alphaWatchNodes

Watch Node objects.

*/
type GetApisCoreEdenRunV1alphaWatchNodes struct {
	Context *middleware.Context
	Handler GetApisCoreEdenRunV1alphaWatchNodesHandler
}

func (o *GetApisCoreEdenRunV1alphaWatchNodes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetApisCoreEdenRunV1alphaWatchNodesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
