// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostApisCoreEdenRunV1alphaNodesHandlerFunc turns a function with the right signature into a post apis core eden run v1alpha nodes handler
type PostApisCoreEdenRunV1alphaNodesHandlerFunc func(PostApisCoreEdenRunV1alphaNodesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostApisCoreEdenRunV1alphaNodesHandlerFunc) Handle(params PostApisCoreEdenRunV1alphaNodesParams) middleware.Responder {
	return fn(params)
}

// PostApisCoreEdenRunV1alphaNodesHandler interface for that can handle valid post apis core eden run v1alpha nodes params
type PostApisCoreEdenRunV1alphaNodesHandler interface {
	Handle(PostApisCoreEdenRunV1alphaNodesParams) middleware.Responder
}

// NewPostApisCoreEdenRunV1alphaNodes creates a new http.Handler for the post apis core eden run v1alpha nodes operation
func NewPostApisCoreEdenRunV1alphaNodes(ctx *middleware.Context, handler PostApisCoreEdenRunV1alphaNodesHandler) *PostApisCoreEdenRunV1alphaNodes {
	return &PostApisCoreEdenRunV1alphaNodes{Context: ctx, Handler: handler}
}

/*PostApisCoreEdenRunV1alphaNodes swagger:route POST /apis/core.eden.run/v1alpha/nodes core.eden.run/v1alpha/nodes postApisCoreEdenRunV1alphaNodes

Update a Node object.

*/
type PostApisCoreEdenRunV1alphaNodes struct {
	Context *middleware.Context
	Handler PostApisCoreEdenRunV1alphaNodesHandler
}

func (o *PostApisCoreEdenRunV1alphaNodes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostApisCoreEdenRunV1alphaNodesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
