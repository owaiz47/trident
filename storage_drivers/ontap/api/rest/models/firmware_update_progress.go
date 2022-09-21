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

// FirmwareUpdateProgress firmware update progress
//
// swagger:model firmware_update_progress
type FirmwareUpdateProgress struct {

	// job
	Job *JobLink `json:"job,omitempty"`

	// update state
	UpdateState []*FirmwareUpdateProgressState `json:"update_state,omitempty"`

	// Specifies the type of update.
	// Read Only: true
	// Enum: [manual_update automatic_update]
	UpdateType string `json:"update_type,omitempty"`

	// zip file name
	// Example: disk_firmware.zip
	// Read Only: true
	ZipFileName string `json:"zip_file_name,omitempty"`
}

// Validate validates this firmware update progress
func (m *FirmwareUpdateProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareUpdateProgress) validateJob(formats strfmt.Registry) error {
	if swag.IsZero(m.Job) { // not required
		return nil
	}

	if m.Job != nil {
		if err := m.Job.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareUpdateProgress) validateUpdateState(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateState) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateState); i++ {
		if swag.IsZero(m.UpdateState[i]) { // not required
			continue
		}

		if m.UpdateState[i] != nil {
			if err := m.UpdateState[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_state" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var firmwareUpdateProgressTypeUpdateTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["manual_update","automatic_update"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		firmwareUpdateProgressTypeUpdateTypePropEnum = append(firmwareUpdateProgressTypeUpdateTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// firmware_update_progress
	// FirmwareUpdateProgress
	// update_type
	// UpdateType
	// manual_update
	// END DEBUGGING
	// FirmwareUpdateProgressUpdateTypeManualUpdate captures enum value "manual_update"
	FirmwareUpdateProgressUpdateTypeManualUpdate string = "manual_update"

	// BEGIN DEBUGGING
	// firmware_update_progress
	// FirmwareUpdateProgress
	// update_type
	// UpdateType
	// automatic_update
	// END DEBUGGING
	// FirmwareUpdateProgressUpdateTypeAutomaticUpdate captures enum value "automatic_update"
	FirmwareUpdateProgressUpdateTypeAutomaticUpdate string = "automatic_update"
)

// prop value enum
func (m *FirmwareUpdateProgress) validateUpdateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, firmwareUpdateProgressTypeUpdateTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FirmwareUpdateProgress) validateUpdateType(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUpdateTypeEnum("update_type", "body", m.UpdateType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this firmware update progress based on the context it is used
func (m *FirmwareUpdateProgress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZipFileName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareUpdateProgress) contextValidateJob(ctx context.Context, formats strfmt.Registry) error {

	if m.Job != nil {
		if err := m.Job.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareUpdateProgress) contextValidateUpdateState(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UpdateState); i++ {

		if m.UpdateState[i] != nil {
			if err := m.UpdateState[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_state" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FirmwareUpdateProgress) contextValidateUpdateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "update_type", "body", string(m.UpdateType)); err != nil {
		return err
	}

	return nil
}

func (m *FirmwareUpdateProgress) contextValidateZipFileName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "zip_file_name", "body", string(m.ZipFileName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareUpdateProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareUpdateProgress) UnmarshalBinary(b []byte) error {
	var res FirmwareUpdateProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
