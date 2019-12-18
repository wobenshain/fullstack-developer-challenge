// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/wobenshain/fullstack-developer-challenge/go/restapi/models"
)

// GetItemOKCode is the HTTP code returned for type GetItemOK
const GetItemOKCode int = 200

/*GetItemOK list all items

swagger:response getItemOK
*/
type GetItemOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Item `json:"body,omitempty"`
}

// NewGetItemOK creates GetItemOK with default headers values
func NewGetItemOK() *GetItemOK {

	return &GetItemOK{}
}

// WithPayload adds the payload to the get item o k response
func (o *GetItemOK) WithPayload(payload []*models.Item) *GetItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get item o k response
func (o *GetItemOK) SetPayload(payload []*models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Item, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetItemDefault generic error response

swagger:response getItemDefault
*/
type GetItemDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetItemDefault creates GetItemDefault with default headers values
func NewGetItemDefault(code int) *GetItemDefault {
	if code <= 0 {
		code = 500
	}

	return &GetItemDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get item default response
func (o *GetItemDefault) WithStatusCode(code int) *GetItemDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get item default response
func (o *GetItemDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get item default response
func (o *GetItemDefault) WithPayload(payload *models.Error) *GetItemDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get item default response
func (o *GetItemDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
