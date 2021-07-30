// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NameMapping Name mapping is used to map CIFS identities to UNIX identities, Kerberos identities to UNIX identities, and UNIX identities to CIFS identities. It needs this information to obtain user credentials and provide proper file access regardless of whether they are connecting from an NFS client or a CIFS client.
//
// swagger:model name_mapping
type NameMapping struct {

	// links
	Links *NameMappingLinks `json:"_links,omitempty"`

	// Client workstation IP Address which is matched when searching for the pattern.
	//   You can specify the value in any of the following formats:
	// * As an IPv4 address with a subnet mask expressed as a number of bits; for instance, 10.1.12.0/24
	// * As an IPv6 address with a subnet mask expressed as a number of bits; for instance, fd20:8b1e:b255:4071::/64
	// * As an IPv4 address with a network mask; for instance, 10.1.16.0/255.255.255.0
	// * As a hostname
	//
	// Example: 10.254.101.111/28
	ClientMatch string `json:"client_match,omitempty"`

	// Direction in which the name mapping is applied. The possible values are:
	//   * krb_unix  - Kerberos principal name to UNIX user name
	//   * win_unix  - Windows user name to UNIX user name
	//   * unix_win  - UNIX user name to Windows user name mapping
	//
	// Example: win_unix
	// Enum: [win_unix unix_win krb_unix]
	Direction string `json:"direction,omitempty"`

	// Position in the list of name mappings.
	// Example: 1
	// Maximum: 2.147483647e+09
	// Minimum: 1
	Index int64 `json:"index,omitempty"`

	// Pattern used to match the name while searching for a name that can be used as a replacement. The pattern is a UNIX-style regular expression. Regular expressions are case-insensitive when mapping from Windows to UNIX, and they are case-sensitive for mappings from Kerberos to UNIX and UNIX to Windows.
	// Example: ENGCIFS_AD_USER
	Pattern string `json:"pattern,omitempty"`

	// The name that is used as a replacement, if the pattern associated with this entry matches.
	// Example: unix_user1
	Replacement string `json:"replacement,omitempty"`

	// svm
	Svm *NameMappingSvm `json:"svm,omitempty"`
}

// Validate validates this name mapping
func (m *NameMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameMapping) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

var nameMappingTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["win_unix","unix_win","krb_unix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nameMappingTypeDirectionPropEnum = append(nameMappingTypeDirectionPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// name_mapping
	// NameMapping
	// direction
	// Direction
	// win_unix
	// END RIPPY DEBUGGING
	// NameMappingDirectionWinUnix captures enum value "win_unix"
	NameMappingDirectionWinUnix string = "win_unix"

	// BEGIN RIPPY DEBUGGING
	// name_mapping
	// NameMapping
	// direction
	// Direction
	// unix_win
	// END RIPPY DEBUGGING
	// NameMappingDirectionUnixWin captures enum value "unix_win"
	NameMappingDirectionUnixWin string = "unix_win"

	// BEGIN RIPPY DEBUGGING
	// name_mapping
	// NameMapping
	// direction
	// Direction
	// krb_unix
	// END RIPPY DEBUGGING
	// NameMappingDirectionKrbUnix captures enum value "krb_unix"
	NameMappingDirectionKrbUnix string = "krb_unix"
)

// prop value enum
func (m *NameMapping) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nameMappingTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NameMapping) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *NameMapping) validateIndex(formats strfmt.Registry) error {
	if swag.IsZero(m.Index) { // not required
		return nil
	}

	if err := validate.MinimumInt("index", "body", m.Index, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("index", "body", m.Index, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *NameMapping) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this name mapping based on the context it is used
func (m *NameMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameMapping) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *NameMapping) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameMapping) UnmarshalBinary(b []byte) error {
	var res NameMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NameMappingLinks name mapping links
//
// swagger:model NameMappingLinks
type NameMappingLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this name mapping links
func (m *NameMappingLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameMappingLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this name mapping links based on the context it is used
func (m *NameMappingLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameMappingLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameMappingLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameMappingLinks) UnmarshalBinary(b []byte) error {
	var res NameMappingLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NameMappingSvm name mapping svm
//
// swagger:model NameMappingSvm
type NameMappingSvm struct {

	// links
	Links *NameMappingSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this name mapping svm
func (m *NameMappingSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameMappingSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this name mapping svm based on the context it is used
func (m *NameMappingSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameMappingSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameMappingSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameMappingSvm) UnmarshalBinary(b []byte) error {
	var res NameMappingSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NameMappingSvmLinks name mapping svm links
//
// swagger:model NameMappingSvmLinks
type NameMappingSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this name mapping svm links
func (m *NameMappingSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameMappingSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this name mapping svm links based on the context it is used
func (m *NameMappingSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameMappingSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameMappingSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameMappingSvmLinks) UnmarshalBinary(b []byte) error {
	var res NameMappingSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY