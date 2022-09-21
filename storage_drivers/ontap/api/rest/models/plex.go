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

// Plex plex
//
// swagger:model plex
type Plex struct {

	// aggregate
	Aggregate *PlexAggregate `json:"aggregate,omitempty"`

	// Plex name
	// Example: plex0
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Plex is online
	// Read Only: true
	Online *bool `json:"online,omitempty"`

	// SyncMirror pool assignment
	// Read Only: true
	// Enum: [pool0 pool1]
	Pool string `json:"pool,omitempty"`

	// raid groups
	// Read Only: true
	RaidGroups []*RaidGroup `json:"raid_groups,omitempty"`

	// resync
	Resync *PlexResync `json:"resync,omitempty"`

	// Plex state
	// Read Only: true
	// Enum: [normal failed out_of_date]
	State string `json:"state,omitempty"`
}

// Validate validates this plex
func (m *Plex) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaidGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResync(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plex) validateAggregate(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregate) { // not required
		return nil
	}

	if m.Aggregate != nil {
		if err := m.Aggregate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

var plexTypePoolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pool0","pool1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		plexTypePoolPropEnum = append(plexTypePoolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// plex
	// Plex
	// pool
	// Pool
	// pool0
	// END DEBUGGING
	// PlexPoolPool0 captures enum value "pool0"
	PlexPoolPool0 string = "pool0"

	// BEGIN DEBUGGING
	// plex
	// Plex
	// pool
	// Pool
	// pool1
	// END DEBUGGING
	// PlexPoolPool1 captures enum value "pool1"
	PlexPoolPool1 string = "pool1"
)

// prop value enum
func (m *Plex) validatePoolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, plexTypePoolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Plex) validatePool(formats strfmt.Registry) error {
	if swag.IsZero(m.Pool) { // not required
		return nil
	}

	// value enum
	if err := m.validatePoolEnum("pool", "body", m.Pool); err != nil {
		return err
	}

	return nil
}

func (m *Plex) validateRaidGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.RaidGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.RaidGroups); i++ {
		if swag.IsZero(m.RaidGroups[i]) { // not required
			continue
		}

		if m.RaidGroups[i] != nil {
			if err := m.RaidGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raid_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Plex) validateResync(formats strfmt.Registry) error {
	if swag.IsZero(m.Resync) { // not required
		return nil
	}

	if m.Resync != nil {
		if err := m.Resync.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resync")
			}
			return err
		}
	}

	return nil
}

var plexTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["normal","failed","out_of_date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		plexTypeStatePropEnum = append(plexTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// plex
	// Plex
	// state
	// State
	// normal
	// END DEBUGGING
	// PlexStateNormal captures enum value "normal"
	PlexStateNormal string = "normal"

	// BEGIN DEBUGGING
	// plex
	// Plex
	// state
	// State
	// failed
	// END DEBUGGING
	// PlexStateFailed captures enum value "failed"
	PlexStateFailed string = "failed"

	// BEGIN DEBUGGING
	// plex
	// Plex
	// state
	// State
	// out_of_date
	// END DEBUGGING
	// PlexStateOutOfDate captures enum value "out_of_date"
	PlexStateOutOfDate string = "out_of_date"
)

// prop value enum
func (m *Plex) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, plexTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Plex) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this plex based on the context it is used
func (m *Plex) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaidGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResync(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plex) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregate != nil {
		if err := m.Aggregate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

func (m *Plex) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Plex) contextValidateOnline(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "online", "body", m.Online); err != nil {
		return err
	}

	return nil
}

func (m *Plex) contextValidatePool(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "pool", "body", string(m.Pool)); err != nil {
		return err
	}

	return nil
}

func (m *Plex) contextValidateRaidGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "raid_groups", "body", []*RaidGroup(m.RaidGroups)); err != nil {
		return err
	}

	for i := 0; i < len(m.RaidGroups); i++ {

		if m.RaidGroups[i] != nil {
			if err := m.RaidGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raid_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Plex) contextValidateResync(ctx context.Context, formats strfmt.Registry) error {

	if m.Resync != nil {
		if err := m.Resync.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resync")
			}
			return err
		}
	}

	return nil
}

func (m *Plex) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plex) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plex) UnmarshalBinary(b []byte) error {
	var res Plex
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlexAggregate plex aggregate
//
// swagger:model PlexAggregate
type PlexAggregate struct {

	// links
	Links *PlexAggregateLinks `json:"_links,omitempty"`

	// name
	// Example: aggr1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this plex aggregate
func (m *PlexAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlexAggregate) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plex aggregate based on the context it is used
func (m *PlexAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlexAggregate) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlexAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlexAggregate) UnmarshalBinary(b []byte) error {
	var res PlexAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlexAggregateLinks plex aggregate links
//
// swagger:model PlexAggregateLinks
type PlexAggregateLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this plex aggregate links
func (m *PlexAggregateLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlexAggregateLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plex aggregate links based on the context it is used
func (m *PlexAggregateLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlexAggregateLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlexAggregateLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlexAggregateLinks) UnmarshalBinary(b []byte) error {
	var res PlexAggregateLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlexResync plex resync
//
// swagger:model PlexResync
type PlexResync struct {

	// Plex is being resynchronized to its mirrored plex
	// Read Only: true
	Active *bool `json:"active,omitempty"`

	// Plex resyncing level
	// Read Only: true
	// Enum: [full incremental]
	Level string `json:"level,omitempty"`

	// Plex resyncing percentage
	// Example: 10
	// Read Only: true
	Percent int64 `json:"percent,omitempty"`
}

// Validate validates this plex resync
func (m *PlexResync) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var plexResyncTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","incremental"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		plexResyncTypeLevelPropEnum = append(plexResyncTypeLevelPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// PlexResync
	// PlexResync
	// level
	// Level
	// full
	// END DEBUGGING
	// PlexResyncLevelFull captures enum value "full"
	PlexResyncLevelFull string = "full"

	// BEGIN DEBUGGING
	// PlexResync
	// PlexResync
	// level
	// Level
	// incremental
	// END DEBUGGING
	// PlexResyncLevelIncremental captures enum value "incremental"
	PlexResyncLevelIncremental string = "incremental"
)

// prop value enum
func (m *PlexResync) validateLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, plexResyncTypeLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlexResync) validateLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("resync"+"."+"level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this plex resync based on the context it is used
func (m *PlexResync) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePercent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlexResync) contextValidateActive(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "resync"+"."+"active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *PlexResync) contextValidateLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "resync"+"."+"level", "body", string(m.Level)); err != nil {
		return err
	}

	return nil
}

func (m *PlexResync) contextValidatePercent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "resync"+"."+"percent", "body", int64(m.Percent)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlexResync) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlexResync) UnmarshalBinary(b []byte) error {
	var res PlexResync
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
