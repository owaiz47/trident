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

// NewMultiAdminVerifyRequestCollectionGetParams creates a new MultiAdminVerifyRequestCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyRequestCollectionGetParams() *MultiAdminVerifyRequestCollectionGetParams {
	return &MultiAdminVerifyRequestCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyRequestCollectionGetParamsWithTimeout creates a new MultiAdminVerifyRequestCollectionGetParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyRequestCollectionGetParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyRequestCollectionGetParams {
	return &MultiAdminVerifyRequestCollectionGetParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyRequestCollectionGetParamsWithContext creates a new MultiAdminVerifyRequestCollectionGetParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyRequestCollectionGetParamsWithContext(ctx context.Context) *MultiAdminVerifyRequestCollectionGetParams {
	return &MultiAdminVerifyRequestCollectionGetParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyRequestCollectionGetParamsWithHTTPClient creates a new MultiAdminVerifyRequestCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyRequestCollectionGetParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyRequestCollectionGetParams {
	return &MultiAdminVerifyRequestCollectionGetParams{
		HTTPClient: client,
	}
}

/* MultiAdminVerifyRequestCollectionGetParams contains all the parameters to send to the API endpoint
   for the multi admin verify request collection get operation.

   Typically these are written to a http.Request.
*/
type MultiAdminVerifyRequestCollectionGetParams struct {

	/* ApproveExpiryTime.

	   Filter by approve_expiry_time
	*/
	ApproveExpiryTimeQueryParameter *string

	/* ApproveTime.

	   Filter by approve_time
	*/
	ApproveTimeQueryParameter *string

	/* ApprovedUsers.

	   Filter by approved_users
	*/
	ApprovedUsersQueryParameter *string

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* CreateTime.

	   Filter by create_time
	*/
	CreateTimeQueryParameter *string

	/* ExecutionExpiryTime.

	   Filter by execution_expiry_time
	*/
	ExecutionExpiryTimeQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Operation.

	   Filter by operation
	*/
	OperationQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerNameQueryParameter *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUIDQueryParameter *string

	/* PendingApprovers.

	   Filter by pending_approvers
	*/
	PendingApproversQueryParameter *int64

	/* PermittedUsers.

	   Filter by permitted_users
	*/
	PermittedUsersQueryParameter *string

	/* PotentialApprovers.

	   Filter by potential_approvers
	*/
	PotentialApproversQueryParameter *string

	/* Query.

	   Filter by query
	*/
	QueryQueryParameter *string

	/* RequiredApprovers.

	   Filter by required_approvers
	*/
	RequiredApproversQueryParameter *int64

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

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	/* UserRequested.

	   Filter by user_requested
	*/
	UserRequestedQueryParameter *string

	/* UserVetoed.

	   Filter by user_vetoed
	*/
	UserVetoedQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify request collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestCollectionGetParams) WithDefaults() *MultiAdminVerifyRequestCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify request collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := MultiAdminVerifyRequestCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithContext(ctx context.Context) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApproveExpiryTimeQueryParameter adds the approveExpiryTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithApproveExpiryTimeQueryParameter(approveExpiryTime *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetApproveExpiryTimeQueryParameter(approveExpiryTime)
	return o
}

// SetApproveExpiryTimeQueryParameter adds the approveExpiryTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetApproveExpiryTimeQueryParameter(approveExpiryTime *string) {
	o.ApproveExpiryTimeQueryParameter = approveExpiryTime
}

// WithApproveTimeQueryParameter adds the approveTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithApproveTimeQueryParameter(approveTime *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetApproveTimeQueryParameter(approveTime)
	return o
}

// SetApproveTimeQueryParameter adds the approveTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetApproveTimeQueryParameter(approveTime *string) {
	o.ApproveTimeQueryParameter = approveTime
}

// WithApprovedUsersQueryParameter adds the approvedUsers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithApprovedUsersQueryParameter(approvedUsers *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetApprovedUsersQueryParameter(approvedUsers)
	return o
}

// SetApprovedUsersQueryParameter adds the approvedUsers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetApprovedUsersQueryParameter(approvedUsers *string) {
	o.ApprovedUsersQueryParameter = approvedUsers
}

// WithCommentQueryParameter adds the comment to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithCommentQueryParameter(comment *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithCreateTimeQueryParameter adds the createTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithCreateTimeQueryParameter(createTime *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetCreateTimeQueryParameter(createTime)
	return o
}

// SetCreateTimeQueryParameter adds the createTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetCreateTimeQueryParameter(createTime *string) {
	o.CreateTimeQueryParameter = createTime
}

// WithExecutionExpiryTimeQueryParameter adds the executionExpiryTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithExecutionExpiryTimeQueryParameter(executionExpiryTime *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetExecutionExpiryTimeQueryParameter(executionExpiryTime)
	return o
}

// SetExecutionExpiryTimeQueryParameter adds the executionExpiryTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetExecutionExpiryTimeQueryParameter(executionExpiryTime *string) {
	o.ExecutionExpiryTimeQueryParameter = executionExpiryTime
}

// WithFieldsQueryParameter adds the fields to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithFieldsQueryParameter(fields []string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexQueryParameter adds the index to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithIndexQueryParameter(index *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithMaxRecordsQueryParameter adds the maxRecords to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOperationQueryParameter adds the operation to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithOperationQueryParameter(operation *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetOperationQueryParameter(operation)
	return o
}

// SetOperationQueryParameter adds the operation to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetOperationQueryParameter(operation *string) {
	o.OperationQueryParameter = operation
}

// WithOrderByQueryParameter adds the orderBy to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithOwnerNameQueryParameter adds the ownerName to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithOwnerNameQueryParameter(ownerName *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetOwnerNameQueryParameter(ownerName)
	return o
}

// SetOwnerNameQueryParameter adds the ownerName to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetOwnerNameQueryParameter(ownerName *string) {
	o.OwnerNameQueryParameter = ownerName
}

// WithOwnerUUIDQueryParameter adds the ownerUUID to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithOwnerUUIDQueryParameter(ownerUUID *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetOwnerUUIDQueryParameter(ownerUUID)
	return o
}

// SetOwnerUUIDQueryParameter adds the ownerUuid to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetOwnerUUIDQueryParameter(ownerUUID *string) {
	o.OwnerUUIDQueryParameter = ownerUUID
}

// WithPendingApproversQueryParameter adds the pendingApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithPendingApproversQueryParameter(pendingApprovers *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetPendingApproversQueryParameter(pendingApprovers)
	return o
}

// SetPendingApproversQueryParameter adds the pendingApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetPendingApproversQueryParameter(pendingApprovers *int64) {
	o.PendingApproversQueryParameter = pendingApprovers
}

// WithPermittedUsersQueryParameter adds the permittedUsers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithPermittedUsersQueryParameter(permittedUsers *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetPermittedUsersQueryParameter(permittedUsers)
	return o
}

// SetPermittedUsersQueryParameter adds the permittedUsers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetPermittedUsersQueryParameter(permittedUsers *string) {
	o.PermittedUsersQueryParameter = permittedUsers
}

// WithPotentialApproversQueryParameter adds the potentialApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithPotentialApproversQueryParameter(potentialApprovers *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetPotentialApproversQueryParameter(potentialApprovers)
	return o
}

// SetPotentialApproversQueryParameter adds the potentialApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetPotentialApproversQueryParameter(potentialApprovers *string) {
	o.PotentialApproversQueryParameter = potentialApprovers
}

// WithQueryQueryParameter adds the query to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithQueryQueryParameter(query *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetQueryQueryParameter(query)
	return o
}

// SetQueryQueryParameter adds the query to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetQueryQueryParameter(query *string) {
	o.QueryQueryParameter = query
}

// WithRequiredApproversQueryParameter adds the requiredApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithRequiredApproversQueryParameter(requiredApprovers *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetRequiredApproversQueryParameter(requiredApprovers)
	return o
}

// SetRequiredApproversQueryParameter adds the requiredApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetRequiredApproversQueryParameter(requiredApprovers *int64) {
	o.RequiredApproversQueryParameter = requiredApprovers
}

// WithReturnRecordsQueryParameter adds the returnRecords to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithStateQueryParameter adds the state to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithStateQueryParameter(state *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WithUserRequestedQueryParameter adds the userRequested to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithUserRequestedQueryParameter(userRequested *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetUserRequestedQueryParameter(userRequested)
	return o
}

// SetUserRequestedQueryParameter adds the userRequested to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetUserRequestedQueryParameter(userRequested *string) {
	o.UserRequestedQueryParameter = userRequested
}

// WithUserVetoedQueryParameter adds the userVetoed to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithUserVetoedQueryParameter(userVetoed *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetUserVetoedQueryParameter(userVetoed)
	return o
}

// SetUserVetoedQueryParameter adds the userVetoed to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetUserVetoedQueryParameter(userVetoed *string) {
	o.UserVetoedQueryParameter = userVetoed
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyRequestCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApproveExpiryTimeQueryParameter != nil {

		// query param approve_expiry_time
		var qrApproveExpiryTime string

		if o.ApproveExpiryTimeQueryParameter != nil {
			qrApproveExpiryTime = *o.ApproveExpiryTimeQueryParameter
		}
		qApproveExpiryTime := qrApproveExpiryTime
		if qApproveExpiryTime != "" {

			if err := r.SetQueryParam("approve_expiry_time", qApproveExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.ApproveTimeQueryParameter != nil {

		// query param approve_time
		var qrApproveTime string

		if o.ApproveTimeQueryParameter != nil {
			qrApproveTime = *o.ApproveTimeQueryParameter
		}
		qApproveTime := qrApproveTime
		if qApproveTime != "" {

			if err := r.SetQueryParam("approve_time", qApproveTime); err != nil {
				return err
			}
		}
	}

	if o.ApprovedUsersQueryParameter != nil {

		// query param approved_users
		var qrApprovedUsers string

		if o.ApprovedUsersQueryParameter != nil {
			qrApprovedUsers = *o.ApprovedUsersQueryParameter
		}
		qApprovedUsers := qrApprovedUsers
		if qApprovedUsers != "" {

			if err := r.SetQueryParam("approved_users", qApprovedUsers); err != nil {
				return err
			}
		}
	}

	if o.CommentQueryParameter != nil {

		// query param comment
		var qrComment string

		if o.CommentQueryParameter != nil {
			qrComment = *o.CommentQueryParameter
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.CreateTimeQueryParameter != nil {

		// query param create_time
		var qrCreateTime string

		if o.CreateTimeQueryParameter != nil {
			qrCreateTime = *o.CreateTimeQueryParameter
		}
		qCreateTime := qrCreateTime
		if qCreateTime != "" {

			if err := r.SetQueryParam("create_time", qCreateTime); err != nil {
				return err
			}
		}
	}

	if o.ExecutionExpiryTimeQueryParameter != nil {

		// query param execution_expiry_time
		var qrExecutionExpiryTime string

		if o.ExecutionExpiryTimeQueryParameter != nil {
			qrExecutionExpiryTime = *o.ExecutionExpiryTimeQueryParameter
		}
		qExecutionExpiryTime := qrExecutionExpiryTime
		if qExecutionExpiryTime != "" {

			if err := r.SetQueryParam("execution_expiry_time", qExecutionExpiryTime); err != nil {
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

	if o.IndexQueryParameter != nil {

		// query param index
		var qrIndex int64

		if o.IndexQueryParameter != nil {
			qrIndex = *o.IndexQueryParameter
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
				return err
			}
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

	if o.OperationQueryParameter != nil {

		// query param operation
		var qrOperation string

		if o.OperationQueryParameter != nil {
			qrOperation = *o.OperationQueryParameter
		}
		qOperation := qrOperation
		if qOperation != "" {

			if err := r.SetQueryParam("operation", qOperation); err != nil {
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

	if o.OwnerNameQueryParameter != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerNameQueryParameter != nil {
			qrOwnerName = *o.OwnerNameQueryParameter
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUIDQueryParameter != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUIDQueryParameter != nil {
			qrOwnerUUID = *o.OwnerUUIDQueryParameter
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
				return err
			}
		}
	}

	if o.PendingApproversQueryParameter != nil {

		// query param pending_approvers
		var qrPendingApprovers int64

		if o.PendingApproversQueryParameter != nil {
			qrPendingApprovers = *o.PendingApproversQueryParameter
		}
		qPendingApprovers := swag.FormatInt64(qrPendingApprovers)
		if qPendingApprovers != "" {

			if err := r.SetQueryParam("pending_approvers", qPendingApprovers); err != nil {
				return err
			}
		}
	}

	if o.PermittedUsersQueryParameter != nil {

		// query param permitted_users
		var qrPermittedUsers string

		if o.PermittedUsersQueryParameter != nil {
			qrPermittedUsers = *o.PermittedUsersQueryParameter
		}
		qPermittedUsers := qrPermittedUsers
		if qPermittedUsers != "" {

			if err := r.SetQueryParam("permitted_users", qPermittedUsers); err != nil {
				return err
			}
		}
	}

	if o.PotentialApproversQueryParameter != nil {

		// query param potential_approvers
		var qrPotentialApprovers string

		if o.PotentialApproversQueryParameter != nil {
			qrPotentialApprovers = *o.PotentialApproversQueryParameter
		}
		qPotentialApprovers := qrPotentialApprovers
		if qPotentialApprovers != "" {

			if err := r.SetQueryParam("potential_approvers", qPotentialApprovers); err != nil {
				return err
			}
		}
	}

	if o.QueryQueryParameter != nil {

		// query param query
		var qrQuery string

		if o.QueryQueryParameter != nil {
			qrQuery = *o.QueryQueryParameter
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.RequiredApproversQueryParameter != nil {

		// query param required_approvers
		var qrRequiredApprovers int64

		if o.RequiredApproversQueryParameter != nil {
			qrRequiredApprovers = *o.RequiredApproversQueryParameter
		}
		qRequiredApprovers := swag.FormatInt64(qrRequiredApprovers)
		if qRequiredApprovers != "" {

			if err := r.SetQueryParam("required_approvers", qRequiredApprovers); err != nil {
				return err
			}
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

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.UserRequestedQueryParameter != nil {

		// query param user_requested
		var qrUserRequested string

		if o.UserRequestedQueryParameter != nil {
			qrUserRequested = *o.UserRequestedQueryParameter
		}
		qUserRequested := qrUserRequested
		if qUserRequested != "" {

			if err := r.SetQueryParam("user_requested", qUserRequested); err != nil {
				return err
			}
		}
	}

	if o.UserVetoedQueryParameter != nil {

		// query param user_vetoed
		var qrUserVetoed string

		if o.UserVetoedQueryParameter != nil {
			qrUserVetoed = *o.UserVetoedQueryParameter
		}
		qUserVetoed := qrUserVetoed
		if qUserVetoed != "" {

			if err := r.SetQueryParam("user_vetoed", qUserVetoed); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamMultiAdminVerifyRequestCollectionGet binds the parameter fields
func (o *MultiAdminVerifyRequestCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamMultiAdminVerifyRequestCollectionGet binds the parameter order_by
func (o *MultiAdminVerifyRequestCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
