// Code generated by go-swagger; DO NOT EDIT.

package purge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xuesea/goharbor-client/v5/apiv2/model"
)

// CreatePurgeScheduleReader is a Reader for the CreatePurgeSchedule structure.
type CreatePurgeScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePurgeScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePurgeScheduleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePurgeScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreatePurgeScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePurgeScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePurgeScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePurgeScheduleCreated creates a CreatePurgeScheduleCreated with default headers values
func NewCreatePurgeScheduleCreated() *CreatePurgeScheduleCreated {
	return &CreatePurgeScheduleCreated{}
}

/*CreatePurgeScheduleCreated handles this case with default header values.

Created
*/
type CreatePurgeScheduleCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreatePurgeScheduleCreated) Error() string {
	return fmt.Sprintf("[POST /system/purgeaudit/schedule][%d] createPurgeScheduleCreated ", 201)
}

func (o *CreatePurgeScheduleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewCreatePurgeScheduleBadRequest creates a CreatePurgeScheduleBadRequest with default headers values
func NewCreatePurgeScheduleBadRequest() *CreatePurgeScheduleBadRequest {
	return &CreatePurgeScheduleBadRequest{}
}

/*CreatePurgeScheduleBadRequest handles this case with default header values.

Bad request
*/
type CreatePurgeScheduleBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreatePurgeScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /system/purgeaudit/schedule][%d] createPurgeScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePurgeScheduleBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreatePurgeScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePurgeScheduleUnauthorized creates a CreatePurgeScheduleUnauthorized with default headers values
func NewCreatePurgeScheduleUnauthorized() *CreatePurgeScheduleUnauthorized {
	return &CreatePurgeScheduleUnauthorized{}
}

/*CreatePurgeScheduleUnauthorized handles this case with default header values.

Unauthorized
*/
type CreatePurgeScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreatePurgeScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /system/purgeaudit/schedule][%d] createPurgeScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *CreatePurgeScheduleUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreatePurgeScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePurgeScheduleForbidden creates a CreatePurgeScheduleForbidden with default headers values
func NewCreatePurgeScheduleForbidden() *CreatePurgeScheduleForbidden {
	return &CreatePurgeScheduleForbidden{}
}

/*CreatePurgeScheduleForbidden handles this case with default header values.

Forbidden
*/
type CreatePurgeScheduleForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreatePurgeScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /system/purgeaudit/schedule][%d] createPurgeScheduleForbidden  %+v", 403, o.Payload)
}

func (o *CreatePurgeScheduleForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreatePurgeScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePurgeScheduleInternalServerError creates a CreatePurgeScheduleInternalServerError with default headers values
func NewCreatePurgeScheduleInternalServerError() *CreatePurgeScheduleInternalServerError {
	return &CreatePurgeScheduleInternalServerError{}
}

/*CreatePurgeScheduleInternalServerError handles this case with default header values.

Internal server error
*/
type CreatePurgeScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreatePurgeScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /system/purgeaudit/schedule][%d] createPurgeScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePurgeScheduleInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreatePurgeScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
