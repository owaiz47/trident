// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FpolicyPolicies fpolicy policies
//
// swagger:model fpolicy_policies
type FpolicyPolicies struct {

	// Specifies if the policy is enabled on the SVM or not. If no value is
	// mentioned for this field but priority is set, then this policy will be enabled.
	//
	Enabled bool `json:"enabled,omitempty"`

	// engine
	Engine *FpolicyEngineReference `json:"engine,omitempty"`

	// events
	// Example: ["event_nfs_close","event_open"]
	Events []*FpolicyEventReference `json:"events,omitempty"`

	// Specifies what action to take on a file access event in a case when all primary and secondary servers are down or no response is received from the FPolicy servers within a given timeout period. When this parameter is set to true, file access events will be denied under these circumstances.
	Mandatory *bool `json:"mandatory,omitempty"`

	// Specifies the name of the policy.
	// Example: fp_policy_1
	Name string `json:"name,omitempty"`

	// Specifies whether passthrough-read should be allowed for FPolicy servers
	// registered for the policy. Passthrough-read is a way to read data for
	// offline files without restoring the files to primary storage. Offline
	// files are files that have been moved to secondary storage.
	//
	PassthroughRead *bool `json:"passthrough_read,omitempty"`

	// Specifies the priority that is assigned to this policy.
	// Maximum: 10
	// Minimum: 1
	Priority int64 `json:"priority,omitempty"`

	// Specifies the privileged user name for accessing files on the cluster
	// using a separate data channel with privileged access. The input for
	// this field should be in "domain\username" format.
	//
	// Example: mydomain\\testuser
	PrivilegedUser string `json:"privileged_user,omitempty"`

	// scope
	Scope *FpolicyPoliciesScope `json:"scope,omitempty"`
}

// Validate validates this fpolicy policies
func (m *FpolicyPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyPolicies) validateEngine(formats strfmt.Registry) error {
	if swag.IsZero(m.Engine) { // not required
		return nil
	}

	if m.Engine != nil {
		if err := m.Engine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("engine")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyPolicies) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FpolicyPolicies) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if err := validate.MinimumInt("priority", "body", m.Priority, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("priority", "body", m.Priority, 10, false); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyPolicies) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fpolicy policies based on the context it is used
func (m *FpolicyPolicies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEngine(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyPolicies) contextValidateEngine(ctx context.Context, formats strfmt.Registry) error {

	if m.Engine != nil {
		if err := m.Engine.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("engine")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyPolicies) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {
			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FpolicyPolicies) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {
		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyPolicies) UnmarshalBinary(b []byte) error {
	var res FpolicyPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyPoliciesScope fpolicy policies scope
//
// swagger:model FpolicyPoliciesScope
type FpolicyPoliciesScope struct {

	// Specifies whether the file name extension checks also apply to directory objects. If this parameter is set to true,
	// the directory objects are subjected to the same extension checks as regular files. If this parameter is set to false,
	// the directory names are not matched for extensions and notifications are sent for directories even if their name
	// extensions do not match. Default is false.
	//
	CheckExtensionsOnDirectories *bool `json:"check_extensions_on_directories,omitempty"`

	// exclude export policies
	ExcludeExportPolicies []string `json:"exclude_export_policies,omitempty"`

	// exclude extension
	ExcludeExtension []string `json:"exclude_extension,omitempty"`

	// exclude shares
	ExcludeShares []string `json:"exclude_shares,omitempty"`

	// exclude volumes
	// Example: ["vol1","vol_svm1","*"]
	ExcludeVolumes []string `json:"exclude_volumes,omitempty"`

	// include export policies
	IncludeExportPolicies []string `json:"include_export_policies,omitempty"`

	// include extension
	IncludeExtension []string `json:"include_extension,omitempty"`

	// include shares
	// Example: ["sh1","share_cifs"]
	IncludeShares []string `json:"include_shares,omitempty"`

	// include volumes
	// Example: ["vol1","vol_svm1"]
	IncludeVolumes []string `json:"include_volumes,omitempty"`

	// Specifies whether the extension checks also apply to objects with no extension. If this parameter is set to true,
	// all objects with or without extensions are monitored. Default is false.
	//
	ObjectMonitoringWithNoExtension *bool `json:"object_monitoring_with_no_extension,omitempty"`
}

// Validate validates this fpolicy policies scope
func (m *FpolicyPoliciesScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy policies scope based on context it is used
func (m *FpolicyPoliciesScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyPoliciesScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyPoliciesScope) UnmarshalBinary(b []byte) error {
	var res FpolicyPoliciesScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
