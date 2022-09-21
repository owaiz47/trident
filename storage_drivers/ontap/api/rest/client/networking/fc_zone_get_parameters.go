// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewFcZoneGetParams creates a new FcZoneGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcZoneGetParams() *FcZoneGetParams {
	return &FcZoneGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcZoneGetParamsWithTimeout creates a new FcZoneGetParams object
// with the ability to set a timeout on a request.
func NewFcZoneGetParamsWithTimeout(timeout time.Duration) *FcZoneGetParams {
	return &FcZoneGetParams{
		timeout: timeout,
	}
}

// NewFcZoneGetParamsWithContext creates a new FcZoneGetParams object
// with the ability to set a context for a request.
func NewFcZoneGetParamsWithContext(ctx context.Context) *FcZoneGetParams {
	return &FcZoneGetParams{
		Context: ctx,
	}
}

// NewFcZoneGetParamsWithHTTPClient creates a new FcZoneGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcZoneGetParamsWithHTTPClient(client *http.Client) *FcZoneGetParams {
	return &FcZoneGetParams{
		HTTPClient: client,
	}
}

/* FcZoneGetParams contains all the parameters to send to the API endpoint
   for the fc zone get operation.

   Typically these are written to a http.Request.
*/
type FcZoneGetParams struct {

	/* CacheMaximumAge.

	   The maximum age of data in the Fibre Channel fabric cache before it should be refreshed from the fabric. The default is 15 minutes.

	   Format: iso8601
	   Default: "15 minutes"
	*/
	CacheMaximumAgeQueryParameter *string

	/* FabricName.

	   The WWN of the primary switch of the Fibre Channel fabric.

	*/
	FabricNamePathParameter string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Name.

	   The name of a zone in the active zoneset the Fibre Channel fabric.

	*/
	NamePathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fc zone get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcZoneGetParams) WithDefaults() *FcZoneGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fc zone get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcZoneGetParams) SetDefaults() {
	var (
		cacheMaximumAgeQueryParameterDefault = string("15 minutes")
	)

	val := FcZoneGetParams{
		CacheMaximumAgeQueryParameter: &cacheMaximumAgeQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fc zone get params
func (o *FcZoneGetParams) WithTimeout(timeout time.Duration) *FcZoneGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fc zone get params
func (o *FcZoneGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fc zone get params
func (o *FcZoneGetParams) WithContext(ctx context.Context) *FcZoneGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fc zone get params
func (o *FcZoneGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fc zone get params
func (o *FcZoneGetParams) WithHTTPClient(client *http.Client) *FcZoneGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fc zone get params
func (o *FcZoneGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheMaximumAgeQueryParameter adds the cacheMaximumAge to the fc zone get params
func (o *FcZoneGetParams) WithCacheMaximumAgeQueryParameter(cacheMaximumAge *string) *FcZoneGetParams {
	o.SetCacheMaximumAgeQueryParameter(cacheMaximumAge)
	return o
}

// SetCacheMaximumAgeQueryParameter adds the cacheMaximumAge to the fc zone get params
func (o *FcZoneGetParams) SetCacheMaximumAgeQueryParameter(cacheMaximumAge *string) {
	o.CacheMaximumAgeQueryParameter = cacheMaximumAge
}

// WithFabricNamePathParameter adds the fabricName to the fc zone get params
func (o *FcZoneGetParams) WithFabricNamePathParameter(fabricName string) *FcZoneGetParams {
	o.SetFabricNamePathParameter(fabricName)
	return o
}

// SetFabricNamePathParameter adds the fabricName to the fc zone get params
func (o *FcZoneGetParams) SetFabricNamePathParameter(fabricName string) {
	o.FabricNamePathParameter = fabricName
}

// WithFieldsQueryParameter adds the fields to the fc zone get params
func (o *FcZoneGetParams) WithFieldsQueryParameter(fields []string) *FcZoneGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the fc zone get params
func (o *FcZoneGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithNamePathParameter adds the name to the fc zone get params
func (o *FcZoneGetParams) WithNamePathParameter(name string) *FcZoneGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the fc zone get params
func (o *FcZoneGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WriteToRequest writes these params to a swagger request
func (o *FcZoneGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CacheMaximumAgeQueryParameter != nil {

		// query param cache.maximum_age
		var qrCacheMaximumAge string

		if o.CacheMaximumAgeQueryParameter != nil {
			qrCacheMaximumAge = *o.CacheMaximumAgeQueryParameter
		}
		qCacheMaximumAge := qrCacheMaximumAge
		if qCacheMaximumAge != "" {

			if err := r.SetQueryParam("cache.maximum_age", qCacheMaximumAge); err != nil {
				return err
			}
		}
	}

	// path param fabric.name
	if err := r.SetPathParam("fabric.name", o.FabricNamePathParameter); err != nil {
		return err
	}

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

// bindParamFcZoneGet binds the parameter fields
func (o *FcZoneGetParams) bindParamFields(formats strfmt.Registry) []string {
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
