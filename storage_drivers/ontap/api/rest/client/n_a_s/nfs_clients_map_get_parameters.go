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

// NewNfsClientsMapGetParams creates a new NfsClientsMapGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNfsClientsMapGetParams() *NfsClientsMapGetParams {
	return &NfsClientsMapGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNfsClientsMapGetParamsWithTimeout creates a new NfsClientsMapGetParams object
// with the ability to set a timeout on a request.
func NewNfsClientsMapGetParamsWithTimeout(timeout time.Duration) *NfsClientsMapGetParams {
	return &NfsClientsMapGetParams{
		timeout: timeout,
	}
}

// NewNfsClientsMapGetParamsWithContext creates a new NfsClientsMapGetParams object
// with the ability to set a context for a request.
func NewNfsClientsMapGetParamsWithContext(ctx context.Context) *NfsClientsMapGetParams {
	return &NfsClientsMapGetParams{
		Context: ctx,
	}
}

// NewNfsClientsMapGetParamsWithHTTPClient creates a new NfsClientsMapGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNfsClientsMapGetParamsWithHTTPClient(client *http.Client) *NfsClientsMapGetParams {
	return &NfsClientsMapGetParams{
		HTTPClient: client,
	}
}

/* NfsClientsMapGetParams contains all the parameters to send to the API endpoint
   for the nfs clients map get operation.

   Typically these are written to a http.Request.
*/
type NfsClientsMapGetParams struct {

	/* ClientIps.

	   Filter by client_ips
	*/
	ClientIPsQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

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

	/* ServerIP.

	   Filter by server_ip
	*/
	ServerIPQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nfs clients map get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsClientsMapGetParams) WithDefaults() *NfsClientsMapGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nfs clients map get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsClientsMapGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NfsClientsMapGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithTimeout(timeout time.Duration) *NfsClientsMapGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithContext(ctx context.Context) *NfsClientsMapGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithHTTPClient(client *http.Client) *NfsClientsMapGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientIPsQueryParameter adds the clientIps to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithClientIPsQueryParameter(clientIps *string) *NfsClientsMapGetParams {
	o.SetClientIPsQueryParameter(clientIps)
	return o
}

// SetClientIPsQueryParameter adds the clientIps to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetClientIPsQueryParameter(clientIps *string) {
	o.ClientIPsQueryParameter = clientIps
}

// WithFieldsQueryParameter adds the fields to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithFieldsQueryParameter(fields []string) *NfsClientsMapGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NfsClientsMapGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNodeNameQueryParameter adds the nodeName to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithNodeNameQueryParameter(nodeName *string) *NfsClientsMapGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *NfsClientsMapGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderByQueryParameter adds the orderBy to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithOrderByQueryParameter(orderBy []string) *NfsClientsMapGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NfsClientsMapGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NfsClientsMapGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithServerIPQueryParameter adds the serverIP to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithServerIPQueryParameter(serverIP *string) *NfsClientsMapGetParams {
	o.SetServerIPQueryParameter(serverIP)
	return o
}

// SetServerIPQueryParameter adds the serverIp to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetServerIPQueryParameter(serverIP *string) {
	o.ServerIPQueryParameter = serverIP
}

// WithSVMNameQueryParameter adds the svmName to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithSVMNameQueryParameter(svmName *string) *NfsClientsMapGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the nfs clients map get params
func (o *NfsClientsMapGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NfsClientsMapGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the nfs clients map get params
func (o *NfsClientsMapGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NfsClientsMapGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientIPsQueryParameter != nil {

		// query param client_ips
		var qrClientIps string

		if o.ClientIPsQueryParameter != nil {
			qrClientIps = *o.ClientIPsQueryParameter
		}
		qClientIps := qrClientIps
		if qClientIps != "" {

			if err := r.SetQueryParam("client_ips", qClientIps); err != nil {
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

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.ServerIPQueryParameter != nil {

		// query param server_ip
		var qrServerIP string

		if o.ServerIPQueryParameter != nil {
			qrServerIP = *o.ServerIPQueryParameter
		}
		qServerIP := qrServerIP
		if qServerIP != "" {

			if err := r.SetQueryParam("server_ip", qServerIP); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNfsClientsMapGet binds the parameter fields
func (o *NfsClientsMapGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNfsClientsMapGet binds the parameter order_by
func (o *NfsClientsMapGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
