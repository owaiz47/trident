// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewConsistencyGroupSnapshotModifyParams creates a new ConsistencyGroupSnapshotModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConsistencyGroupSnapshotModifyParams() *ConsistencyGroupSnapshotModifyParams {
	return &ConsistencyGroupSnapshotModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConsistencyGroupSnapshotModifyParamsWithTimeout creates a new ConsistencyGroupSnapshotModifyParams object
// with the ability to set a timeout on a request.
func NewConsistencyGroupSnapshotModifyParamsWithTimeout(timeout time.Duration) *ConsistencyGroupSnapshotModifyParams {
	return &ConsistencyGroupSnapshotModifyParams{
		timeout: timeout,
	}
}

// NewConsistencyGroupSnapshotModifyParamsWithContext creates a new ConsistencyGroupSnapshotModifyParams object
// with the ability to set a context for a request.
func NewConsistencyGroupSnapshotModifyParamsWithContext(ctx context.Context) *ConsistencyGroupSnapshotModifyParams {
	return &ConsistencyGroupSnapshotModifyParams{
		Context: ctx,
	}
}

// NewConsistencyGroupSnapshotModifyParamsWithHTTPClient creates a new ConsistencyGroupSnapshotModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewConsistencyGroupSnapshotModifyParamsWithHTTPClient(client *http.Client) *ConsistencyGroupSnapshotModifyParams {
	return &ConsistencyGroupSnapshotModifyParams{
		HTTPClient: client,
	}
}

/*
ConsistencyGroupSnapshotModifyParams contains all the parameters to send to the API endpoint

	for the consistency group snapshot modify operation.

	Typically these are written to a http.Request.
*/
type ConsistencyGroupSnapshotModifyParams struct {

	/* Action.

	   Commits the Snapshot copy. The commit must be called within the timeout value specified during POST. If the commit is not called within the specified time, then the Snapshot copy create operation is aborted.

	*/
	Action *string

	/* ConsistencyGroupUUID.

	   Unique identifier of the consistency group's Snapshot copy for deletion.

	*/
	ConsistencyGroupUUID string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Snapshot copy UUID.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the consistency group snapshot modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsistencyGroupSnapshotModifyParams) WithDefaults() *ConsistencyGroupSnapshotModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the consistency group snapshot modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsistencyGroupSnapshotModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := ConsistencyGroupSnapshotModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) WithTimeout(timeout time.Duration) *ConsistencyGroupSnapshotModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) WithContext(ctx context.Context) *ConsistencyGroupSnapshotModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) WithHTTPClient(client *http.Client) *ConsistencyGroupSnapshotModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) WithAction(action *string) *ConsistencyGroupSnapshotModifyParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) SetAction(action *string) {
	o.Action = action
}

// WithConsistencyGroupUUID adds the consistencyGroupUUID to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) WithConsistencyGroupUUID(consistencyGroupUUID string) *ConsistencyGroupSnapshotModifyParams {
	o.SetConsistencyGroupUUID(consistencyGroupUUID)
	return o
}

// SetConsistencyGroupUUID adds the consistencyGroupUuid to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) SetConsistencyGroupUUID(consistencyGroupUUID string) {
	o.ConsistencyGroupUUID = consistencyGroupUUID
}

// WithReturnTimeout adds the returnTimeout to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) WithReturnTimeout(returnTimeout *int64) *ConsistencyGroupSnapshotModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) WithUUID(uuid string) *ConsistencyGroupSnapshotModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the consistency group snapshot modify params
func (o *ConsistencyGroupSnapshotModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ConsistencyGroupSnapshotModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	// path param consistency_group.uuid
	if err := r.SetPathParam("consistency_group.uuid", o.ConsistencyGroupUUID); err != nil {
		return err
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}