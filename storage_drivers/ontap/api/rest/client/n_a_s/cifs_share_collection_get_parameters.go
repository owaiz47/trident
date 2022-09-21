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

// NewCifsShareCollectionGetParams creates a new CifsShareCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsShareCollectionGetParams() *CifsShareCollectionGetParams {
	return &CifsShareCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsShareCollectionGetParamsWithTimeout creates a new CifsShareCollectionGetParams object
// with the ability to set a timeout on a request.
func NewCifsShareCollectionGetParamsWithTimeout(timeout time.Duration) *CifsShareCollectionGetParams {
	return &CifsShareCollectionGetParams{
		timeout: timeout,
	}
}

// NewCifsShareCollectionGetParamsWithContext creates a new CifsShareCollectionGetParams object
// with the ability to set a context for a request.
func NewCifsShareCollectionGetParamsWithContext(ctx context.Context) *CifsShareCollectionGetParams {
	return &CifsShareCollectionGetParams{
		Context: ctx,
	}
}

// NewCifsShareCollectionGetParamsWithHTTPClient creates a new CifsShareCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsShareCollectionGetParamsWithHTTPClient(client *http.Client) *CifsShareCollectionGetParams {
	return &CifsShareCollectionGetParams{
		HTTPClient: client,
	}
}

/* CifsShareCollectionGetParams contains all the parameters to send to the API endpoint
   for the cifs share collection get operation.

   Typically these are written to a http.Request.
*/
type CifsShareCollectionGetParams struct {

	/* AccessBasedEnumeration.

	   Filter by access_based_enumeration
	*/
	AccessBasedEnumerationQueryParameter *bool

	/* AclsPermission.

	   Filter by acls.permission
	*/
	AclsPermissionQueryParameter *string

	/* AclsType.

	   Filter by acls.type
	*/
	AclsTypeQueryParameter *string

	/* AclsUserOrGroup.

	   Filter by acls.user_or_group
	*/
	AclsUserOrGroupQueryParameter *string

	/* AllowUnencryptedAccess.

	   Filter by allow_unencrypted_access
	*/
	AllowUnencryptedAccessQueryParameter *bool

	/* ChangeNotify.

	   Filter by change_notify
	*/
	ChangeNotifyQueryParameter *bool

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* ContinuouslyAvailable.

	   Filter by continuously_available
	*/
	ContinuouslyAvailableQueryParameter *bool

	/* DirUmask.

	   Filter by dir_umask
	*/
	DirUmaskQueryParameter *int64

	/* Encryption.

	   Filter by encryption
	*/
	EncryptionQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* FileUmask.

	   Filter by file_umask
	*/
	FileUmaskQueryParameter *int64

	/* ForceGroupForCreate.

	   Filter by force_group_for_create
	*/
	ForceGroupForCreateQueryParameter *string

	/* HomeDirectory.

	   Filter by home_directory
	*/
	HomeDirectoryQueryParameter *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* NamespaceCaching.

	   Filter by namespace_caching
	*/
	NamespaceCachingQueryParameter *bool

	/* NoStrictSecurity.

	   Filter by no_strict_security
	*/
	NoStrictSecurityQueryParameter *bool

	/* OfflineFiles.

	   Filter by offline_files
	*/
	OfflineFilesQueryParameter *string

	/* Oplocks.

	   Filter by oplocks
	*/
	OplocksQueryParameter *bool

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* Path.

	   Filter by path
	*/
	PathQueryParameter *string

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

	/* ShowSnapshot.

	   Filter by show_snapshot
	*/
	ShowSnapshotQueryParameter *bool

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* UnixSymlink.

	   Filter by unix_symlink
	*/
	UnixSymlinkQueryParameter *string

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeNameQueryParameter *string

	/* VolumeUUID.

	   Filter by volume.uuid
	*/
	VolumeUUIDQueryParameter *string

	/* VscanProfile.

	   Filter by vscan_profile
	*/
	VscanProfileQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs share collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareCollectionGetParams) WithDefaults() *CifsShareCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs share collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := CifsShareCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithTimeout(timeout time.Duration) *CifsShareCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithContext(ctx context.Context) *CifsShareCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithHTTPClient(client *http.Client) *CifsShareCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessBasedEnumerationQueryParameter adds the accessBasedEnumeration to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithAccessBasedEnumerationQueryParameter(accessBasedEnumeration *bool) *CifsShareCollectionGetParams {
	o.SetAccessBasedEnumerationQueryParameter(accessBasedEnumeration)
	return o
}

// SetAccessBasedEnumerationQueryParameter adds the accessBasedEnumeration to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetAccessBasedEnumerationQueryParameter(accessBasedEnumeration *bool) {
	o.AccessBasedEnumerationQueryParameter = accessBasedEnumeration
}

// WithAclsPermissionQueryParameter adds the aclsPermission to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithAclsPermissionQueryParameter(aclsPermission *string) *CifsShareCollectionGetParams {
	o.SetAclsPermissionQueryParameter(aclsPermission)
	return o
}

// SetAclsPermissionQueryParameter adds the aclsPermission to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetAclsPermissionQueryParameter(aclsPermission *string) {
	o.AclsPermissionQueryParameter = aclsPermission
}

// WithAclsTypeQueryParameter adds the aclsType to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithAclsTypeQueryParameter(aclsType *string) *CifsShareCollectionGetParams {
	o.SetAclsTypeQueryParameter(aclsType)
	return o
}

// SetAclsTypeQueryParameter adds the aclsType to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetAclsTypeQueryParameter(aclsType *string) {
	o.AclsTypeQueryParameter = aclsType
}

// WithAclsUserOrGroupQueryParameter adds the aclsUserOrGroup to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithAclsUserOrGroupQueryParameter(aclsUserOrGroup *string) *CifsShareCollectionGetParams {
	o.SetAclsUserOrGroupQueryParameter(aclsUserOrGroup)
	return o
}

// SetAclsUserOrGroupQueryParameter adds the aclsUserOrGroup to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetAclsUserOrGroupQueryParameter(aclsUserOrGroup *string) {
	o.AclsUserOrGroupQueryParameter = aclsUserOrGroup
}

// WithAllowUnencryptedAccessQueryParameter adds the allowUnencryptedAccess to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithAllowUnencryptedAccessQueryParameter(allowUnencryptedAccess *bool) *CifsShareCollectionGetParams {
	o.SetAllowUnencryptedAccessQueryParameter(allowUnencryptedAccess)
	return o
}

// SetAllowUnencryptedAccessQueryParameter adds the allowUnencryptedAccess to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetAllowUnencryptedAccessQueryParameter(allowUnencryptedAccess *bool) {
	o.AllowUnencryptedAccessQueryParameter = allowUnencryptedAccess
}

// WithChangeNotifyQueryParameter adds the changeNotify to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithChangeNotifyQueryParameter(changeNotify *bool) *CifsShareCollectionGetParams {
	o.SetChangeNotifyQueryParameter(changeNotify)
	return o
}

// SetChangeNotifyQueryParameter adds the changeNotify to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetChangeNotifyQueryParameter(changeNotify *bool) {
	o.ChangeNotifyQueryParameter = changeNotify
}

// WithCommentQueryParameter adds the comment to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithCommentQueryParameter(comment *string) *CifsShareCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithContinuouslyAvailableQueryParameter adds the continuouslyAvailable to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithContinuouslyAvailableQueryParameter(continuouslyAvailable *bool) *CifsShareCollectionGetParams {
	o.SetContinuouslyAvailableQueryParameter(continuouslyAvailable)
	return o
}

// SetContinuouslyAvailableQueryParameter adds the continuouslyAvailable to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetContinuouslyAvailableQueryParameter(continuouslyAvailable *bool) {
	o.ContinuouslyAvailableQueryParameter = continuouslyAvailable
}

// WithDirUmaskQueryParameter adds the dirUmask to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithDirUmaskQueryParameter(dirUmask *int64) *CifsShareCollectionGetParams {
	o.SetDirUmaskQueryParameter(dirUmask)
	return o
}

// SetDirUmaskQueryParameter adds the dirUmask to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetDirUmaskQueryParameter(dirUmask *int64) {
	o.DirUmaskQueryParameter = dirUmask
}

// WithEncryptionQueryParameter adds the encryption to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithEncryptionQueryParameter(encryption *bool) *CifsShareCollectionGetParams {
	o.SetEncryptionQueryParameter(encryption)
	return o
}

// SetEncryptionQueryParameter adds the encryption to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetEncryptionQueryParameter(encryption *bool) {
	o.EncryptionQueryParameter = encryption
}

// WithFieldsQueryParameter adds the fields to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithFieldsQueryParameter(fields []string) *CifsShareCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithFileUmaskQueryParameter adds the fileUmask to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithFileUmaskQueryParameter(fileUmask *int64) *CifsShareCollectionGetParams {
	o.SetFileUmaskQueryParameter(fileUmask)
	return o
}

// SetFileUmaskQueryParameter adds the fileUmask to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetFileUmaskQueryParameter(fileUmask *int64) {
	o.FileUmaskQueryParameter = fileUmask
}

// WithForceGroupForCreateQueryParameter adds the forceGroupForCreate to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithForceGroupForCreateQueryParameter(forceGroupForCreate *string) *CifsShareCollectionGetParams {
	o.SetForceGroupForCreateQueryParameter(forceGroupForCreate)
	return o
}

// SetForceGroupForCreateQueryParameter adds the forceGroupForCreate to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetForceGroupForCreateQueryParameter(forceGroupForCreate *string) {
	o.ForceGroupForCreateQueryParameter = forceGroupForCreate
}

// WithHomeDirectoryQueryParameter adds the homeDirectory to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithHomeDirectoryQueryParameter(homeDirectory *bool) *CifsShareCollectionGetParams {
	o.SetHomeDirectoryQueryParameter(homeDirectory)
	return o
}

// SetHomeDirectoryQueryParameter adds the homeDirectory to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetHomeDirectoryQueryParameter(homeDirectory *bool) {
	o.HomeDirectoryQueryParameter = homeDirectory
}

// WithMaxRecordsQueryParameter adds the maxRecords to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *CifsShareCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithNameQueryParameter(name *string) *CifsShareCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithNamespaceCachingQueryParameter adds the namespaceCaching to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithNamespaceCachingQueryParameter(namespaceCaching *bool) *CifsShareCollectionGetParams {
	o.SetNamespaceCachingQueryParameter(namespaceCaching)
	return o
}

// SetNamespaceCachingQueryParameter adds the namespaceCaching to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetNamespaceCachingQueryParameter(namespaceCaching *bool) {
	o.NamespaceCachingQueryParameter = namespaceCaching
}

// WithNoStrictSecurityQueryParameter adds the noStrictSecurity to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithNoStrictSecurityQueryParameter(noStrictSecurity *bool) *CifsShareCollectionGetParams {
	o.SetNoStrictSecurityQueryParameter(noStrictSecurity)
	return o
}

// SetNoStrictSecurityQueryParameter adds the noStrictSecurity to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetNoStrictSecurityQueryParameter(noStrictSecurity *bool) {
	o.NoStrictSecurityQueryParameter = noStrictSecurity
}

// WithOfflineFilesQueryParameter adds the offlineFiles to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithOfflineFilesQueryParameter(offlineFiles *string) *CifsShareCollectionGetParams {
	o.SetOfflineFilesQueryParameter(offlineFiles)
	return o
}

// SetOfflineFilesQueryParameter adds the offlineFiles to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetOfflineFilesQueryParameter(offlineFiles *string) {
	o.OfflineFilesQueryParameter = offlineFiles
}

// WithOplocksQueryParameter adds the oplocks to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithOplocksQueryParameter(oplocks *bool) *CifsShareCollectionGetParams {
	o.SetOplocksQueryParameter(oplocks)
	return o
}

// SetOplocksQueryParameter adds the oplocks to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetOplocksQueryParameter(oplocks *bool) {
	o.OplocksQueryParameter = oplocks
}

// WithOrderByQueryParameter adds the orderBy to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *CifsShareCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPathQueryParameter adds the path to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithPathQueryParameter(path *string) *CifsShareCollectionGetParams {
	o.SetPathQueryParameter(path)
	return o
}

// SetPathQueryParameter adds the path to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetPathQueryParameter(path *string) {
	o.PathQueryParameter = path
}

// WithReturnRecordsQueryParameter adds the returnRecords to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *CifsShareCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *CifsShareCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithShowSnapshotQueryParameter adds the showSnapshot to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithShowSnapshotQueryParameter(showSnapshot *bool) *CifsShareCollectionGetParams {
	o.SetShowSnapshotQueryParameter(showSnapshot)
	return o
}

// SetShowSnapshotQueryParameter adds the showSnapshot to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetShowSnapshotQueryParameter(showSnapshot *bool) {
	o.ShowSnapshotQueryParameter = showSnapshot
}

// WithSVMNameQueryParameter adds the svmName to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *CifsShareCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *CifsShareCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUnixSymlinkQueryParameter adds the unixSymlink to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithUnixSymlinkQueryParameter(unixSymlink *string) *CifsShareCollectionGetParams {
	o.SetUnixSymlinkQueryParameter(unixSymlink)
	return o
}

// SetUnixSymlinkQueryParameter adds the unixSymlink to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetUnixSymlinkQueryParameter(unixSymlink *string) {
	o.UnixSymlinkQueryParameter = unixSymlink
}

// WithVolumeNameQueryParameter adds the volumeName to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithVolumeNameQueryParameter(volumeName *string) *CifsShareCollectionGetParams {
	o.SetVolumeNameQueryParameter(volumeName)
	return o
}

// SetVolumeNameQueryParameter adds the volumeName to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetVolumeNameQueryParameter(volumeName *string) {
	o.VolumeNameQueryParameter = volumeName
}

// WithVolumeUUIDQueryParameter adds the volumeUUID to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithVolumeUUIDQueryParameter(volumeUUID *string) *CifsShareCollectionGetParams {
	o.SetVolumeUUIDQueryParameter(volumeUUID)
	return o
}

// SetVolumeUUIDQueryParameter adds the volumeUuid to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetVolumeUUIDQueryParameter(volumeUUID *string) {
	o.VolumeUUIDQueryParameter = volumeUUID
}

// WithVscanProfileQueryParameter adds the vscanProfile to the cifs share collection get params
func (o *CifsShareCollectionGetParams) WithVscanProfileQueryParameter(vscanProfile *string) *CifsShareCollectionGetParams {
	o.SetVscanProfileQueryParameter(vscanProfile)
	return o
}

// SetVscanProfileQueryParameter adds the vscanProfile to the cifs share collection get params
func (o *CifsShareCollectionGetParams) SetVscanProfileQueryParameter(vscanProfile *string) {
	o.VscanProfileQueryParameter = vscanProfile
}

// WriteToRequest writes these params to a swagger request
func (o *CifsShareCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessBasedEnumerationQueryParameter != nil {

		// query param access_based_enumeration
		var qrAccessBasedEnumeration bool

		if o.AccessBasedEnumerationQueryParameter != nil {
			qrAccessBasedEnumeration = *o.AccessBasedEnumerationQueryParameter
		}
		qAccessBasedEnumeration := swag.FormatBool(qrAccessBasedEnumeration)
		if qAccessBasedEnumeration != "" {

			if err := r.SetQueryParam("access_based_enumeration", qAccessBasedEnumeration); err != nil {
				return err
			}
		}
	}

	if o.AclsPermissionQueryParameter != nil {

		// query param acls.permission
		var qrAclsPermission string

		if o.AclsPermissionQueryParameter != nil {
			qrAclsPermission = *o.AclsPermissionQueryParameter
		}
		qAclsPermission := qrAclsPermission
		if qAclsPermission != "" {

			if err := r.SetQueryParam("acls.permission", qAclsPermission); err != nil {
				return err
			}
		}
	}

	if o.AclsTypeQueryParameter != nil {

		// query param acls.type
		var qrAclsType string

		if o.AclsTypeQueryParameter != nil {
			qrAclsType = *o.AclsTypeQueryParameter
		}
		qAclsType := qrAclsType
		if qAclsType != "" {

			if err := r.SetQueryParam("acls.type", qAclsType); err != nil {
				return err
			}
		}
	}

	if o.AclsUserOrGroupQueryParameter != nil {

		// query param acls.user_or_group
		var qrAclsUserOrGroup string

		if o.AclsUserOrGroupQueryParameter != nil {
			qrAclsUserOrGroup = *o.AclsUserOrGroupQueryParameter
		}
		qAclsUserOrGroup := qrAclsUserOrGroup
		if qAclsUserOrGroup != "" {

			if err := r.SetQueryParam("acls.user_or_group", qAclsUserOrGroup); err != nil {
				return err
			}
		}
	}

	if o.AllowUnencryptedAccessQueryParameter != nil {

		// query param allow_unencrypted_access
		var qrAllowUnencryptedAccess bool

		if o.AllowUnencryptedAccessQueryParameter != nil {
			qrAllowUnencryptedAccess = *o.AllowUnencryptedAccessQueryParameter
		}
		qAllowUnencryptedAccess := swag.FormatBool(qrAllowUnencryptedAccess)
		if qAllowUnencryptedAccess != "" {

			if err := r.SetQueryParam("allow_unencrypted_access", qAllowUnencryptedAccess); err != nil {
				return err
			}
		}
	}

	if o.ChangeNotifyQueryParameter != nil {

		// query param change_notify
		var qrChangeNotify bool

		if o.ChangeNotifyQueryParameter != nil {
			qrChangeNotify = *o.ChangeNotifyQueryParameter
		}
		qChangeNotify := swag.FormatBool(qrChangeNotify)
		if qChangeNotify != "" {

			if err := r.SetQueryParam("change_notify", qChangeNotify); err != nil {
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

	if o.ContinuouslyAvailableQueryParameter != nil {

		// query param continuously_available
		var qrContinuouslyAvailable bool

		if o.ContinuouslyAvailableQueryParameter != nil {
			qrContinuouslyAvailable = *o.ContinuouslyAvailableQueryParameter
		}
		qContinuouslyAvailable := swag.FormatBool(qrContinuouslyAvailable)
		if qContinuouslyAvailable != "" {

			if err := r.SetQueryParam("continuously_available", qContinuouslyAvailable); err != nil {
				return err
			}
		}
	}

	if o.DirUmaskQueryParameter != nil {

		// query param dir_umask
		var qrDirUmask int64

		if o.DirUmaskQueryParameter != nil {
			qrDirUmask = *o.DirUmaskQueryParameter
		}
		qDirUmask := swag.FormatInt64(qrDirUmask)
		if qDirUmask != "" {

			if err := r.SetQueryParam("dir_umask", qDirUmask); err != nil {
				return err
			}
		}
	}

	if o.EncryptionQueryParameter != nil {

		// query param encryption
		var qrEncryption bool

		if o.EncryptionQueryParameter != nil {
			qrEncryption = *o.EncryptionQueryParameter
		}
		qEncryption := swag.FormatBool(qrEncryption)
		if qEncryption != "" {

			if err := r.SetQueryParam("encryption", qEncryption); err != nil {
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

	if o.FileUmaskQueryParameter != nil {

		// query param file_umask
		var qrFileUmask int64

		if o.FileUmaskQueryParameter != nil {
			qrFileUmask = *o.FileUmaskQueryParameter
		}
		qFileUmask := swag.FormatInt64(qrFileUmask)
		if qFileUmask != "" {

			if err := r.SetQueryParam("file_umask", qFileUmask); err != nil {
				return err
			}
		}
	}

	if o.ForceGroupForCreateQueryParameter != nil {

		// query param force_group_for_create
		var qrForceGroupForCreate string

		if o.ForceGroupForCreateQueryParameter != nil {
			qrForceGroupForCreate = *o.ForceGroupForCreateQueryParameter
		}
		qForceGroupForCreate := qrForceGroupForCreate
		if qForceGroupForCreate != "" {

			if err := r.SetQueryParam("force_group_for_create", qForceGroupForCreate); err != nil {
				return err
			}
		}
	}

	if o.HomeDirectoryQueryParameter != nil {

		// query param home_directory
		var qrHomeDirectory bool

		if o.HomeDirectoryQueryParameter != nil {
			qrHomeDirectory = *o.HomeDirectoryQueryParameter
		}
		qHomeDirectory := swag.FormatBool(qrHomeDirectory)
		if qHomeDirectory != "" {

			if err := r.SetQueryParam("home_directory", qHomeDirectory); err != nil {
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

	if o.NamespaceCachingQueryParameter != nil {

		// query param namespace_caching
		var qrNamespaceCaching bool

		if o.NamespaceCachingQueryParameter != nil {
			qrNamespaceCaching = *o.NamespaceCachingQueryParameter
		}
		qNamespaceCaching := swag.FormatBool(qrNamespaceCaching)
		if qNamespaceCaching != "" {

			if err := r.SetQueryParam("namespace_caching", qNamespaceCaching); err != nil {
				return err
			}
		}
	}

	if o.NoStrictSecurityQueryParameter != nil {

		// query param no_strict_security
		var qrNoStrictSecurity bool

		if o.NoStrictSecurityQueryParameter != nil {
			qrNoStrictSecurity = *o.NoStrictSecurityQueryParameter
		}
		qNoStrictSecurity := swag.FormatBool(qrNoStrictSecurity)
		if qNoStrictSecurity != "" {

			if err := r.SetQueryParam("no_strict_security", qNoStrictSecurity); err != nil {
				return err
			}
		}
	}

	if o.OfflineFilesQueryParameter != nil {

		// query param offline_files
		var qrOfflineFiles string

		if o.OfflineFilesQueryParameter != nil {
			qrOfflineFiles = *o.OfflineFilesQueryParameter
		}
		qOfflineFiles := qrOfflineFiles
		if qOfflineFiles != "" {

			if err := r.SetQueryParam("offline_files", qOfflineFiles); err != nil {
				return err
			}
		}
	}

	if o.OplocksQueryParameter != nil {

		// query param oplocks
		var qrOplocks bool

		if o.OplocksQueryParameter != nil {
			qrOplocks = *o.OplocksQueryParameter
		}
		qOplocks := swag.FormatBool(qrOplocks)
		if qOplocks != "" {

			if err := r.SetQueryParam("oplocks", qOplocks); err != nil {
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

	if o.PathQueryParameter != nil {

		// query param path
		var qrPath string

		if o.PathQueryParameter != nil {
			qrPath = *o.PathQueryParameter
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
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

	if o.ShowSnapshotQueryParameter != nil {

		// query param show_snapshot
		var qrShowSnapshot bool

		if o.ShowSnapshotQueryParameter != nil {
			qrShowSnapshot = *o.ShowSnapshotQueryParameter
		}
		qShowSnapshot := swag.FormatBool(qrShowSnapshot)
		if qShowSnapshot != "" {

			if err := r.SetQueryParam("show_snapshot", qShowSnapshot); err != nil {
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

	if o.UnixSymlinkQueryParameter != nil {

		// query param unix_symlink
		var qrUnixSymlink string

		if o.UnixSymlinkQueryParameter != nil {
			qrUnixSymlink = *o.UnixSymlinkQueryParameter
		}
		qUnixSymlink := qrUnixSymlink
		if qUnixSymlink != "" {

			if err := r.SetQueryParam("unix_symlink", qUnixSymlink); err != nil {
				return err
			}
		}
	}

	if o.VolumeNameQueryParameter != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeNameQueryParameter != nil {
			qrVolumeName = *o.VolumeNameQueryParameter
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	if o.VolumeUUIDQueryParameter != nil {

		// query param volume.uuid
		var qrVolumeUUID string

		if o.VolumeUUIDQueryParameter != nil {
			qrVolumeUUID = *o.VolumeUUIDQueryParameter
		}
		qVolumeUUID := qrVolumeUUID
		if qVolumeUUID != "" {

			if err := r.SetQueryParam("volume.uuid", qVolumeUUID); err != nil {
				return err
			}
		}
	}

	if o.VscanProfileQueryParameter != nil {

		// query param vscan_profile
		var qrVscanProfile string

		if o.VscanProfileQueryParameter != nil {
			qrVscanProfile = *o.VscanProfileQueryParameter
		}
		qVscanProfile := qrVscanProfile
		if qVscanProfile != "" {

			if err := r.SetQueryParam("vscan_profile", qVscanProfile); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsShareCollectionGet binds the parameter fields
func (o *CifsShareCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCifsShareCollectionGet binds the parameter order_by
func (o *CifsShareCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
