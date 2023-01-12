// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xuesea/goharbor-client/v5/apiv2/model"
)

// GetRegistryInfoReader is a Reader for the GetRegistryInfo structure.
type GetRegistryInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistryInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistryInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegistryInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegistryInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistryInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistryInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistryInfoOK creates a GetRegistryInfoOK with default headers values
func NewGetRegistryInfoOK() *GetRegistryInfoOK {
	return &GetRegistryInfoOK{}
}

/*GetRegistryInfoOK handles this case with default header values.

Success
*/
type GetRegistryInfoOK struct {
	Payload *model.RegistryInfo
}

func (o *GetRegistryInfoOK) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistryInfoOK  %+v", 200, o.Payload)
}

func (o *GetRegistryInfoOK) GetPayload() *model.RegistryInfo {
	return o.Payload
}

func (o *GetRegistryInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RegistryInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryInfoUnauthorized creates a GetRegistryInfoUnauthorized with default headers values
func NewGetRegistryInfoUnauthorized() *GetRegistryInfoUnauthorized {
	return &GetRegistryInfoUnauthorized{}
}

/*GetRegistryInfoUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRegistryInfoUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRegistryInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistryInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistryInfoUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRegistryInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryInfoForbidden creates a GetRegistryInfoForbidden with default headers values
func NewGetRegistryInfoForbidden() *GetRegistryInfoForbidden {
	return &GetRegistryInfoForbidden{}
}

/*GetRegistryInfoForbidden handles this case with default header values.

Forbidden
*/
type GetRegistryInfoForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRegistryInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistryInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetRegistryInfoForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRegistryInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryInfoNotFound creates a GetRegistryInfoNotFound with default headers values
func NewGetRegistryInfoNotFound() *GetRegistryInfoNotFound {
	return &GetRegistryInfoNotFound{}
}

/*GetRegistryInfoNotFound handles this case with default header values.

Not found
*/
type GetRegistryInfoNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRegistryInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistryInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistryInfoNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRegistryInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryInfoInternalServerError creates a GetRegistryInfoInternalServerError with default headers values
func NewGetRegistryInfoInternalServerError() *GetRegistryInfoInternalServerError {
	return &GetRegistryInfoInternalServerError{}
}

/*GetRegistryInfoInternalServerError handles this case with default header values.

Internal server error
*/
type GetRegistryInfoInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRegistryInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/info][%d] getRegistryInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistryInfoInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRegistryInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
