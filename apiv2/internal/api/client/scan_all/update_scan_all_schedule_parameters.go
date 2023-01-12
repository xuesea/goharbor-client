// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/xuesea/goharbor-client/v5/apiv2/model"
)

// NewUpdateScanAllScheduleParams creates a new UpdateScanAllScheduleParams object
// with the default values initialized.
func NewUpdateScanAllScheduleParams() *UpdateScanAllScheduleParams {
	var ()
	return &UpdateScanAllScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScanAllScheduleParamsWithTimeout creates a new UpdateScanAllScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateScanAllScheduleParamsWithTimeout(timeout time.Duration) *UpdateScanAllScheduleParams {
	var ()
	return &UpdateScanAllScheduleParams{

		timeout: timeout,
	}
}

// NewUpdateScanAllScheduleParamsWithContext creates a new UpdateScanAllScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateScanAllScheduleParamsWithContext(ctx context.Context) *UpdateScanAllScheduleParams {
	var ()
	return &UpdateScanAllScheduleParams{

		Context: ctx,
	}
}

// NewUpdateScanAllScheduleParamsWithHTTPClient creates a new UpdateScanAllScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateScanAllScheduleParamsWithHTTPClient(client *http.Client) *UpdateScanAllScheduleParams {
	var ()
	return &UpdateScanAllScheduleParams{
		HTTPClient: client,
	}
}

/*UpdateScanAllScheduleParams contains all the parameters to send to the API endpoint
for the update scan all schedule operation typically these are written to a http.Request
*/
type UpdateScanAllScheduleParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Schedule
	  Updates the schedule of scan all job, which scans all of images in Harbor.

	*/
	Schedule *model.Schedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithTimeout(timeout time.Duration) *UpdateScanAllScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithContext(ctx context.Context) *UpdateScanAllScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithHTTPClient(client *http.Client) *UpdateScanAllScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithXRequestID(xRequestID *string) *UpdateScanAllScheduleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSchedule adds the schedule to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) WithSchedule(schedule *model.Schedule) *UpdateScanAllScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the update scan all schedule params
func (o *UpdateScanAllScheduleParams) SetSchedule(schedule *model.Schedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScanAllScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
