// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xuesea/goharbor-client/v5/apiv2/model"
)

// DownloadScanDataReader is a Reader for the DownloadScanData structure.
type DownloadScanDataReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadScanDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadScanDataOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDownloadScanDataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadScanDataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadScanDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadScanDataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadScanDataOK creates a DownloadScanDataOK with default headers values
func NewDownloadScanDataOK(writer io.Writer) *DownloadScanDataOK {
	return &DownloadScanDataOK{
		Payload: writer,
	}
}

/*DownloadScanDataOK handles this case with default header values.

Data file containing the export data
*/
type DownloadScanDataOK struct {
	/*Value is a CSV formatted file; filename=export.csv
	 */
	ContentDisposition string

	Payload io.Writer
}

func (o *DownloadScanDataOK) Error() string {
	return fmt.Sprintf("[GET /export/cve/download/{execution_id}][%d] downloadScanDataOK  %+v", 200, o.Payload)
}

func (o *DownloadScanDataOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadScanDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Disposition
	o.ContentDisposition = response.GetHeader("Content-Disposition")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadScanDataUnauthorized creates a DownloadScanDataUnauthorized with default headers values
func NewDownloadScanDataUnauthorized() *DownloadScanDataUnauthorized {
	return &DownloadScanDataUnauthorized{}
}

/*DownloadScanDataUnauthorized handles this case with default header values.

Unauthorized
*/
type DownloadScanDataUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DownloadScanDataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /export/cve/download/{execution_id}][%d] downloadScanDataUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadScanDataUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DownloadScanDataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadScanDataForbidden creates a DownloadScanDataForbidden with default headers values
func NewDownloadScanDataForbidden() *DownloadScanDataForbidden {
	return &DownloadScanDataForbidden{}
}

/*DownloadScanDataForbidden handles this case with default header values.

Forbidden
*/
type DownloadScanDataForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DownloadScanDataForbidden) Error() string {
	return fmt.Sprintf("[GET /export/cve/download/{execution_id}][%d] downloadScanDataForbidden  %+v", 403, o.Payload)
}

func (o *DownloadScanDataForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DownloadScanDataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadScanDataNotFound creates a DownloadScanDataNotFound with default headers values
func NewDownloadScanDataNotFound() *DownloadScanDataNotFound {
	return &DownloadScanDataNotFound{}
}

/*DownloadScanDataNotFound handles this case with default header values.

Not found
*/
type DownloadScanDataNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DownloadScanDataNotFound) Error() string {
	return fmt.Sprintf("[GET /export/cve/download/{execution_id}][%d] downloadScanDataNotFound  %+v", 404, o.Payload)
}

func (o *DownloadScanDataNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DownloadScanDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadScanDataInternalServerError creates a DownloadScanDataInternalServerError with default headers values
func NewDownloadScanDataInternalServerError() *DownloadScanDataInternalServerError {
	return &DownloadScanDataInternalServerError{}
}

/*DownloadScanDataInternalServerError handles this case with default header values.

Internal server error
*/
type DownloadScanDataInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DownloadScanDataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /export/cve/download/{execution_id}][%d] downloadScanDataInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadScanDataInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DownloadScanDataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
