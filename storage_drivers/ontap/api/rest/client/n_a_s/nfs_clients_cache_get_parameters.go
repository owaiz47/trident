// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewNfsClientsCacheGetParams creates a new NfsClientsCacheGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNfsClientsCacheGetParams() *NfsClientsCacheGetParams {
	return &NfsClientsCacheGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNfsClientsCacheGetParamsWithTimeout creates a new NfsClientsCacheGetParams object
// with the ability to set a timeout on a request.
func NewNfsClientsCacheGetParamsWithTimeout(timeout time.Duration) *NfsClientsCacheGetParams {
	return &NfsClientsCacheGetParams{
		timeout: timeout,
	}
}

// NewNfsClientsCacheGetParamsWithContext creates a new NfsClientsCacheGetParams object
// with the ability to set a context for a request.
func NewNfsClientsCacheGetParamsWithContext(ctx context.Context) *NfsClientsCacheGetParams {
	return &NfsClientsCacheGetParams{
		Context: ctx,
	}
}

// NewNfsClientsCacheGetParamsWithHTTPClient creates a new NfsClientsCacheGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNfsClientsCacheGetParamsWithHTTPClient(client *http.Client) *NfsClientsCacheGetParams {
	return &NfsClientsCacheGetParams{
		HTTPClient: client,
	}
}

/*
NfsClientsCacheGetParams contains all the parameters to send to the API endpoint

	for the nfs clients cache get operation.

	Typically these are written to a http.Request.
*/
type NfsClientsCacheGetParams struct {

	/* ClientRetentionInterval.

	   Filter by client_retention_interval
	*/
	ClientRetentionInterval *string

	/* EnableNfsClientsDeletion.

	   Filter by enable_nfs_clients_deletion
	*/
	EnableNfsClientsDeletion *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UpdateInterval.

	   Filter by update_interval
	*/
	UpdateInterval *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nfs clients cache get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsClientsCacheGetParams) WithDefaults() *NfsClientsCacheGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nfs clients cache get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsClientsCacheGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) WithTimeout(timeout time.Duration) *NfsClientsCacheGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) WithContext(ctx context.Context) *NfsClientsCacheGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) WithHTTPClient(client *http.Client) *NfsClientsCacheGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientRetentionInterval adds the clientRetentionInterval to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) WithClientRetentionInterval(clientRetentionInterval *string) *NfsClientsCacheGetParams {
	o.SetClientRetentionInterval(clientRetentionInterval)
	return o
}

// SetClientRetentionInterval adds the clientRetentionInterval to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) SetClientRetentionInterval(clientRetentionInterval *string) {
	o.ClientRetentionInterval = clientRetentionInterval
}

// WithEnableNfsClientsDeletion adds the enableNfsClientsDeletion to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) WithEnableNfsClientsDeletion(enableNfsClientsDeletion *bool) *NfsClientsCacheGetParams {
	o.SetEnableNfsClientsDeletion(enableNfsClientsDeletion)
	return o
}

// SetEnableNfsClientsDeletion adds the enableNfsClientsDeletion to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) SetEnableNfsClientsDeletion(enableNfsClientsDeletion *bool) {
	o.EnableNfsClientsDeletion = enableNfsClientsDeletion
}

// WithFields adds the fields to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) WithFields(fields []string) *NfsClientsCacheGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUpdateInterval adds the updateInterval to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) WithUpdateInterval(updateInterval *string) *NfsClientsCacheGetParams {
	o.SetUpdateInterval(updateInterval)
	return o
}

// SetUpdateInterval adds the updateInterval to the nfs clients cache get params
func (o *NfsClientsCacheGetParams) SetUpdateInterval(updateInterval *string) {
	o.UpdateInterval = updateInterval
}

// WriteToRequest writes these params to a swagger request
func (o *NfsClientsCacheGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientRetentionInterval != nil {

		// query param client_retention_interval
		var qrClientRetentionInterval string

		if o.ClientRetentionInterval != nil {
			qrClientRetentionInterval = *o.ClientRetentionInterval
		}
		qClientRetentionInterval := qrClientRetentionInterval
		if qClientRetentionInterval != "" {

			if err := r.SetQueryParam("client_retention_interval", qClientRetentionInterval); err != nil {
				return err
			}
		}
	}

	if o.EnableNfsClientsDeletion != nil {

		// query param enable_nfs_clients_deletion
		var qrEnableNfsClientsDeletion bool

		if o.EnableNfsClientsDeletion != nil {
			qrEnableNfsClientsDeletion = *o.EnableNfsClientsDeletion
		}
		qEnableNfsClientsDeletion := swag.FormatBool(qrEnableNfsClientsDeletion)
		if qEnableNfsClientsDeletion != "" {

			if err := r.SetQueryParam("enable_nfs_clients_deletion", qEnableNfsClientsDeletion); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.UpdateInterval != nil {

		// query param update_interval
		var qrUpdateInterval string

		if o.UpdateInterval != nil {
			qrUpdateInterval = *o.UpdateInterval
		}
		qUpdateInterval := qrUpdateInterval
		if qUpdateInterval != "" {

			if err := r.SetQueryParam("update_interval", qUpdateInterval); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNfsClientsCacheGet binds the parameter fields
func (o *NfsClientsCacheGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}