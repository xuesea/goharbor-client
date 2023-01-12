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

// StopRunningJobReader is a Reader for the StopRunningJob structure.
type StopRunningJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopRunningJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopRunningJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopRunningJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopRunningJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopRunningJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopRunningJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopRunningJobOK creates a StopRunningJobOK with default headers values
func NewStopRunningJobOK() *StopRunningJobOK {
	return &StopRunningJobOK{}
}

/*StopRunningJobOK handles this case with default header values.

Stop worker successfully.
*/
type StopRunningJobOK struct {
}

func (o *StopRunningJobOK) Error() string {
	return fmt.Sprintf("[PUT /jobservice/jobs/{job_id}][%d] stopRunningJobOK ", 200)
}

func (o *StopRunningJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopRunningJobUnauthorized creates a StopRunningJobUnauthorized with default headers values
func NewStopRunningJobUnauthorized() *StopRunningJobUnauthorized {
	return &StopRunningJobUnauthorized{}
}

/*StopRunningJobUnauthorized handles this case with default header values.

Unauthorized
*/
type StopRunningJobUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopRunningJobUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /jobservice/jobs/{job_id}][%d] stopRunningJobUnauthorized  %+v", 401, o.Payload)
}

func (o *StopRunningJobUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopRunningJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopRunningJobForbidden creates a StopRunningJobForbidden with default headers values
func NewStopRunningJobForbidden() *StopRunningJobForbidden {
	return &StopRunningJobForbidden{}
}

/*StopRunningJobForbidden handles this case with default header values.

Forbidden
*/
type StopRunningJobForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopRunningJobForbidden) Error() string {
	return fmt.Sprintf("[PUT /jobservice/jobs/{job_id}][%d] stopRunningJobForbidden  %+v", 403, o.Payload)
}

func (o *StopRunningJobForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopRunningJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopRunningJobNotFound creates a StopRunningJobNotFound with default headers values
func NewStopRunningJobNotFound() *StopRunningJobNotFound {
	return &StopRunningJobNotFound{}
}

/*StopRunningJobNotFound handles this case with default header values.

Not found
*/
type StopRunningJobNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopRunningJobNotFound) Error() string {
	return fmt.Sprintf("[PUT /jobservice/jobs/{job_id}][%d] stopRunningJobNotFound  %+v", 404, o.Payload)
}

func (o *StopRunningJobNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopRunningJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopRunningJobInternalServerError creates a StopRunningJobInternalServerError with default headers values
func NewStopRunningJobInternalServerError() *StopRunningJobInternalServerError {
	return &StopRunningJobInternalServerError{}
}

/*StopRunningJobInternalServerError handles this case with default header values.

Internal server error
*/
type StopRunningJobInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopRunningJobInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /jobservice/jobs/{job_id}][%d] stopRunningJobInternalServerError  %+v", 500, o.Payload)
}

func (o *StopRunningJobInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopRunningJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
