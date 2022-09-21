// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeProtocolRawPerformanceStatOther volume protocol raw performance stat other
//
// swagger:model volume_protocol_raw_performance_stat_other
type VolumeProtocolRawPerformanceStatOther struct {

	// Number of operations of the given type performed on this volume.
	// Example: 1000
	Count int64 `json:"count,omitempty"`

	// The raw data component latency in microseconds measured within ONTAP for all operations of the given type.
	// Example: 200
	TotalTime int64 `json:"total_time,omitempty"`
}

// Validate validates this volume protocol raw performance stat other
func (m *VolumeProtocolRawPerformanceStatOther) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume protocol raw performance stat other based on context it is used
func (m *VolumeProtocolRawPerformanceStatOther) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeProtocolRawPerformanceStatOther) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeProtocolRawPerformanceStatOther) UnmarshalBinary(b []byte) error {
	var res VolumeProtocolRawPerformanceStatOther
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
