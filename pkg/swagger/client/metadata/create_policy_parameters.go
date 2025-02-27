// Code generated by go-swagger; DO NOT EDIT.

package metadata

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

	"github.com/fugue/regula/pkg/swagger/models"
)

// NewCreatePolicyParams creates a new CreatePolicyParams object
// with the default values initialized.
func NewCreatePolicyParams() *CreatePolicyParams {
	var ()
	return &CreatePolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePolicyParamsWithTimeout creates a new CreatePolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePolicyParamsWithTimeout(timeout time.Duration) *CreatePolicyParams {
	var ()
	return &CreatePolicyParams{

		timeout: timeout,
	}
}

// NewCreatePolicyParamsWithContext creates a new CreatePolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePolicyParamsWithContext(ctx context.Context) *CreatePolicyParams {
	var ()
	return &CreatePolicyParams{

		Context: ctx,
	}
}

// NewCreatePolicyParamsWithHTTPClient creates a new CreatePolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePolicyParamsWithHTTPClient(client *http.Client) *CreatePolicyParams {
	var ()
	return &CreatePolicyParams{
		HTTPClient: client,
	}
}

/*CreatePolicyParams contains all the parameters to send to the API endpoint
for the create policy operation typically these are written to a http.Request
*/
type CreatePolicyParams struct {

	/*Input
	  List of resource types to be able to survey and remediate.

	*/
	Input *models.CreatePolicyInput
	/*Provider
	  Name of the cloud provider.

	*/
	Provider string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create policy params
func (o *CreatePolicyParams) WithTimeout(timeout time.Duration) *CreatePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policy params
func (o *CreatePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policy params
func (o *CreatePolicyParams) WithContext(ctx context.Context) *CreatePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policy params
func (o *CreatePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policy params
func (o *CreatePolicyParams) WithHTTPClient(client *http.Client) *CreatePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policy params
func (o *CreatePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create policy params
func (o *CreatePolicyParams) WithInput(input *models.CreatePolicyInput) *CreatePolicyParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create policy params
func (o *CreatePolicyParams) SetInput(input *models.CreatePolicyInput) {
	o.Input = input
}

// WithProvider adds the provider to the create policy params
func (o *CreatePolicyParams) WithProvider(provider string) *CreatePolicyParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the create policy params
func (o *CreatePolicyParams) SetProvider(provider string) {
	o.Provider = provider
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
