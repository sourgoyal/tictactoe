// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "tictactoe/gen/models"
)

// GetAPIV1GamesOKCode is the HTTP code returned for type GetAPIV1GamesOK
const GetAPIV1GamesOKCode int = 200

/*GetAPIV1GamesOK Successful response, returns an array of games, returns an empty array if no users found

swagger:response getApiV1GamesOK
*/
type GetAPIV1GamesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Game `json:"body,omitempty"`
}

// NewGetAPIV1GamesOK creates GetAPIV1GamesOK with default headers values
func NewGetAPIV1GamesOK() *GetAPIV1GamesOK {

	return &GetAPIV1GamesOK{}
}

// WithPayload adds the payload to the get Api v1 games o k response
func (o *GetAPIV1GamesOK) WithPayload(payload []*models.Game) *GetAPIV1GamesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api v1 games o k response
func (o *GetAPIV1GamesOK) SetPayload(payload []*models.Game) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIV1GamesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Game, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAPIV1GamesBadRequestCode is the HTTP code returned for type GetAPIV1GamesBadRequest
const GetAPIV1GamesBadRequestCode int = 400

/*GetAPIV1GamesBadRequest Bad request

swagger:response getApiV1GamesBadRequest
*/
type GetAPIV1GamesBadRequest struct {
}

// NewGetAPIV1GamesBadRequest creates GetAPIV1GamesBadRequest with default headers values
func NewGetAPIV1GamesBadRequest() *GetAPIV1GamesBadRequest {

	return &GetAPIV1GamesBadRequest{}
}

// WriteResponse to the client
func (o *GetAPIV1GamesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAPIV1GamesNotFoundCode is the HTTP code returned for type GetAPIV1GamesNotFound
const GetAPIV1GamesNotFoundCode int = 404

/*GetAPIV1GamesNotFound Resource not found

swagger:response getApiV1GamesNotFound
*/
type GetAPIV1GamesNotFound struct {
}

// NewGetAPIV1GamesNotFound creates GetAPIV1GamesNotFound with default headers values
func NewGetAPIV1GamesNotFound() *GetAPIV1GamesNotFound {

	return &GetAPIV1GamesNotFound{}
}

// WriteResponse to the client
func (o *GetAPIV1GamesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetAPIV1GamesInternalServerErrorCode is the HTTP code returned for type GetAPIV1GamesInternalServerError
const GetAPIV1GamesInternalServerErrorCode int = 500

/*GetAPIV1GamesInternalServerError Internal server error

swagger:response getApiV1GamesInternalServerError
*/
type GetAPIV1GamesInternalServerError struct {
}

// NewGetAPIV1GamesInternalServerError creates GetAPIV1GamesInternalServerError with default headers values
func NewGetAPIV1GamesInternalServerError() *GetAPIV1GamesInternalServerError {

	return &GetAPIV1GamesInternalServerError{}
}

// WriteResponse to the client
func (o *GetAPIV1GamesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
