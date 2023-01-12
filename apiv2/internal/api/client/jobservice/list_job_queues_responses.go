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

// ListJobQueuesReader is a Reader for the ListJobQueues structure.
type ListJobQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListJobQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListJobQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListJobQueuesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListJobQueuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListJobQueuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListJobQueuesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListJobQueuesOK creates a ListJobQueuesOK with default headers values
func NewListJobQueuesOK() *ListJobQueuesOK {
	return &ListJobQueuesOK{}
}

/*ListJobQueuesOK handles this case with default header values.

List job queue successfully.
*/
type ListJobQueuesOK struct {
	Payload []*model.JobQueue
}

func (o *ListJobQueuesOK) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesOK  %+v", 200, o.Payload)
}

func (o *ListJobQueuesOK) GetPayload() []*model.JobQueue {
	return o.Payload
}

func (o *ListJobQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobQueuesUnauthorized creates a ListJobQueuesUnauthorized with default headers values
func NewListJobQueuesUnauthorized() *ListJobQueuesUnauthorized {
	return &ListJobQueuesUnauthorized{}
}

/*ListJobQueuesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListJobQueuesUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListJobQueuesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListJobQueuesUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListJobQueuesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobQueuesForbidden creates a ListJobQueuesForbidden with default headers values
func NewListJobQueuesForbidden() *ListJobQueuesForbidden {
	return &ListJobQueuesForbidden{}
}

/*ListJobQueuesForbidden handles this case with default header values.

Forbidden
*/
type ListJobQueuesForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListJobQueuesForbidden) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesForbidden  %+v", 403, o.Payload)
}

func (o *ListJobQueuesForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListJobQueuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobQueuesNotFound creates a ListJobQueuesNotFound with default headers values
func NewListJobQueuesNotFound() *ListJobQueuesNotFound {
	return &ListJobQueuesNotFound{}
}

/*ListJobQueuesNotFound handles this case with default header values.

Not found
*/
type ListJobQueuesNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListJobQueuesNotFound) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesNotFound  %+v", 404, o.Payload)
}

func (o *ListJobQueuesNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListJobQueuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobQueuesInternalServerError creates a ListJobQueuesInternalServerError with default headers values
func NewListJobQueuesInternalServerError() *ListJobQueuesInternalServerError {
	return &ListJobQueuesInternalServerError{}
}

/*ListJobQueuesInternalServerError handles this case with default header values.

Internal server error
*/
type ListJobQueuesInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListJobQueuesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListJobQueuesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListJobQueuesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
