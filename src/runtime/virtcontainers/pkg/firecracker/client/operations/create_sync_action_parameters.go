// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// NewCreateSyncActionParams creates a new CreateSyncActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSyncActionParams() *CreateSyncActionParams {
	return &CreateSyncActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSyncActionParamsWithTimeout creates a new CreateSyncActionParams object
// with the ability to set a timeout on a request.
func NewCreateSyncActionParamsWithTimeout(timeout time.Duration) *CreateSyncActionParams {
	return &CreateSyncActionParams{
		timeout: timeout,
	}
}

// NewCreateSyncActionParamsWithContext creates a new CreateSyncActionParams object
// with the ability to set a context for a request.
func NewCreateSyncActionParamsWithContext(ctx context.Context) *CreateSyncActionParams {
	return &CreateSyncActionParams{
		Context: ctx,
	}
}

// NewCreateSyncActionParamsWithHTTPClient creates a new CreateSyncActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSyncActionParamsWithHTTPClient(client *http.Client) *CreateSyncActionParams {
	return &CreateSyncActionParams{
		HTTPClient: client,
	}
}

/* CreateSyncActionParams contains all the parameters to send to the API endpoint
   for the create sync action operation.

   Typically these are written to a http.Request.
*/
type CreateSyncActionParams struct {

	// Info.
	Info *models.InstanceActionInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create sync action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSyncActionParams) WithDefaults() *CreateSyncActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create sync action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSyncActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create sync action params
func (o *CreateSyncActionParams) WithTimeout(timeout time.Duration) *CreateSyncActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create sync action params
func (o *CreateSyncActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create sync action params
func (o *CreateSyncActionParams) WithContext(ctx context.Context) *CreateSyncActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create sync action params
func (o *CreateSyncActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create sync action params
func (o *CreateSyncActionParams) WithHTTPClient(client *http.Client) *CreateSyncActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create sync action params
func (o *CreateSyncActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the create sync action params
func (o *CreateSyncActionParams) WithInfo(info *models.InstanceActionInfo) *CreateSyncActionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the create sync action params
func (o *CreateSyncActionParams) SetInfo(info *models.InstanceActionInfo) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSyncActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
