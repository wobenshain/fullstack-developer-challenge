// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/wobenshain/fullstack-developer-challenge/go/restapi/models"
)

// GetItemIDOKCode is the HTTP code returned for type GetItemIDOK
const GetItemIDOKCode int = 200

/*GetItemIDOK get an item

swagger:response getItemIdOK
*/
type GetItemIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewGetItemIDOK creates GetItemIDOK with default headers values
func NewGetItemIDOK() *GetItemIDOK {

	return &GetItemIDOK{}
}

// WithPayload adds the payload to the get item Id o k response
func (o *GetItemIDOK) WithPayload(payload *models.Item) *GetItemIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get item Id o k response
func (o *GetItemIDOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetItemIDDefault generic error response

swagger:response getItemIdDefault
*/
type GetItemIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetItemIDDefault creates GetItemIDDefault with default headers values
func NewGetItemIDDefault(code int) *GetItemIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetItemIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get item ID default response
func (o *GetItemIDDefault) WithStatusCode(code int) *GetItemIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get item ID default response
func (o *GetItemIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get item ID default response
func (o *GetItemIDDefault) WithPayload(payload *models.Error) *GetItemIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get item ID default response
func (o *GetItemIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
