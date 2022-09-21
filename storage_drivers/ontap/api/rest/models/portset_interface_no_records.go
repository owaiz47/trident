// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PortsetInterfaceNoRecords A container for either a Fibre Channel network interface or an IP network interface. On POST `fc` and `ip` are mutually exclusive.
//
//
// swagger:model portset_interface_no_records
type PortsetInterfaceNoRecords struct {

	// links
	Links *PortsetInterfaceNoRecordsLinks `json:"_links,omitempty"`

	// fc
	Fc *PortsetInterfaceNoRecordsFc `json:"fc,omitempty"`

	// ip
	IP *PortsetInterfaceNoRecordsIP `json:"ip,omitempty"`

	// The unique identifier of the network interface.
	//
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this portset interface no records
func (m *PortsetInterfaceNoRecords) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecords) validateLinks(formats strfmt.Registry) error {
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

func (m *PortsetInterfaceNoRecords) validateFc(formats strfmt.Registry) error {
	if swag.IsZero(m.Fc) { // not required
		return nil
	}

	if m.Fc != nil {
		if err := m.Fc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc")
			}
			return err
		}
	}

	return nil
}

func (m *PortsetInterfaceNoRecords) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portset interface no records based on the context it is used
func (m *PortsetInterfaceNoRecords) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecords) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortsetInterfaceNoRecords) contextValidateFc(ctx context.Context, formats strfmt.Registry) error {

	if m.Fc != nil {
		if err := m.Fc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc")
			}
			return err
		}
	}

	return nil
}

func (m *PortsetInterfaceNoRecords) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {
		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *PortsetInterfaceNoRecords) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortsetInterfaceNoRecords) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortsetInterfaceNoRecords) UnmarshalBinary(b []byte) error {
	var res PortsetInterfaceNoRecords
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortsetInterfaceNoRecordsFc An FC interface.
//
//
// swagger:model PortsetInterfaceNoRecordsFc
type PortsetInterfaceNoRecordsFc struct {

	// links
	Links *PortsetInterfaceNoRecordsFcLinks `json:"_links,omitempty"`

	// The name of the FC interface.
	//
	// Example: fc_lif1
	Name string `json:"name,omitempty"`

	// The unique identifier of the FC interface.
	//
	// Example: 3a09ab42-4da1-32cf-9d35-3385a6101a0b
	UUID string `json:"uuid,omitempty"`

	// The WWPN of the FC interface.
	//
	// Example: 20:00:00:50:56:b4:13:a8
	// Read Only: true
	Wwpn string `json:"wwpn,omitempty"`
}

// Validate validates this portset interface no records fc
func (m *PortsetInterfaceNoRecordsFc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsFc) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portset interface no records fc based on the context it is used
func (m *PortsetInterfaceNoRecordsFc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsFc) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *PortsetInterfaceNoRecordsFc) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fc"+"."+"wwpn", "body", string(m.Wwpn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsFc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsFc) UnmarshalBinary(b []byte) error {
	var res PortsetInterfaceNoRecordsFc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortsetInterfaceNoRecordsFcLinks portset interface no records fc links
//
// swagger:model PortsetInterfaceNoRecordsFcLinks
type PortsetInterfaceNoRecordsFcLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this portset interface no records fc links
func (m *PortsetInterfaceNoRecordsFcLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsFcLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portset interface no records fc links based on the context it is used
func (m *PortsetInterfaceNoRecordsFcLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsFcLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsFcLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsFcLinks) UnmarshalBinary(b []byte) error {
	var res PortsetInterfaceNoRecordsFcLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortsetInterfaceNoRecordsIP portset interface no records IP
//
// swagger:model PortsetInterfaceNoRecordsIP
type PortsetInterfaceNoRecordsIP struct {

	// links
	Links *PortsetInterfaceNoRecordsIPLinks `json:"_links,omitempty"`

	// ip
	IP *PortsetInterfaceNoRecordsIPIP `json:"ip,omitempty"`

	// The name of the interface. If only the name is provided, the SVM scope
	// must be provided by the object this object is embedded in.
	//
	// Example: lif1
	Name string `json:"name,omitempty"`

	// The UUID that uniquely identifies the interface.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this portset interface no records IP
func (m *PortsetInterfaceNoRecordsIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsIP) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *PortsetInterfaceNoRecordsIP) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "ip")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portset interface no records IP based on the context it is used
func (m *PortsetInterfaceNoRecordsIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsIP) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *PortsetInterfaceNoRecordsIP) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {
		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "ip")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsIP) UnmarshalBinary(b []byte) error {
	var res PortsetInterfaceNoRecordsIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortsetInterfaceNoRecordsIPIP IP information
//
// swagger:model PortsetInterfaceNoRecordsIPIP
type PortsetInterfaceNoRecordsIPIP struct {

	// address
	Address IPAddressReadonly `json:"address,omitempty"`
}

// Validate validates this portset interface no records IP IP
func (m *PortsetInterfaceNoRecordsIPIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsIPIP) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := m.Address.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ip" + "." + "ip" + "." + "address")
		}
		return err
	}

	return nil
}

// ContextValidate validate this portset interface no records IP IP based on the context it is used
func (m *PortsetInterfaceNoRecordsIPIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsIPIP) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Address.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ip" + "." + "ip" + "." + "address")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsIPIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsIPIP) UnmarshalBinary(b []byte) error {
	var res PortsetInterfaceNoRecordsIPIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortsetInterfaceNoRecordsIPLinks portset interface no records IP links
//
// swagger:model PortsetInterfaceNoRecordsIPLinks
type PortsetInterfaceNoRecordsIPLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this portset interface no records IP links
func (m *PortsetInterfaceNoRecordsIPLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsIPLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this portset interface no records IP links based on the context it is used
func (m *PortsetInterfaceNoRecordsIPLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsIPLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsIPLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsIPLinks) UnmarshalBinary(b []byte) error {
	var res PortsetInterfaceNoRecordsIPLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortsetInterfaceNoRecordsLinks portset interface no records links
//
// swagger:model PortsetInterfaceNoRecordsLinks
type PortsetInterfaceNoRecordsLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this portset interface no records links
func (m *PortsetInterfaceNoRecordsLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this portset interface no records links based on the context it is used
func (m *PortsetInterfaceNoRecordsLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortsetInterfaceNoRecordsLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PortsetInterfaceNoRecordsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortsetInterfaceNoRecordsLinks) UnmarshalBinary(b []byte) error {
	var res PortsetInterfaceNoRecordsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
