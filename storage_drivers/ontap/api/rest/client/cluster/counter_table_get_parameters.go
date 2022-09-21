// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewCounterTableGetParams creates a new CounterTableGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCounterTableGetParams() *CounterTableGetParams {
	return &CounterTableGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCounterTableGetParamsWithTimeout creates a new CounterTableGetParams object
// with the ability to set a timeout on a request.
func NewCounterTableGetParamsWithTimeout(timeout time.Duration) *CounterTableGetParams {
	return &CounterTableGetParams{
		timeout: timeout,
	}
}

// NewCounterTableGetParamsWithContext creates a new CounterTableGetParams object
// with the ability to set a context for a request.
func NewCounterTableGetParamsWithContext(ctx context.Context) *CounterTableGetParams {
	return &CounterTableGetParams{
		Context: ctx,
	}
}

// NewCounterTableGetParamsWithHTTPClient creates a new CounterTableGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCounterTableGetParamsWithHTTPClient(client *http.Client) *CounterTableGetParams {
	return &CounterTableGetParams{
		HTTPClient: client,
	}
}

/* CounterTableGetParams contains all the parameters to send to the API endpoint
   for the counter table get operation.

   Typically these are written to a http.Request.
*/
type CounterTableGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Name.

	   Counter table name.
	*/
	NamePathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the counter table get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CounterTableGetParams) WithDefaults() *CounterTableGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the counter table get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CounterTableGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the counter table get params
func (o *CounterTableGetParams) WithTimeout(timeout time.Duration) *CounterTableGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the counter table get params
func (o *CounterTableGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the counter table get params
func (o *CounterTableGetParams) WithContext(ctx context.Context) *CounterTableGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the counter table get params
func (o *CounterTableGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the counter table get params
func (o *CounterTableGetParams) WithHTTPClient(client *http.Client) *CounterTableGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the counter table get params
func (o *CounterTableGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the counter table get params
func (o *CounterTableGetParams) WithFieldsQueryParameter(fields []string) *CounterTableGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the counter table get params
func (o *CounterTableGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithNamePathParameter adds the name to the counter table get params
func (o *CounterTableGetParams) WithNamePathParameter(name string) *CounterTableGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the counter table get params
func (o *CounterTableGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WriteToRequest writes these params to a swagger request
func (o *CounterTableGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCounterTableGet binds the parameter fields
func (o *CounterTableGetParams) bindParamFields(formats strfmt.Registry) []string {
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
