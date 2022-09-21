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

// NewCifsOpenFileGetParams creates a new CifsOpenFileGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsOpenFileGetParams() *CifsOpenFileGetParams {
	return &CifsOpenFileGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsOpenFileGetParamsWithTimeout creates a new CifsOpenFileGetParams object
// with the ability to set a timeout on a request.
func NewCifsOpenFileGetParamsWithTimeout(timeout time.Duration) *CifsOpenFileGetParams {
	return &CifsOpenFileGetParams{
		timeout: timeout,
	}
}

// NewCifsOpenFileGetParamsWithContext creates a new CifsOpenFileGetParams object
// with the ability to set a context for a request.
func NewCifsOpenFileGetParamsWithContext(ctx context.Context) *CifsOpenFileGetParams {
	return &CifsOpenFileGetParams{
		Context: ctx,
	}
}

// NewCifsOpenFileGetParamsWithHTTPClient creates a new CifsOpenFileGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsOpenFileGetParamsWithHTTPClient(client *http.Client) *CifsOpenFileGetParams {
	return &CifsOpenFileGetParams{
		HTTPClient: client,
	}
}

/* CifsOpenFileGetParams contains all the parameters to send to the API endpoint
   for the cifs open file get operation.

   Typically these are written to a http.Request.
*/
type CifsOpenFileGetParams struct {

	/* ConnectionIdentifier.

	   Connection ID
	*/
	ConnectionIDentifierPathParameter int64

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Identifier.

	   File ID

	   Format: int64
	*/
	IdentifierPathParameter int64

	/* NodeUUID.

	   Node UUID.
	*/
	NodeUUIDPathParameter string

	/* SessionIdentifier.

	   Session ID

	   Format: int64
	*/
	SessionIDentifierPathParameter int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs open file get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsOpenFileGetParams) WithDefaults() *CifsOpenFileGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs open file get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsOpenFileGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs open file get params
func (o *CifsOpenFileGetParams) WithTimeout(timeout time.Duration) *CifsOpenFileGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs open file get params
func (o *CifsOpenFileGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs open file get params
func (o *CifsOpenFileGetParams) WithContext(ctx context.Context) *CifsOpenFileGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs open file get params
func (o *CifsOpenFileGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs open file get params
func (o *CifsOpenFileGetParams) WithHTTPClient(client *http.Client) *CifsOpenFileGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs open file get params
func (o *CifsOpenFileGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionIDentifierPathParameter adds the connectionIdentifier to the cifs open file get params
func (o *CifsOpenFileGetParams) WithConnectionIDentifierPathParameter(connectionIdentifier int64) *CifsOpenFileGetParams {
	o.SetConnectionIDentifierPathParameter(connectionIdentifier)
	return o
}

// SetConnectionIDentifierPathParameter adds the connectionIdentifier to the cifs open file get params
func (o *CifsOpenFileGetParams) SetConnectionIDentifierPathParameter(connectionIdentifier int64) {
	o.ConnectionIDentifierPathParameter = connectionIdentifier
}

// WithFieldsQueryParameter adds the fields to the cifs open file get params
func (o *CifsOpenFileGetParams) WithFieldsQueryParameter(fields []string) *CifsOpenFileGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cifs open file get params
func (o *CifsOpenFileGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIdentifierPathParameter adds the identifier to the cifs open file get params
func (o *CifsOpenFileGetParams) WithIdentifierPathParameter(identifier int64) *CifsOpenFileGetParams {
	o.SetIdentifierPathParameter(identifier)
	return o
}

// SetIdentifierPathParameter adds the identifier to the cifs open file get params
func (o *CifsOpenFileGetParams) SetIdentifierPathParameter(identifier int64) {
	o.IdentifierPathParameter = identifier
}

// WithNodeUUIDPathParameter adds the nodeUUID to the cifs open file get params
func (o *CifsOpenFileGetParams) WithNodeUUIDPathParameter(nodeUUID string) *CifsOpenFileGetParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the cifs open file get params
func (o *CifsOpenFileGetParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WithSessionIDentifierPathParameter adds the sessionIdentifier to the cifs open file get params
func (o *CifsOpenFileGetParams) WithSessionIDentifierPathParameter(sessionIdentifier int64) *CifsOpenFileGetParams {
	o.SetSessionIDentifierPathParameter(sessionIdentifier)
	return o
}

// SetSessionIDentifierPathParameter adds the sessionIdentifier to the cifs open file get params
func (o *CifsOpenFileGetParams) SetSessionIDentifierPathParameter(sessionIdentifier int64) {
	o.SessionIDentifierPathParameter = sessionIdentifier
}

// WithSVMUUIDPathParameter adds the svmUUID to the cifs open file get params
func (o *CifsOpenFileGetParams) WithSVMUUIDPathParameter(svmUUID string) *CifsOpenFileGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the cifs open file get params
func (o *CifsOpenFileGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsOpenFileGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection.identifier
	if err := r.SetPathParam("connection.identifier", swag.FormatInt64(o.ConnectionIDentifierPathParameter)); err != nil {
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

	// path param identifier
	if err := r.SetPathParam("identifier", swag.FormatInt64(o.IdentifierPathParameter)); err != nil {
		return err
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
		return err
	}

	// path param session.identifier
	if err := r.SetPathParam("session.identifier", swag.FormatInt64(o.SessionIDentifierPathParameter)); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsOpenFileGet binds the parameter fields
func (o *CifsOpenFileGetParams) bindParamFields(formats strfmt.Registry) []string {
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
