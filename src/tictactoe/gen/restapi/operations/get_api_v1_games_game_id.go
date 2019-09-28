// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAPIV1GamesGameIDHandlerFunc turns a function with the right signature into a get API v1 games game ID handler
type GetAPIV1GamesGameIDHandlerFunc func(GetAPIV1GamesGameIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAPIV1GamesGameIDHandlerFunc) Handle(params GetAPIV1GamesGameIDParams) middleware.Responder {
	return fn(params)
}

// GetAPIV1GamesGameIDHandler interface for that can handle valid get API v1 games game ID params
type GetAPIV1GamesGameIDHandler interface {
	Handle(GetAPIV1GamesGameIDParams) middleware.Responder
}

// NewGetAPIV1GamesGameID creates a new http.Handler for the get API v1 games game ID operation
func NewGetAPIV1GamesGameID(ctx *middleware.Context, handler GetAPIV1GamesGameIDHandler) *GetAPIV1GamesGameID {
	return &GetAPIV1GamesGameID{Context: ctx, Handler: handler}
}

/*GetAPIV1GamesGameID swagger:route GET /api/v1/games/{game_id} getApiV1GamesGameId

Get a game.

*/
type GetAPIV1GamesGameID struct {
	Context *middleware.Context
	Handler GetAPIV1GamesGameIDHandler
}

func (o *GetAPIV1GamesGameID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAPIV1GamesGameIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}