// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteAPIV1GamesGameIDHandlerFunc turns a function with the right signature into a delete API v1 games game ID handler
type DeleteAPIV1GamesGameIDHandlerFunc func(DeleteAPIV1GamesGameIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAPIV1GamesGameIDHandlerFunc) Handle(params DeleteAPIV1GamesGameIDParams) middleware.Responder {
	return fn(params)
}

// DeleteAPIV1GamesGameIDHandler interface for that can handle valid delete API v1 games game ID params
type DeleteAPIV1GamesGameIDHandler interface {
	Handle(DeleteAPIV1GamesGameIDParams) middleware.Responder
}

// NewDeleteAPIV1GamesGameID creates a new http.Handler for the delete API v1 games game ID operation
func NewDeleteAPIV1GamesGameID(ctx *middleware.Context, handler DeleteAPIV1GamesGameIDHandler) *DeleteAPIV1GamesGameID {
	return &DeleteAPIV1GamesGameID{Context: ctx, Handler: handler}
}

/*DeleteAPIV1GamesGameID swagger:route DELETE /api/v1/games/{game_id} deleteApiV1GamesGameId

Delete a game.

*/
type DeleteAPIV1GamesGameID struct {
	Context *middleware.Context
	Handler DeleteAPIV1GamesGameIDHandler
}

func (o *DeleteAPIV1GamesGameID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAPIV1GamesGameIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}