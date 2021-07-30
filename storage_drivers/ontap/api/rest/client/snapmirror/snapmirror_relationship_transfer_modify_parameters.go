// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSnapmirrorRelationshipTransferModifyParams creates a new SnapmirrorRelationshipTransferModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapmirrorRelationshipTransferModifyParams() *SnapmirrorRelationshipTransferModifyParams {
	return &SnapmirrorRelationshipTransferModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapmirrorRelationshipTransferModifyParamsWithTimeout creates a new SnapmirrorRelationshipTransferModifyParams object
// with the ability to set a timeout on a request.
func NewSnapmirrorRelationshipTransferModifyParamsWithTimeout(timeout time.Duration) *SnapmirrorRelationshipTransferModifyParams {
	return &SnapmirrorRelationshipTransferModifyParams{
		timeout: timeout,
	}
}

// NewSnapmirrorRelationshipTransferModifyParamsWithContext creates a new SnapmirrorRelationshipTransferModifyParams object
// with the ability to set a context for a request.
func NewSnapmirrorRelationshipTransferModifyParamsWithContext(ctx context.Context) *SnapmirrorRelationshipTransferModifyParams {
	return &SnapmirrorRelationshipTransferModifyParams{
		Context: ctx,
	}
}

// NewSnapmirrorRelationshipTransferModifyParamsWithHTTPClient creates a new SnapmirrorRelationshipTransferModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapmirrorRelationshipTransferModifyParamsWithHTTPClient(client *http.Client) *SnapmirrorRelationshipTransferModifyParams {
	return &SnapmirrorRelationshipTransferModifyParams{
		HTTPClient: client,
	}
}

/* SnapmirrorRelationshipTransferModifyParams contains all the parameters to send to the API endpoint
   for the snapmirror relationship transfer modify operation.

   Typically these are written to a http.Request.
*/
type SnapmirrorRelationshipTransferModifyParams struct {

	/* Info.

	   Information on the SnapMirror transfer
	*/
	Info *models.SnapmirrorTransfer

	/* RelationshipUUID.

	   Relationship UUID
	*/
	RelationshIPUUIDPathParameter string

	/* UUID.

	   Transfer UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapmirror relationship transfer modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorRelationshipTransferModifyParams) WithDefaults() *SnapmirrorRelationshipTransferModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapmirror relationship transfer modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorRelationshipTransferModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) WithTimeout(timeout time.Duration) *SnapmirrorRelationshipTransferModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) WithContext(ctx context.Context) *SnapmirrorRelationshipTransferModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) WithHTTPClient(client *http.Client) *SnapmirrorRelationshipTransferModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) WithInfo(info *models.SnapmirrorTransfer) *SnapmirrorRelationshipTransferModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) SetInfo(info *models.SnapmirrorTransfer) {
	o.Info = info
}

// WithRelationshIPUUIDPathParameter adds the relationshipUUID to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) WithRelationshIPUUIDPathParameter(relationshipUUID string) *SnapmirrorRelationshipTransferModifyParams {
	o.SetRelationshIPUUIDPathParameter(relationshipUUID)
	return o
}

// SetRelationshIPUUIDPathParameter adds the relationshipUuid to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) SetRelationshIPUUIDPathParameter(relationshipUUID string) {
	o.RelationshIPUUIDPathParameter = relationshipUUID
}

// WithUUIDPathParameter adds the uuid to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) WithUUIDPathParameter(uuid string) *SnapmirrorRelationshipTransferModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the snapmirror relationship transfer modify params
func (o *SnapmirrorRelationshipTransferModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapmirrorRelationshipTransferModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param relationship.uuid
	if err := r.SetPathParam("relationship.uuid", o.RelationshIPUUIDPathParameter); err != nil {
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