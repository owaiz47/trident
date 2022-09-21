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

// NewMultiAdminVerifyApprovalGroupGetParams creates a new MultiAdminVerifyApprovalGroupGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyApprovalGroupGetParams() *MultiAdminVerifyApprovalGroupGetParams {
	return &MultiAdminVerifyApprovalGroupGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyApprovalGroupGetParamsWithTimeout creates a new MultiAdminVerifyApprovalGroupGetParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyApprovalGroupGetParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyApprovalGroupGetParams {
	return &MultiAdminVerifyApprovalGroupGetParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyApprovalGroupGetParamsWithContext creates a new MultiAdminVerifyApprovalGroupGetParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyApprovalGroupGetParamsWithContext(ctx context.Context) *MultiAdminVerifyApprovalGroupGetParams {
	return &MultiAdminVerifyApprovalGroupGetParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyApprovalGroupGetParamsWithHTTPClient creates a new MultiAdminVerifyApprovalGroupGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyApprovalGroupGetParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyApprovalGroupGetParams {
	return &MultiAdminVerifyApprovalGroupGetParams{
		HTTPClient: client,
	}
}

/* MultiAdminVerifyApprovalGroupGetParams contains all the parameters to send to the API endpoint
   for the multi admin verify approval group get operation.

   Typically these are written to a http.Request.
*/
type MultiAdminVerifyApprovalGroupGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	// Name.
	NamePathParameter string

	// OwnerUUID.
	OwnerUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify approval group get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyApprovalGroupGetParams) WithDefaults() *MultiAdminVerifyApprovalGroupGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify approval group get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyApprovalGroupGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyApprovalGroupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) WithContext(ctx context.Context) *MultiAdminVerifyApprovalGroupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyApprovalGroupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) WithFieldsQueryParameter(fields []string) *MultiAdminVerifyApprovalGroupGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithNamePathParameter adds the name to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) WithNamePathParameter(name string) *MultiAdminVerifyApprovalGroupGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithOwnerUUIDPathParameter adds the ownerUUID to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) WithOwnerUUIDPathParameter(ownerUUID string) *MultiAdminVerifyApprovalGroupGetParams {
	o.SetOwnerUUIDPathParameter(ownerUUID)
	return o
}

// SetOwnerUUIDPathParameter adds the ownerUuid to the multi admin verify approval group get params
func (o *MultiAdminVerifyApprovalGroupGetParams) SetOwnerUUIDPathParameter(ownerUUID string) {
	o.OwnerUUIDPathParameter = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyApprovalGroupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamMultiAdminVerifyApprovalGroupGet binds the parameter fields
func (o *MultiAdminVerifyApprovalGroupGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
