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

// KeyServerStateArray A container for holding an array of key server connectivity states for each node.
//
//
// swagger:model key_server_state_array
type KeyServerStateArray struct {

	// Set to true when key server connectivity state is available on all nodes of the cluster.
	// Read Only: true
	ClusterAvailability *bool `json:"cluster_availability,omitempty"`

	// An array of key server connectivity states for each node.
	//
	// Read Only: true
	Records []*KeyServerState `json:"records,omitempty"`
}

// Validate validates this key server state array
func (m *KeyServerStateArray) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyServerStateArray) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this key server state array based on the context it is used
func (m *KeyServerStateArray) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyServerStateArray) contextValidateClusterAvailability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cluster_availability", "body", m.ClusterAvailability); err != nil {
		return err
	}

	return nil
}

func (m *KeyServerStateArray) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "records", "body", []*KeyServerState(m.Records)); err != nil {
		return err
	}

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyServerStateArray) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyServerStateArray) UnmarshalBinary(b []byte) error {
	var res KeyServerStateArray
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY