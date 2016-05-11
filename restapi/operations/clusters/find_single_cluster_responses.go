package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/redhatanalytics/oshinko-rest/models"
)

/*FindSingleClusterOK Cluster detail response

swagger:response findSingleClusterOK
*/
type FindSingleClusterOK struct {

	// In: body
	Payload *models.SingleCluster `json:"body,omitempty"`
}

// NewFindSingleClusterOK creates FindSingleClusterOK with default headers values
func NewFindSingleClusterOK() *FindSingleClusterOK {
	return &FindSingleClusterOK{}
}

// WithPayload adds the payload to the find single cluster o k response
func (o *FindSingleClusterOK) WithPayload(payload *models.SingleCluster) *FindSingleClusterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find single cluster o k response
func (o *FindSingleClusterOK) SetPayload(payload *models.SingleCluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindSingleClusterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindSingleClusterDefault Unexpected error

swagger:response findSingleClusterDefault
*/
type FindSingleClusterDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFindSingleClusterDefault creates FindSingleClusterDefault with default headers values
func NewFindSingleClusterDefault(code int) *FindSingleClusterDefault {
	if code <= 0 {
		code = 500
	}

	return &FindSingleClusterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find single cluster default response
func (o *FindSingleClusterDefault) WithStatusCode(code int) *FindSingleClusterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find single cluster default response
func (o *FindSingleClusterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find single cluster default response
func (o *FindSingleClusterDefault) WithPayload(payload *models.ErrorResponse) *FindSingleClusterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find single cluster default response
func (o *FindSingleClusterDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindSingleClusterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
