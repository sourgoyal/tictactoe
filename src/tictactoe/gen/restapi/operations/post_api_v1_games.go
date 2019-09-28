// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostAPIV1GamesHandlerFunc turns a function with the right signature into a post API v1 games handler
type PostAPIV1GamesHandlerFunc func(PostAPIV1GamesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAPIV1GamesHandlerFunc) Handle(params PostAPIV1GamesParams) middleware.Responder {
	return fn(params)
}

// PostAPIV1GamesHandler interface for that can handle valid post API v1 games params
type PostAPIV1GamesHandler interface {
	Handle(PostAPIV1GamesParams) middleware.Responder
}

// NewPostAPIV1Games creates a new http.Handler for the post API v1 games operation
func NewPostAPIV1Games(ctx *middleware.Context, handler PostAPIV1GamesHandler) *PostAPIV1Games {
	return &PostAPIV1Games{Context: ctx, Handler: handler}
}

/*PostAPIV1Games swagger:route POST /api/v1/games postApiV1Games

Start a new game.

*/
type PostAPIV1Games struct {
	Context *middleware.Context
	Handler PostAPIV1GamesHandler
}

func (o *PostAPIV1Games) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAPIV1GamesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAPIV1GamesBadRequestBody post API v1 games bad request body
// swagger:model PostAPIV1GamesBadRequestBody
type PostAPIV1GamesBadRequestBody struct {

	// Why the game failed to start
	Reason string `json:"reason,omitempty"`
}

// Validate validates this post API v1 games bad request body
func (o *PostAPIV1GamesBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAPIV1GamesBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAPIV1GamesBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostAPIV1GamesBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAPIV1GamesCreatedBody post API v1 games created body
// swagger:model PostAPIV1GamesCreatedBody
type PostAPIV1GamesCreatedBody struct {

	// URL of the started game
	Location string `json:"location,omitempty"`
}

// Validate validates this post API v1 games created body
func (o *PostAPIV1GamesCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAPIV1GamesCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAPIV1GamesCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostAPIV1GamesCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}