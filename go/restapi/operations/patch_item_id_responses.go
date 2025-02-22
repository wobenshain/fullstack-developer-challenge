// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/wobenshain/fullstack-developer-challenge/go/restapi/models"
)

// PatchItemIDOKCode is the HTTP code returned for type PatchItemIDOK
const PatchItemIDOKCode int = 200

/*PatchItemIDOK update an item

swagger:response patchItemIdOK
*/
type PatchItemIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewPatchItemIDOK creates PatchItemIDOK with default headers values
func NewPatchItemIDOK() *PatchItemIDOK {

	return &PatchItemIDOK{}
}

// WithPayload adds the payload to the patch item Id o k response
func (o *PatchItemIDOK) WithPayload(payload *models.Item) *PatchItemIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch item Id o k response
func (o *PatchItemIDOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchItemIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchItemIDDefault generic error response

swagger:response patchItemIdDefault
*/
type PatchItemIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchItemIDDefault creates PatchItemIDDefault with default headers values
func NewPatchItemIDDefault(code int) *PatchItemIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchItemIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch item ID default response
func (o *PatchItemIDDefault) WithStatusCode(code int) *PatchItemIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch item ID default response
func (o *PatchItemIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch item ID default response
func (o *PatchItemIDDefault) WithPayload(payload *models.Error) *PatchItemIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch item ID default response
func (o *PatchItemIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchItemIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
