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

// TopMetricsSvmFileResponse top metrics svm file response
//
// swagger:model top_metrics_svm_file_response
type TopMetricsSvmFileResponse struct {

	// links
	Links *TopMetricsSvmFileResponseLinks `json:"_links,omitempty"`

	// List of volumes that are not included in the SVM activity tracking REST API.
	ExcludedVolumes []*TopMetricsSvmFileExcludedVolume `json:"excluded_volumes,omitempty"`

	// notice
	Notice *TopMetricsSvmFileResponseNotice `json:"notice,omitempty"`

	// Number of records.
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*TopMetricsSvmFile `json:"records,omitempty"`
}

// Validate validates this top metrics svm file response
func (m *TopMetricsSvmFileResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludedVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmFileResponse) validateExcludedVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludedVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.ExcludedVolumes); i++ {
		if swag.IsZero(m.ExcludedVolumes[i]) { // not required
			continue
		}

		if m.ExcludedVolumes[i] != nil {
			if err := m.ExcludedVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excluded_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TopMetricsSvmFileResponse) validateNotice(formats strfmt.Registry) error {
	if swag.IsZero(m.Notice) { // not required
		return nil
	}

	if m.Notice != nil {
		if err := m.Notice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notice")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFileResponse) validateRecords(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm file response based on the context it is used
func (m *TopMetricsSvmFileResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExcludedVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotice(ctx, formats); err != nil {
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

func (m *TopMetricsSvmFileResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsSvmFileResponse) contextValidateExcludedVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExcludedVolumes); i++ {

		if m.ExcludedVolumes[i] != nil {
			if err := m.ExcludedVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excluded_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TopMetricsSvmFileResponse) contextValidateNotice(ctx context.Context, formats strfmt.Registry) error {

	if m.Notice != nil {
		if err := m.Notice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notice")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFileResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmFileResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileResponse) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileResponseLinks top metrics svm file response links
//
// swagger:model TopMetricsSvmFileResponseLinks
type TopMetricsSvmFileResponseLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics svm file response links
func (m *TopMetricsSvmFileResponseLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileResponseLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFileResponseLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm file response links based on the context it is used
func (m *TopMetricsSvmFileResponseLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileResponseLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmFileResponseLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmFileResponseLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileResponseLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileResponseLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmFileResponseNotice Optional field that indicates why no records are returned by the SVM activity tracking REST API.
//
// swagger:model TopMetricsSvmFileResponseNotice
type TopMetricsSvmFileResponseNotice struct {

	// Warning code indicating why no records are returned.
	// Example: 111411207
	// Read Only: true
	Code string `json:"code,omitempty"`

	// Details why no records are returned.
	// Example: The volume is offline.
	// Read Only: true
	Message string `json:"message,omitempty"`
}

// Validate validates this top metrics svm file response notice
func (m *TopMetricsSvmFileResponseNotice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this top metrics svm file response notice based on the context it is used
func (m *TopMetricsSvmFileResponseNotice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmFileResponseNotice) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "notice"+"."+"code", "body", string(m.Code)); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmFileResponseNotice) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "notice"+"."+"message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmFileResponseNotice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmFileResponseNotice) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmFileResponseNotice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
