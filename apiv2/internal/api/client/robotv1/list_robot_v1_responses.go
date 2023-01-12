// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/xuesea/goharbor-client/v5/apiv2/model"
)

// ListRobotV1Reader is a Reader for the ListRobotV1 structure.
type ListRobotV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRobotV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRobotV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListRobotV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRobotV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRobotV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRobotV1OK creates a ListRobotV1OK with default headers values
func NewListRobotV1OK() *ListRobotV1OK {
	return &ListRobotV1OK{}
}

/*ListRobotV1OK handles this case with default header values.

Success
*/
type ListRobotV1OK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of robot accounts
	 */
	XTotalCount int64

	Payload []*model.Robot
}

func (o *ListRobotV1OK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots][%d] listRobotV1OK  %+v", 200, o.Payload)
}

func (o *ListRobotV1OK) GetPayload() []*model.Robot {
	return o.Payload
}

func (o *ListRobotV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotV1BadRequest creates a ListRobotV1BadRequest with default headers values
func NewListRobotV1BadRequest() *ListRobotV1BadRequest {
	return &ListRobotV1BadRequest{}
}

/*ListRobotV1BadRequest handles this case with default header values.

Bad request
*/
type ListRobotV1BadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRobotV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots][%d] listRobotV1BadRequest  %+v", 400, o.Payload)
}

func (o *ListRobotV1BadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRobotV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotV1NotFound creates a ListRobotV1NotFound with default headers values
func NewListRobotV1NotFound() *ListRobotV1NotFound {
	return &ListRobotV1NotFound{}
}

/*ListRobotV1NotFound handles this case with default header values.

Not found
*/
type ListRobotV1NotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRobotV1NotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots][%d] listRobotV1NotFound  %+v", 404, o.Payload)
}

func (o *ListRobotV1NotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRobotV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotV1InternalServerError creates a ListRobotV1InternalServerError with default headers values
func NewListRobotV1InternalServerError() *ListRobotV1InternalServerError {
	return &ListRobotV1InternalServerError{}
}

/*ListRobotV1InternalServerError handles this case with default header values.

Internal server error
*/
type ListRobotV1InternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRobotV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots][%d] listRobotV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ListRobotV1InternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRobotV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
