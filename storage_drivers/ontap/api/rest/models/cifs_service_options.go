// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CifsServiceOptions cifs service options
//
// swagger:model cifs_service_options
type CifsServiceOptions struct {

	// Specifies whether or not Administrator can be mapped to the UNIX user "root".
	//
	AdminToRootMapping *bool `json:"admin_to_root_mapping,omitempty"`

	// Specifies whether or not the CIFS server supports the advanced sparse file capabilities. This allows
	// CIFS clients to query the allocated ranges of a file and to write zeroes or free data blocks for ranges
	// of a file.
	//
	AdvancedSparseFile *bool `json:"advanced_sparse_file,omitempty"`

	// Specifies whether or not to enable the Copy Offload feature. This feature enables direct
	// data transfers within or between compatible storage devices without transferring the data
	// through the host computer.<br/>
	// Note that this will also enable/disable the direct copy feature accordingly.
	//
	CopyOffload *bool `json:"copy_offload,omitempty"`

	// Specifies whether or not fake open support is enabled. This parameter allows you to optimize the
	// open and close requests coming from SMB 2 clients.
	//
	FakeOpen *bool `json:"fake_open,omitempty"`

	// Specifies whether or not the trim requests (FSCTL_FILE_LEVEL_TRIM) are supported on the CIFS server.
	//
	FsctlTrim *bool `json:"fsctl_trim,omitempty"`

	// Specifies whether or not the reparse point support is enabled. When enabled the CIFS server
	// exposes junction points to Windows clients as reparse points. This parameter is only active
	// if the client has negotiated use of the SMB 2 or SMB 3 protocol. This parameter is not supported
	// for SVMs with Infinite Volume.
	//
	JunctionReparse *bool `json:"junction_reparse,omitempty"`

	// Specifies whether or not SMB clients can send reads up to 1 MB in size.
	LargeMtu *bool `json:"large_mtu,omitempty"`

	// Specifies whether or not the CIFS server supports Multichannel.
	Multichannel *bool `json:"multichannel,omitempty"`

	// Specifies a Windows User or Group name that should be mapped in case of a NULL user
	// value.
	//
	NullUserWindowsName string `json:"null_user_windows_name,omitempty"`

	// Specifies whether or not the path component cache is enabled on the CIFS server.
	PathComponentCache *bool `json:"path_component_cache,omitempty"`

	// Specifies whether or not to refer clients to more optimal LIFs. When enabled, it automatically
	// refers clients to a data LIF local to the node which hosts the root of the requested share.
	//
	Referral *bool `json:"referral,omitempty"`

	// Specifies whether or not to enable the Shadowcopy Feature. This feature enables
	// to take share-based backup copies of data that is in a data-consistent state at
	// a specific point in time where the data is accessed over SMB 3.0 shares.
	//
	Shadowcopy *bool `json:"shadowcopy,omitempty"`

	// Specifies the maximum level of subdirectories on which ONTAP should create shadow copies.
	ShadowcopyDirDepth *int64 `json:"shadowcopy_dir_depth,omitempty"`

	// Specifies the maximum number of outstanding requests on a CIFS connection.
	// Example: 128
	// Maximum: 8192
	// Minimum: 2
	SmbCredits int64 `json:"smb_credits,omitempty"`

	// Specifies the CIFS protocol versions for which the widelink is reported as reparse point.
	//
	// Example: ["smb1"]
	WidelinkReparseVersions []string `json:"widelink_reparse_versions,omitempty"`
}

// Validate validates this cifs service options
func (m *CifsServiceOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSmbCredits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidelinkReparseVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsServiceOptions) validateSmbCredits(formats strfmt.Registry) error {
	if swag.IsZero(m.SmbCredits) { // not required
		return nil
	}

	if err := validate.MinimumInt("smb_credits", "body", m.SmbCredits, 2, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("smb_credits", "body", m.SmbCredits, 8192, false); err != nil {
		return err
	}

	return nil
}

var cifsServiceOptionsWidelinkReparseVersionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["smb1","smb2","smb2_1","smb3","smb3_1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsServiceOptionsWidelinkReparseVersionsItemsEnum = append(cifsServiceOptionsWidelinkReparseVersionsItemsEnum, v)
	}
}

func (m *CifsServiceOptions) validateWidelinkReparseVersionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsServiceOptionsWidelinkReparseVersionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsServiceOptions) validateWidelinkReparseVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.WidelinkReparseVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.WidelinkReparseVersions); i++ {

		// value enum
		if err := m.validateWidelinkReparseVersionsItemsEnum("widelink_reparse_versions"+"."+strconv.Itoa(i), "body", m.WidelinkReparseVersions[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this cifs service options based on context it is used
func (m *CifsServiceOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CifsServiceOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsServiceOptions) UnmarshalBinary(b []byte) error {
	var res CifsServiceOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
