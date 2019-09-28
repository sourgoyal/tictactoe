// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostAPIV1GamesCreatedCode is the HTTP code returned for type PostAPIV1GamesCreated
const PostAPIV1GamesCreatedCode int = 201

/*PostAPIV1GamesCreated Game successfully started

swagger:response postApiV1GamesCreated
*/
type PostAPIV1GamesCreated struct {
	/*URL of the started game

	 */
	Location string `json:"Location"`

	/*
	  In: Body
	*/
	Payload *PostAPIV1GamesCreatedBody `json:"body,omitempty"`
}

// NewPostAPIV1GamesCreated creates PostAPIV1GamesCreated with default headers values
func NewPostAPIV1GamesCreated() *PostAPIV1GamesCreated {

	return &PostAPIV1GamesCreated{}
}

// WithLocation adds the location to the post Api v1 games created response
func (o *PostAPIV1GamesCreated) WithLocation(location string) *PostAPIV1GamesCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the post Api v1 games created response
func (o *PostAPIV1GamesCreated) SetLocation(location string) {
	o.Location = location
}

// WithPayload adds the payload to the post Api v1 games created response
func (o *PostAPIV1GamesCreated) WithPayload(payload *PostAPIV1GamesCreatedBody) *PostAPIV1GamesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Api v1 games created response
func (o *PostAPIV1GamesCreated) SetPayload(payload *PostAPIV1GamesCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAPIV1GamesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAPIV1GamesBadRequestCode is the HTTP code returned for type PostAPIV1GamesBadRequest
const PostAPIV1GamesBadRequestCode int = 400

/*PostAPIV1GamesBadRequest Bad request

swagger:response postApiV1GamesBadRequest
*/
type PostAPIV1GamesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostAPIV1GamesBadRequestBody `json:"body,omitempty"`
}

// NewPostAPIV1GamesBadRequest creates PostAPIV1GamesBadRequest with default headers values
func NewPostAPIV1GamesBadRequest() *PostAPIV1GamesBadRequest {

	return &PostAPIV1GamesBadRequest{}
}

// WithPayload adds the payload to the post Api v1 games bad request response
func (o *PostAPIV1GamesBadRequest) WithPayload(payload *PostAPIV1GamesBadRequestBody) *PostAPIV1GamesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Api v1 games bad request response
func (o *PostAPIV1GamesBadRequest) SetPayload(payload *PostAPIV1GamesBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAPIV1GamesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAPIV1GamesNotFoundCode is the HTTP code returned for type PostAPIV1GamesNotFound
const PostAPIV1GamesNotFoundCode int = 404

/*PostAPIV1GamesNotFound Resource not found

swagger:response postApiV1GamesNotFound
*/
type PostAPIV1GamesNotFound struct {
}

// NewPostAPIV1GamesNotFound creates PostAPIV1GamesNotFound with default headers values
func NewPostAPIV1GamesNotFound() *PostAPIV1GamesNotFound {

	return &PostAPIV1GamesNotFound{}
}

// WriteResponse to the client
func (o *PostAPIV1GamesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PostAPIV1GamesInternalServerErrorCode is the HTTP code returned for type PostAPIV1GamesInternalServerError
const PostAPIV1GamesInternalServerErrorCode int = 500

/*PostAPIV1GamesInternalServerError Internal server error

swagger:response postApiV1GamesInternalServerError
*/
type PostAPIV1GamesInternalServerError struct {
}

// NewPostAPIV1GamesInternalServerError creates PostAPIV1GamesInternalServerError with default headers values
func NewPostAPIV1GamesInternalServerError() *PostAPIV1GamesInternalServerError {

	return &PostAPIV1GamesInternalServerError{}
}

// WriteResponse to the client
func (o *PostAPIV1GamesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
