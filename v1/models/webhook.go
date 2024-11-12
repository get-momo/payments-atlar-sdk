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

// Webhook webhook
//
// swagger:model Webhook
type Webhook struct {

	// The time this resource was created.
	// Example: 2022-07-20T18:31:12.889104898Z
	Created string `json:"created,omitempty"`

	// An inclusive filter to specify which resources and events that you are interested in.
	Filter []*WebhookFilter `json:"filter"`

	// The unique identifier of this webhook.
	// Example: be7e04fe-3fa0-48f5-a993-2008409255f2
	ID string `json:"id,omitempty"`

	// An optional name for your webhook where you could describe its purpose.
	// Example: Payment updates
	Name string `json:"name,omitempty"`

	// Your organization ID.
	// Example: 605e26fc-4fce-495a-a92f-2c3592d7287e
	OrganizationID string `json:"organizationId,omitempty"`

	// The last time this resource was updated.
	// Example: 2022-05-20T18:31:12.889104898Z
	Updated string `json:"updated,omitempty"`

	// The URL to which webhooks will be sent via HTTP POST requests in JSON format. Must use https.
	// Example: https://example.com
	// Required: true
	URL *string `json:"url"`

	// Webhooks won't be sent if the configuration is not verified. Modifying the URL sets verified to false. You can contact [support@atlar.com](mailto:support@atlar.com) to get a webhook configuration verified.
	// Example: true
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this webhook
func (m *Webhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Webhook) validateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	for i := 0; i < len(m.Filter); i++ {
		if swag.IsZero(m.Filter[i]) { // not required
			continue
		}

		if m.Filter[i] != nil {
			if err := m.Filter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filter" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Webhook) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this webhook based on the context it is used
func (m *Webhook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Webhook) contextValidateFilter(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filter); i++ {

		if m.Filter[i] != nil {

			if swag.IsZero(m.Filter[i]) { // not required
				return nil
			}

			if err := m.Filter[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filter" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Webhook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Webhook) UnmarshalBinary(b []byte) error {
	var res Webhook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}