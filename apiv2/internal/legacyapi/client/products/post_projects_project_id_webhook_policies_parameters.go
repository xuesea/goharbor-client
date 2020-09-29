// Code generated by go-swagger; DO NOT EDIT.

package products

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
	"github.com/go-openapi/swag"

	"github.com/mittwald/goharbor-client/v2/apiv2/model/legacy"
)

// NewPostProjectsProjectIDWebhookPoliciesParams creates a new PostProjectsProjectIDWebhookPoliciesParams object
// with the default values initialized.
func NewPostProjectsProjectIDWebhookPoliciesParams() *PostProjectsProjectIDWebhookPoliciesParams {
	var ()
	return &PostProjectsProjectIDWebhookPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProjectsProjectIDWebhookPoliciesParamsWithTimeout creates a new PostProjectsProjectIDWebhookPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProjectsProjectIDWebhookPoliciesParamsWithTimeout(timeout time.Duration) *PostProjectsProjectIDWebhookPoliciesParams {
	var ()
	return &PostProjectsProjectIDWebhookPoliciesParams{

		timeout: timeout,
	}
}

// NewPostProjectsProjectIDWebhookPoliciesParamsWithContext creates a new PostProjectsProjectIDWebhookPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProjectsProjectIDWebhookPoliciesParamsWithContext(ctx context.Context) *PostProjectsProjectIDWebhookPoliciesParams {
	var ()
	return &PostProjectsProjectIDWebhookPoliciesParams{

		Context: ctx,
	}
}

// NewPostProjectsProjectIDWebhookPoliciesParamsWithHTTPClient creates a new PostProjectsProjectIDWebhookPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProjectsProjectIDWebhookPoliciesParamsWithHTTPClient(client *http.Client) *PostProjectsProjectIDWebhookPoliciesParams {
	var ()
	return &PostProjectsProjectIDWebhookPoliciesParams{
		HTTPClient: client,
	}
}

/*PostProjectsProjectIDWebhookPoliciesParams contains all the parameters to send to the API endpoint
for the post projects project ID webhook policies operation typically these are written to a http.Request
*/
type PostProjectsProjectIDWebhookPoliciesParams struct {

	/*Policy
	  Properties "targets" and "event_types" needed.

	*/
	Policy *legacy.WebhookPolicy
	/*ProjectID
	  Relevant project ID

	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) WithTimeout(timeout time.Duration) *PostProjectsProjectIDWebhookPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) WithContext(ctx context.Context) *PostProjectsProjectIDWebhookPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) WithHTTPClient(client *http.Client) *PostProjectsProjectIDWebhookPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) WithPolicy(policy *legacy.WebhookPolicy) *PostProjectsProjectIDWebhookPoliciesParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) SetPolicy(policy *legacy.WebhookPolicy) {
	o.Policy = policy
}

// WithProjectID adds the projectID to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) WithProjectID(projectID int64) *PostProjectsProjectIDWebhookPoliciesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the post projects project ID webhook policies params
func (o *PostProjectsProjectIDWebhookPoliciesParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PostProjectsProjectIDWebhookPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
