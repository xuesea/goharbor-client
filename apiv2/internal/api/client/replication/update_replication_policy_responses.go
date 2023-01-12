// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xuesea/goharbor-client/v5/apiv2/model"
)

// UpdateReplicationPolicyReader is a Reader for the UpdateReplicationPolicy structure.
type UpdateReplicationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReplicationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateReplicationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateReplicationPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateReplicationPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateReplicationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateReplicationPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateReplicationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateReplicationPolicyOK creates a UpdateReplicationPolicyOK with default headers values
func NewUpdateReplicationPolicyOK() *UpdateReplicationPolicyOK {
	return &UpdateReplicationPolicyOK{}
}

/*UpdateReplicationPolicyOK handles this case with default header values.

Success
*/
type UpdateReplicationPolicyOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateReplicationPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyOK ", 200)
}

func (o *UpdateReplicationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateReplicationPolicyUnauthorized creates a UpdateReplicationPolicyUnauthorized with default headers values
func NewUpdateReplicationPolicyUnauthorized() *UpdateReplicationPolicyUnauthorized {
	return &UpdateReplicationPolicyUnauthorized{}
}

/*UpdateReplicationPolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateReplicationPolicyUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateReplicationPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateReplicationPolicyUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReplicationPolicyForbidden creates a UpdateReplicationPolicyForbidden with default headers values
func NewUpdateReplicationPolicyForbidden() *UpdateReplicationPolicyForbidden {
	return &UpdateReplicationPolicyForbidden{}
}

/*UpdateReplicationPolicyForbidden handles this case with default header values.

Forbidden
*/
type UpdateReplicationPolicyForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateReplicationPolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *UpdateReplicationPolicyForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReplicationPolicyNotFound creates a UpdateReplicationPolicyNotFound with default headers values
func NewUpdateReplicationPolicyNotFound() *UpdateReplicationPolicyNotFound {
	return &UpdateReplicationPolicyNotFound{}
}

/*UpdateReplicationPolicyNotFound handles this case with default header values.

Not found
*/
type UpdateReplicationPolicyNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateReplicationPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateReplicationPolicyNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReplicationPolicyConflict creates a UpdateReplicationPolicyConflict with default headers values
func NewUpdateReplicationPolicyConflict() *UpdateReplicationPolicyConflict {
	return &UpdateReplicationPolicyConflict{}
}

/*UpdateReplicationPolicyConflict handles this case with default header values.

Conflict
*/
type UpdateReplicationPolicyConflict struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateReplicationPolicyConflict) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyConflict  %+v", 409, o.Payload)
}

func (o *UpdateReplicationPolicyConflict) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReplicationPolicyInternalServerError creates a UpdateReplicationPolicyInternalServerError with default headers values
func NewUpdateReplicationPolicyInternalServerError() *UpdateReplicationPolicyInternalServerError {
	return &UpdateReplicationPolicyInternalServerError{}
}

/*UpdateReplicationPolicyInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateReplicationPolicyInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateReplicationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /replication/policies/{id}][%d] updateReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateReplicationPolicyInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateReplicationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
