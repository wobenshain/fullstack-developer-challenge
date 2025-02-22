// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostItemHandlerFunc turns a function with the right signature into a post item handler
type PostItemHandlerFunc func(PostItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostItemHandlerFunc) Handle(params PostItemParams) middleware.Responder {
	return fn(params)
}

// PostItemHandler interface for that can handle valid post item params
type PostItemHandler interface {
	Handle(PostItemParams) middleware.Responder
}

// NewPostItem creates a new http.Handler for the post item operation
func NewPostItem(ctx *middleware.Context, handler PostItemHandler) *PostItem {
	return &PostItem{Context: ctx, Handler: handler}
}

/*PostItem swagger:route POST /item postItem

PostItem post item API

*/
type PostItem struct {
	Context *middleware.Context
	Handler PostItemHandler
}

func (o *PostItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
