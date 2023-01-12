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

// GetPurgeScheduleReader is a Reader for the GetPurgeSchedule structure.
type GetPurgeScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPurgeScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPurgeScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPurgeScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPurgeScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPurgeScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPurgeScheduleOK creates a GetPurgeScheduleOK with default headers values
func NewGetPurgeScheduleOK() *GetPurgeScheduleOK {
	return &GetPurgeScheduleOK{}
}

/*GetPurgeScheduleOK handles this case with default header values.

Get purge job's schedule.
*/
type GetPurgeScheduleOK struct {
	Payload *model.ExecHistory
}

func (o *GetPurgeScheduleOK) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/schedule][%d] getPurgeScheduleOK  %+v", 200, o.Payload)
}

func (o *GetPurgeScheduleOK) GetPayload() *model.ExecHistory {
	return o.Payload
}

func (o *GetPurgeScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ExecHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeScheduleUnauthorized creates a GetPurgeScheduleUnauthorized with default headers values
func NewGetPurgeScheduleUnauthorized() *GetPurgeScheduleUnauthorized {
	return &GetPurgeScheduleUnauthorized{}
}

/*GetPurgeScheduleUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPurgeScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/schedule][%d] getPurgeScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPurgeScheduleUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeScheduleForbidden creates a GetPurgeScheduleForbidden with default headers values
func NewGetPurgeScheduleForbidden() *GetPurgeScheduleForbidden {
	return &GetPurgeScheduleForbidden{}
}

/*GetPurgeScheduleForbidden handles this case with default header values.

Forbidden
*/
type GetPurgeScheduleForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/schedule][%d] getPurgeScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetPurgeScheduleForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeScheduleInternalServerError creates a GetPurgeScheduleInternalServerError with default headers values
func NewGetPurgeScheduleInternalServerError() *GetPurgeScheduleInternalServerError {
	return &GetPurgeScheduleInternalServerError{}
}

/*GetPurgeScheduleInternalServerError handles this case with default header values.

Internal server error
*/
type GetPurgeScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit/schedule][%d] getPurgeScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPurgeScheduleInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
