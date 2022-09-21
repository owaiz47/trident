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

// NisService nis service
//
// swagger:model nis_service
type NisService struct {

	// links
	Links *NisServiceLinks `json:"_links,omitempty"`

	// An array of objects where each object represents the NIS server and it's status for a given NIS domain. It is an advanced field.
	BindingDetails []*NisServiceBindingDetailsItems0 `json:"binding_details,omitempty"`

	// bound servers
	// Read Only: true
	BoundServers []string `json:"bound_servers,omitempty"`

	// The NIS domain to which this configuration belongs.
	//
	// Max Length: 64
	// Min Length: 1
	Domain string `json:"domain,omitempty"`

	// A list of hostnames or IP addresses of NIS servers used
	// by the NIS domain configuration.
	//
	// Max Items: 10
	Servers []string `json:"servers,omitempty"`

	// svm
	Svm *NisServiceSvm `json:"svm,omitempty"`
}

// Validate validates this nis service
func (m *NisService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindingDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoundServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
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

func (m *NisService) validateLinks(formats strfmt.Registry) error {
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

func (m *NisService) validateBindingDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.BindingDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.BindingDetails); i++ {
		if swag.IsZero(m.BindingDetails[i]) { // not required
			continue
		}

		if m.BindingDetails[i] != nil {
			if err := m.BindingDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("binding_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NisService) validateBoundServers(formats strfmt.Registry) error {
	if swag.IsZero(m.BoundServers) { // not required
		return nil
	}

	for i := 0; i < len(m.BoundServers); i++ {

		if err := validate.MinLength("bound_servers"+"."+strconv.Itoa(i), "body", m.BoundServers[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength("bound_servers"+"."+strconv.Itoa(i), "body", m.BoundServers[i], 255); err != nil {
			return err
		}

	}

	return nil
}

func (m *NisService) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if err := validate.MinLength("domain", "body", m.Domain, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("domain", "body", m.Domain, 64); err != nil {
		return err
	}

	return nil
}

func (m *NisService) validateServers(formats strfmt.Registry) error {
	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	iServersSize := int64(len(m.Servers))

	if err := validate.MaxItems("servers", "body", iServersSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.Servers); i++ {

		if err := validate.MinLength("servers"+"."+strconv.Itoa(i), "body", m.Servers[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength("servers"+"."+strconv.Itoa(i), "body", m.Servers[i], 255); err != nil {
			return err
		}

	}

	return nil
}

func (m *NisService) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this nis service based on the context it is used
func (m *NisService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBindingDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBoundServers(ctx, formats); err != nil {
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

func (m *NisService) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NisService) contextValidateBindingDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BindingDetails); i++ {

		if m.BindingDetails[i] != nil {
			if err := m.BindingDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("binding_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NisService) contextValidateBoundServers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "bound_servers", "body", []string(m.BoundServers)); err != nil {
		return err
	}

	return nil
}

func (m *NisService) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NisService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NisService) UnmarshalBinary(b []byte) error {
	var res NisService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NisServiceBindingDetailsItems0 nis service binding details items0
//
// swagger:model NisServiceBindingDetailsItems0
type NisServiceBindingDetailsItems0 struct {

	// Hostname/IP address of the NIS server in the domain.
	// Max Length: 255
	// Min Length: 1
	Server string `json:"server,omitempty"`

	// status
	Status *BindingStatus `json:"status,omitempty"`
}

// Validate validates this nis service binding details items0
func (m *NisServiceBindingDetailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisServiceBindingDetailsItems0) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if err := validate.MinLength("server", "body", m.Server, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("server", "body", m.Server, 255); err != nil {
		return err
	}

	return nil
}

func (m *NisServiceBindingDetailsItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nis service binding details items0 based on the context it is used
func (m *NisServiceBindingDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisServiceBindingDetailsItems0) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NisServiceBindingDetailsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NisServiceBindingDetailsItems0) UnmarshalBinary(b []byte) error {
	var res NisServiceBindingDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NisServiceLinks nis service links
//
// swagger:model NisServiceLinks
type NisServiceLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nis service links
func (m *NisServiceLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisServiceLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this nis service links based on the context it is used
func (m *NisServiceLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisServiceLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NisServiceLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NisServiceLinks) UnmarshalBinary(b []byte) error {
	var res NisServiceLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NisServiceSvm nis service svm
//
// swagger:model NisServiceSvm
type NisServiceSvm struct {

	// links
	Links *NisServiceSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this nis service svm
func (m *NisServiceSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisServiceSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this nis service svm based on the context it is used
func (m *NisServiceSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisServiceSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NisServiceSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NisServiceSvm) UnmarshalBinary(b []byte) error {
	var res NisServiceSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NisServiceSvmLinks nis service svm links
//
// swagger:model NisServiceSvmLinks
type NisServiceSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nis service svm links
func (m *NisServiceSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisServiceSvmLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this nis service svm links based on the context it is used
func (m *NisServiceSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NisServiceSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NisServiceSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NisServiceSvmLinks) UnmarshalBinary(b []byte) error {
	var res NisServiceSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
