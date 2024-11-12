// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeprecatedCreditTransferScheme Deprecated. Use `paymentSchemeType` instead.
//
// swagger:model DeprecatedCreditTransferScheme
type DeprecatedCreditTransferScheme struct {

	// display name
	// Example: Sepa Credit Transfer
	DisplayName string `json:"displayName,omitempty"`

	// type
	// Example: SCT
	Type string `json:"type,omitempty"`
}

// Validate validates this deprecated credit transfer scheme
func (m *DeprecatedCreditTransferScheme) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this deprecated credit transfer scheme based on context it is used
func (m *DeprecatedCreditTransferScheme) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeprecatedCreditTransferScheme) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeprecatedCreditTransferScheme) UnmarshalBinary(b []byte) error {
	var res DeprecatedCreditTransferScheme
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}