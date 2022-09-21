// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LdapSchemaAccount ldap schema account
//
// swagger:model ldap_schema_account
type LdapSchemaAccount struct {

	// Attribute name used to retrieve UNIX account information.
	// Example: windowsAccount
	Unix string `json:"unix,omitempty"`

	// Attribute name used to retrieve Windows account information for a UNIX user account.
	// Example: windowsAccount
	Windows string `json:"windows,omitempty"`
}

// Validate validates this ldap schema account
func (m *LdapSchemaAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ldap schema account based on context it is used
func (m *LdapSchemaAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapSchemaAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapSchemaAccount) UnmarshalBinary(b []byte) error {
	var res LdapSchemaAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
