// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewEmsMessageCollectionGetParams creates a new EmsMessageCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsMessageCollectionGetParams() *EmsMessageCollectionGetParams {
	return &EmsMessageCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsMessageCollectionGetParamsWithTimeout creates a new EmsMessageCollectionGetParams object
// with the ability to set a timeout on a request.
func NewEmsMessageCollectionGetParamsWithTimeout(timeout time.Duration) *EmsMessageCollectionGetParams {
	return &EmsMessageCollectionGetParams{
		timeout: timeout,
	}
}

// NewEmsMessageCollectionGetParamsWithContext creates a new EmsMessageCollectionGetParams object
// with the ability to set a context for a request.
func NewEmsMessageCollectionGetParamsWithContext(ctx context.Context) *EmsMessageCollectionGetParams {
	return &EmsMessageCollectionGetParams{
		Context: ctx,
	}
}

// NewEmsMessageCollectionGetParamsWithHTTPClient creates a new EmsMessageCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsMessageCollectionGetParamsWithHTTPClient(client *http.Client) *EmsMessageCollectionGetParams {
	return &EmsMessageCollectionGetParams{
		HTTPClient: client,
	}
}

/* EmsMessageCollectionGetParams contains all the parameters to send to the API endpoint
   for the ems message collection get operation.

   Typically these are written to a http.Request.
*/
type EmsMessageCollectionGetParams struct {

	/* CorrectiveAction.

	   Filter by corrective_action
	*/
	CorrectiveActionQueryParameter *string

	/* Deprecated.

	   Filter by deprecated
	*/
	DeprecatedQueryParameter *bool

	/* Description.

	   Filter by description
	*/
	DescriptionQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* Severity.

	   Filter by severity
	*/
	SeverityQueryParameter *string

	/* SnmpTrapType.

	   Filter by snmp_trap_type
	*/
	SnmpTrapTypeQueryParameter *string

	/* Stateful.

	   Filter by stateful
	*/
	StatefulQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems message collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsMessageCollectionGetParams) WithDefaults() *EmsMessageCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems message collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsMessageCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := EmsMessageCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithTimeout(timeout time.Duration) *EmsMessageCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithContext(ctx context.Context) *EmsMessageCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithHTTPClient(client *http.Client) *EmsMessageCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCorrectiveActionQueryParameter adds the correctiveAction to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithCorrectiveActionQueryParameter(correctiveAction *string) *EmsMessageCollectionGetParams {
	o.SetCorrectiveActionQueryParameter(correctiveAction)
	return o
}

// SetCorrectiveActionQueryParameter adds the correctiveAction to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetCorrectiveActionQueryParameter(correctiveAction *string) {
	o.CorrectiveActionQueryParameter = correctiveAction
}

// WithDeprecatedQueryParameter adds the deprecated to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithDeprecatedQueryParameter(deprecated *bool) *EmsMessageCollectionGetParams {
	o.SetDeprecatedQueryParameter(deprecated)
	return o
}

// SetDeprecatedQueryParameter adds the deprecated to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetDeprecatedQueryParameter(deprecated *bool) {
	o.DeprecatedQueryParameter = deprecated
}

// WithDescriptionQueryParameter adds the description to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithDescriptionQueryParameter(description *string) *EmsMessageCollectionGetParams {
	o.SetDescriptionQueryParameter(description)
	return o
}

// SetDescriptionQueryParameter adds the description to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetDescriptionQueryParameter(description *string) {
	o.DescriptionQueryParameter = description
}

// WithFieldsQueryParameter adds the fields to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithFieldsQueryParameter(fields []string) *EmsMessageCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *EmsMessageCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithNameQueryParameter(name *string) *EmsMessageCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *EmsMessageCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *EmsMessageCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *EmsMessageCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSeverityQueryParameter adds the severity to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithSeverityQueryParameter(severity *string) *EmsMessageCollectionGetParams {
	o.SetSeverityQueryParameter(severity)
	return o
}

// SetSeverityQueryParameter adds the severity to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetSeverityQueryParameter(severity *string) {
	o.SeverityQueryParameter = severity
}

// WithSnmpTrapTypeQueryParameter adds the snmpTrapType to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithSnmpTrapTypeQueryParameter(snmpTrapType *string) *EmsMessageCollectionGetParams {
	o.SetSnmpTrapTypeQueryParameter(snmpTrapType)
	return o
}

// SetSnmpTrapTypeQueryParameter adds the snmpTrapType to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetSnmpTrapTypeQueryParameter(snmpTrapType *string) {
	o.SnmpTrapTypeQueryParameter = snmpTrapType
}

// WithStatefulQueryParameter adds the stateful to the ems message collection get params
func (o *EmsMessageCollectionGetParams) WithStatefulQueryParameter(stateful *bool) *EmsMessageCollectionGetParams {
	o.SetStatefulQueryParameter(stateful)
	return o
}

// SetStatefulQueryParameter adds the stateful to the ems message collection get params
func (o *EmsMessageCollectionGetParams) SetStatefulQueryParameter(stateful *bool) {
	o.StatefulQueryParameter = stateful
}

// WriteToRequest writes these params to a swagger request
func (o *EmsMessageCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CorrectiveActionQueryParameter != nil {

		// query param corrective_action
		var qrCorrectiveAction string

		if o.CorrectiveActionQueryParameter != nil {
			qrCorrectiveAction = *o.CorrectiveActionQueryParameter
		}
		qCorrectiveAction := qrCorrectiveAction
		if qCorrectiveAction != "" {

			if err := r.SetQueryParam("corrective_action", qCorrectiveAction); err != nil {
				return err
			}
		}
	}

	if o.DeprecatedQueryParameter != nil {

		// query param deprecated
		var qrDeprecated bool

		if o.DeprecatedQueryParameter != nil {
			qrDeprecated = *o.DeprecatedQueryParameter
		}
		qDeprecated := swag.FormatBool(qrDeprecated)
		if qDeprecated != "" {

			if err := r.SetQueryParam("deprecated", qDeprecated); err != nil {
				return err
			}
		}
	}

	if o.DescriptionQueryParameter != nil {

		// query param description
		var qrDescription string

		if o.DescriptionQueryParameter != nil {
			qrDescription = *o.DescriptionQueryParameter
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SeverityQueryParameter != nil {

		// query param severity
		var qrSeverity string

		if o.SeverityQueryParameter != nil {
			qrSeverity = *o.SeverityQueryParameter
		}
		qSeverity := qrSeverity
		if qSeverity != "" {

			if err := r.SetQueryParam("severity", qSeverity); err != nil {
				return err
			}
		}
	}

	if o.SnmpTrapTypeQueryParameter != nil {

		// query param snmp_trap_type
		var qrSnmpTrapType string

		if o.SnmpTrapTypeQueryParameter != nil {
			qrSnmpTrapType = *o.SnmpTrapTypeQueryParameter
		}
		qSnmpTrapType := qrSnmpTrapType
		if qSnmpTrapType != "" {

			if err := r.SetQueryParam("snmp_trap_type", qSnmpTrapType); err != nil {
				return err
			}
		}
	}

	if o.StatefulQueryParameter != nil {

		// query param stateful
		var qrStateful bool

		if o.StatefulQueryParameter != nil {
			qrStateful = *o.StatefulQueryParameter
		}
		qStateful := swag.FormatBool(qrStateful)
		if qStateful != "" {

			if err := r.SetQueryParam("stateful", qStateful); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEmsMessageCollectionGet binds the parameter fields
func (o *EmsMessageCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamEmsMessageCollectionGet binds the parameter order_by
func (o *EmsMessageCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
