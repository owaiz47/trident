// Code generated by go-swagger; DO NOT EDIT.

package security

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
)

// NewSecurityKeyManagerKeyServersDeleteParams creates a new SecurityKeyManagerKeyServersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeyManagerKeyServersDeleteParams() *SecurityKeyManagerKeyServersDeleteParams {
	return &SecurityKeyManagerKeyServersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeyManagerKeyServersDeleteParamsWithTimeout creates a new SecurityKeyManagerKeyServersDeleteParams object
// with the ability to set a timeout on a request.
func NewSecurityKeyManagerKeyServersDeleteParamsWithTimeout(timeout time.Duration) *SecurityKeyManagerKeyServersDeleteParams {
	return &SecurityKeyManagerKeyServersDeleteParams{
		timeout: timeout,
	}
}

// NewSecurityKeyManagerKeyServersDeleteParamsWithContext creates a new SecurityKeyManagerKeyServersDeleteParams object
// with the ability to set a context for a request.
func NewSecurityKeyManagerKeyServersDeleteParamsWithContext(ctx context.Context) *SecurityKeyManagerKeyServersDeleteParams {
	return &SecurityKeyManagerKeyServersDeleteParams{
		Context: ctx,
	}
}

// NewSecurityKeyManagerKeyServersDeleteParamsWithHTTPClient creates a new SecurityKeyManagerKeyServersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeyManagerKeyServersDeleteParamsWithHTTPClient(client *http.Client) *SecurityKeyManagerKeyServersDeleteParams {
	return &SecurityKeyManagerKeyServersDeleteParams{
		HTTPClient: client,
	}
}

/* SecurityKeyManagerKeyServersDeleteParams contains all the parameters to send to the API endpoint
   for the security key manager key servers delete operation.

   Typically these are written to a http.Request.
*/
type SecurityKeyManagerKeyServersDeleteParams struct {

	/* Force.

	   Set the force flag to "true" to bypass out of quorum checks when removing a primary key server.

	*/
	ForceQueryParameter *bool

	/* Server.

	   Primary key server configured in the external key manager.
	*/
	ServerPathParameter string

	/* UUID.

	   External key manager UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security key manager key servers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerKeyServersDeleteParams) WithDefaults() *SecurityKeyManagerKeyServersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security key manager key servers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerKeyServersDeleteParams) SetDefaults() {
	var (
		forceQueryParameterDefault = bool(false)
	)

	val := SecurityKeyManagerKeyServersDeleteParams{
		ForceQueryParameter: &forceQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) WithTimeout(timeout time.Duration) *SecurityKeyManagerKeyServersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) WithContext(ctx context.Context) *SecurityKeyManagerKeyServersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) WithHTTPClient(client *http.Client) *SecurityKeyManagerKeyServersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceQueryParameter adds the force to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) WithForceQueryParameter(force *bool) *SecurityKeyManagerKeyServersDeleteParams {
	o.SetForceQueryParameter(force)
	return o
}

// SetForceQueryParameter adds the force to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) SetForceQueryParameter(force *bool) {
	o.ForceQueryParameter = force
}

// WithServerPathParameter adds the server to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) WithServerPathParameter(server string) *SecurityKeyManagerKeyServersDeleteParams {
	o.SetServerPathParameter(server)
	return o
}

// SetServerPathParameter adds the server to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) SetServerPathParameter(server string) {
	o.ServerPathParameter = server
}

// WithUUIDPathParameter adds the uuid to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) WithUUIDPathParameter(uuid string) *SecurityKeyManagerKeyServersDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the security key manager key servers delete params
func (o *SecurityKeyManagerKeyServersDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeyManagerKeyServersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceQueryParameter != nil {

		// query param force
		var qrForce bool

		if o.ForceQueryParameter != nil {
			qrForce = *o.ForceQueryParameter
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param server
	if err := r.SetPathParam("server", o.ServerPathParameter); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
