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

// FpolicyEngines Defines how ONTAP makes and manages connections to external FPolicy servers.
//
// swagger:model fpolicy_engines
type FpolicyEngines struct {

	// buffer size
	BufferSize *FpolicyEnginesBufferSize `json:"buffer_size,omitempty"`

	// certificate
	Certificate *FpolicyEnginesCertificate `json:"certificate,omitempty"`

	// The format for the notification messages sent to the FPolicy servers.
	//   The possible values are:
	//     * xml  - Notifications sent to the FPolicy server will be formatted using the XML schema.
	//     * protobuf - Notifications sent to the FPolicy server will be formatted using Protobuf schema, which is a binary form.
	//
	// Enum: [xml protobuf]
	Format *string `json:"format,omitempty"`

	// Specifies the maximum number of outstanding requests for the FPolicy server. It is used to specify maximum outstanding requests that will be queued up for the FPolicy server. The value for this field must be between 1 and 10000. The default value can be 500 , 1000 or 2000 depending on the hardware platform.
	// Example: 500
	// Maximum: 10000
	// Minimum: 1
	MaxServerRequests int64 `json:"max_server_requests,omitempty"`

	// Specifies the name to assign to the external server configuration.
	// Example: fp_ex_eng
	Name string `json:"name,omitempty"`

	// Port number of the FPolicy server application.
	// Example: 9876
	Port int64 `json:"port,omitempty"`

	// primary servers
	// Example: ["10.132.145.20","10.140.101.109"]
	PrimaryServers []string `json:"primary_servers,omitempty"`

	// Specifies the ISO-8601 timeout duration for a screen request to be aborted by a storage appliance. The allowed range is between 0 to 200 seconds.
	// Example: PT40S
	RequestAbortTimeout *string `json:"request_abort_timeout,omitempty"`

	// Specifies the ISO-8601 timeout duration for a screen request to be processed by an FPolicy server. The allowed range is between 0 to 100 seconds.
	// Example: PT20S
	RequestCancelTimeout *string `json:"request_cancel_timeout,omitempty"`

	// resiliency
	Resiliency *FpolicyEnginesResiliency `json:"resiliency,omitempty"`

	// secondary servers
	// Example: ["10.132.145.20","10.132.145.21"]
	SecondaryServers []string `json:"secondary_servers,omitempty"`

	// Specifies the ISO-8601 timeout duration in which a throttled FPolicy server must complete at least one screen request. If no request is processed within the timeout, connection to the FPolicy server is terminated. The allowed range is between 0 to 100 seconds.
	// Example: PT1M
	ServerProgressTimeout *string `json:"server_progress_timeout,omitempty"`

	// Specifies the SSL option for external communication with the FPolicy server. Possible values include the following:
	// * no_auth       When set to "no_auth", no authentication takes place.
	// * server_auth   When set to "server_auth", only the FPolicy server is authenticated by the SVM. With this option, before creating the FPolicy external engine, the administrator must install the public certificate of the certificate authority (CA) that signed the FPolicy server certificate.
	// * mutual_auth   When set to "mutual_auth", mutual authentication takes place between the SVM and the FPolicy server. This means authentication of the FPolicy server by the SVM along with authentication of the SVM by the FPolicy server. With this option, before creating the FPolicy external engine, the administrator must install the public certificate of the certificate authority (CA) that signed the FPolicy server certificate along with the public certificate and key file for authentication of the SVM.
	//
	// Enum: [no_auth server_auth mutual_auth]
	SslOption *string `json:"ssl_option,omitempty"`

	// Specifies the ISO-8601 interval time for a storage appliance to query a status request from an FPolicy server. The allowed range is between 0 to 50 seconds.
	// Example: PT10S
	StatusRequestInterval *string `json:"status_request_interval,omitempty"`

	// The notification mode determines what ONTAP does after sending notifications to FPolicy servers.
	//   The possible values are:
	//     * synchronous  - After sending a notification, wait for a response from the FPolicy server.
	//     * asynchronous - After sending a notification, file request processing continues.
	//
	// Enum: [synchronous asynchronous]
	Type *string `json:"type,omitempty"`
}

// Validate validates this fpolicy engines
func (m *FpolicyEngines) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBufferSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxServerRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyEngines) validateBufferSize(formats strfmt.Registry) error {
	if swag.IsZero(m.BufferSize) { // not required
		return nil
	}

	if m.BufferSize != nil {
		if err := m.BufferSize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buffer_size")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEngines) validateCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificate) { // not required
		return nil
	}

	if m.Certificate != nil {
		if err := m.Certificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

var fpolicyEnginesTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["xml","protobuf"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyEnginesTypeFormatPropEnum = append(fpolicyEnginesTypeFormatPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// fpolicy_engines
	// FpolicyEngines
	// format
	// Format
	// xml
	// END DEBUGGING
	// FpolicyEnginesFormatXML captures enum value "xml"
	FpolicyEnginesFormatXML string = "xml"

	// BEGIN DEBUGGING
	// fpolicy_engines
	// FpolicyEngines
	// format
	// Format
	// protobuf
	// END DEBUGGING
	// FpolicyEnginesFormatProtobuf captures enum value "protobuf"
	FpolicyEnginesFormatProtobuf string = "protobuf"
)

// prop value enum
func (m *FpolicyEngines) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyEnginesTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyEngines) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", *m.Format); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyEngines) validateMaxServerRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxServerRequests) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_server_requests", "body", m.MaxServerRequests, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_server_requests", "body", m.MaxServerRequests, 10000, false); err != nil {
		return err
	}

	return nil
}

func (m *FpolicyEngines) validateResiliency(formats strfmt.Registry) error {
	if swag.IsZero(m.Resiliency) { // not required
		return nil
	}

	if m.Resiliency != nil {
		if err := m.Resiliency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resiliency")
			}
			return err
		}
	}

	return nil
}

var fpolicyEnginesTypeSslOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no_auth","server_auth","mutual_auth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyEnginesTypeSslOptionPropEnum = append(fpolicyEnginesTypeSslOptionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// fpolicy_engines
	// FpolicyEngines
	// ssl_option
	// SslOption
	// no_auth
	// END DEBUGGING
	// FpolicyEnginesSslOptionNoAuth captures enum value "no_auth"
	FpolicyEnginesSslOptionNoAuth string = "no_auth"

	// BEGIN DEBUGGING
	// fpolicy_engines
	// FpolicyEngines
	// ssl_option
	// SslOption
	// server_auth
	// END DEBUGGING
	// FpolicyEnginesSslOptionServerAuth captures enum value "server_auth"
	FpolicyEnginesSslOptionServerAuth string = "server_auth"

	// BEGIN DEBUGGING
	// fpolicy_engines
	// FpolicyEngines
	// ssl_option
	// SslOption
	// mutual_auth
	// END DEBUGGING
	// FpolicyEnginesSslOptionMutualAuth captures enum value "mutual_auth"
	FpolicyEnginesSslOptionMutualAuth string = "mutual_auth"
)

// prop value enum
func (m *FpolicyEngines) validateSslOptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyEnginesTypeSslOptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyEngines) validateSslOption(formats strfmt.Registry) error {
	if swag.IsZero(m.SslOption) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslOptionEnum("ssl_option", "body", *m.SslOption); err != nil {
		return err
	}

	return nil
}

var fpolicyEnginesTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["synchronous","asynchronous"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fpolicyEnginesTypeTypePropEnum = append(fpolicyEnginesTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// fpolicy_engines
	// FpolicyEngines
	// type
	// Type
	// synchronous
	// END DEBUGGING
	// FpolicyEnginesTypeSynchronous captures enum value "synchronous"
	FpolicyEnginesTypeSynchronous string = "synchronous"

	// BEGIN DEBUGGING
	// fpolicy_engines
	// FpolicyEngines
	// type
	// Type
	// asynchronous
	// END DEBUGGING
	// FpolicyEnginesTypeAsynchronous captures enum value "asynchronous"
	FpolicyEnginesTypeAsynchronous string = "asynchronous"
)

// prop value enum
func (m *FpolicyEngines) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fpolicyEnginesTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FpolicyEngines) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fpolicy engines based on the context it is used
func (m *FpolicyEngines) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBufferSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCertificate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResiliency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FpolicyEngines) contextValidateBufferSize(ctx context.Context, formats strfmt.Registry) error {

	if m.BufferSize != nil {
		if err := m.BufferSize.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buffer_size")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEngines) contextValidateCertificate(ctx context.Context, formats strfmt.Registry) error {

	if m.Certificate != nil {
		if err := m.Certificate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *FpolicyEngines) contextValidateResiliency(ctx context.Context, formats strfmt.Registry) error {

	if m.Resiliency != nil {
		if err := m.Resiliency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resiliency")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEngines) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEngines) UnmarshalBinary(b []byte) error {
	var res FpolicyEngines
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEnginesBufferSize Specifies the send and recieve buffer size of the connected socket for the FPolicy server.
//
// swagger:model FpolicyEnginesBufferSize
type FpolicyEnginesBufferSize struct {

	// Specifies the receive buffer size of the connected socket for the FPolicy server. Default value is 256KB.
	RecvBuffer *int64 `json:"recv_buffer,omitempty"`

	// Specifies the send buffer size of the connected socket for the FPolicy server. Default value 1MB.
	SendBuffer *int64 `json:"send_buffer,omitempty"`
}

// Validate validates this fpolicy engines buffer size
func (m *FpolicyEnginesBufferSize) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy engines buffer size based on context it is used
func (m *FpolicyEnginesBufferSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEnginesBufferSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEnginesBufferSize) UnmarshalBinary(b []byte) error {
	var res FpolicyEnginesBufferSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEnginesCertificate Provides details about certificate used to authenticate the Fpolicy server.
//
// swagger:model FpolicyEnginesCertificate
type FpolicyEnginesCertificate struct {

	// Specifies the certificate authority (CA) name of the certificate
	// used for authentication if SSL authentication between the SVM and the FPolicy
	// server is configured.
	//
	// Example: TASample1
	Ca string `json:"ca,omitempty"`

	// Specifies the certificate name as a fully qualified domain
	// name (FQDN) or custom common name. The certificate is used if SSL authentication
	// between the SVM and the FPolicy server is configured.
	//
	// Example: Sample1-FPolicy-Client
	Name string `json:"name,omitempty"`

	// Specifies the serial number of the certificate used for
	// authentication if SSL authentication between the SVM and the FPolicy
	// server is configured.
	//
	// Example: 8DDE112A114D1FBC
	SerialNumber string `json:"serial_number,omitempty"`
}

// Validate validates this fpolicy engines certificate
func (m *FpolicyEnginesCertificate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy engines certificate based on context it is used
func (m *FpolicyEnginesCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEnginesCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEnginesCertificate) UnmarshalBinary(b []byte) error {
	var res FpolicyEnginesCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FpolicyEnginesResiliency If all primary and secondary servers are down, or if no response is received from the FPolicy servers, file access events are stored inside the storage controller under the specified resiliency-directory-path.
//
// swagger:model FpolicyEnginesResiliency
type FpolicyEnginesResiliency struct {

	// Specifies the directory path under the SVM namespace,
	// where notifications are stored in the files whenever a network outage happens.
	//
	// Example: /dir1
	DirectoryPath string `json:"directory_path,omitempty"`

	// Specifies whether the resiliency feature is enabled or not.
	// Default is false.
	//
	Enabled *bool `json:"enabled,omitempty"`

	// Specifies the ISO-8601 duration, for which the notifications are written
	// to files inside the storage controller during a network outage. The value for
	// this field must be between 0 and 600 seconds. Default is 180 seconds.
	//
	// Example: PT3M
	RetentionDuration string `json:"retention_duration,omitempty"`
}

// Validate validates this fpolicy engines resiliency
func (m *FpolicyEnginesResiliency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fpolicy engines resiliency based on context it is used
func (m *FpolicyEnginesResiliency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FpolicyEnginesResiliency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FpolicyEnginesResiliency) UnmarshalBinary(b []byte) error {
	var res FpolicyEnginesResiliency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
