// Code generated by go-swagger; DO NOT EDIT.

package jobservice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xuesea/goharbor-client/v5/apiv2/model"
)

// GetWorkerPoolsReader is a Reader for the GetWorkerPools structure.
type GetWorkerPoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkerPoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkerPoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkerPoolsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkerPoolsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkerPoolsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkerPoolsOK creates a GetWorkerPoolsOK with default headers values
func NewGetWorkerPoolsOK() *GetWorkerPoolsOK {
	return &GetWorkerPoolsOK{}
}

/*GetWorkerPoolsOK handles this case with default header values.

Get worker pools successfully.
*/
type GetWorkerPoolsOK struct {
	Payload []*model.WorkerPool
}

func (o *GetWorkerPoolsOK) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools][%d] getWorkerPoolsOK  %+v", 200, o.Payload)
}

func (o *GetWorkerPoolsOK) GetPayload() []*model.WorkerPool {
	return o.Payload
}

func (o *GetWorkerPoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkerPoolsUnauthorized creates a GetWorkerPoolsUnauthorized with default headers values
func NewGetWorkerPoolsUnauthorized() *GetWorkerPoolsUnauthorized {
	return &GetWorkerPoolsUnauthorized{}
}

/*GetWorkerPoolsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetWorkerPoolsUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWorkerPoolsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools][%d] getWorkerPoolsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkerPoolsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWorkerPoolsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkerPoolsForbidden creates a GetWorkerPoolsForbidden with default headers values
func NewGetWorkerPoolsForbidden() *GetWorkerPoolsForbidden {
	return &GetWorkerPoolsForbidden{}
}

/*GetWorkerPoolsForbidden handles this case with default header values.

Forbidden
*/
type GetWorkerPoolsForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWorkerPoolsForbidden) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools][%d] getWorkerPoolsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkerPoolsForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWorkerPoolsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkerPoolsInternalServerError creates a GetWorkerPoolsInternalServerError with default headers values
func NewGetWorkerPoolsInternalServerError() *GetWorkerPoolsInternalServerError {
	return &GetWorkerPoolsInternalServerError{}
}

/*GetWorkerPoolsInternalServerError handles this case with default header values.

Internal server error
*/
type GetWorkerPoolsInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWorkerPoolsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools][%d] getWorkerPoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkerPoolsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWorkerPoolsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
