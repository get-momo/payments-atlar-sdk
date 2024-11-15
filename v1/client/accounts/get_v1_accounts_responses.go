// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// GetV1AccountsReader is a Reader for the GetV1Accounts structure.
type GetV1AccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1AccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1AccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/accounts] GetV1Accounts", response, response.Code())
	}
}

// NewGetV1AccountsOK creates a GetV1AccountsOK with default headers values
func NewGetV1AccountsOK() *GetV1AccountsOK {
	return &GetV1AccountsOK{}
}

/*
GetV1AccountsOK describes a response with status code 200, with default header values.

QueryResponse with list of Accounts
*/
type GetV1AccountsOK struct {
	Payload *GetV1AccountsOKBody
}

// IsSuccess returns true when this get v1 accounts o k response has a 2xx status code
func (o *GetV1AccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 accounts o k response has a 3xx status code
func (o *GetV1AccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 accounts o k response has a 4xx status code
func (o *GetV1AccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 accounts o k response has a 5xx status code
func (o *GetV1AccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 accounts o k response a status code equal to that given
func (o *GetV1AccountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 accounts o k response
func (o *GetV1AccountsOK) Code() int {
	return 200
}

func (o *GetV1AccountsOK) Error() string {
	return fmt.Sprintf("[GET /v1/accounts][%d] getV1AccountsOK  %+v", 200, o.Payload)
}

func (o *GetV1AccountsOK) String() string {
	return fmt.Sprintf("[GET /v1/accounts][%d] getV1AccountsOK  %+v", 200, o.Payload)
}

func (o *GetV1AccountsOK) GetPayload() *GetV1AccountsOKBody {
	return o.Payload
}

func (o *GetV1AccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1AccountsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1AccountsOKBody get v1 accounts o k body
swagger:model GetV1AccountsOKBody
*/
type GetV1AccountsOKBody struct {
	models.QueryResponse

	// items
	Items []*models.Account `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1AccountsOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1AccountsOKBodyAO0
	var getV1AccountsOKBodyAO0 models.QueryResponse
	if err := swag.ReadJSON(raw, &getV1AccountsOKBodyAO0); err != nil {
		return err
	}
	o.QueryResponse = getV1AccountsOKBodyAO0

	// GetV1AccountsOKBodyAO1
	var dataGetV1AccountsOKBodyAO1 struct {
		Items []*models.Account `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1AccountsOKBodyAO1); err != nil {
		return err
	}

	o.Items = dataGetV1AccountsOKBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1AccountsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1AccountsOKBodyAO0, err := swag.WriteJSON(o.QueryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1AccountsOKBodyAO0)
	var dataGetV1AccountsOKBodyAO1 struct {
		Items []*models.Account `json:"items"`
	}

	dataGetV1AccountsOKBodyAO1.Items = o.Items

	jsonDataGetV1AccountsOKBodyAO1, errGetV1AccountsOKBodyAO1 := swag.WriteJSON(dataGetV1AccountsOKBodyAO1)
	if errGetV1AccountsOKBodyAO1 != nil {
		return nil, errGetV1AccountsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1AccountsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 accounts o k body
func (o *GetV1AccountsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.QueryResponse
	if err := o.QueryResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV1AccountsOKBody) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1AccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1AccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 accounts o k body based on the context it is used
func (o *GetV1AccountsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.QueryResponse
	if err := o.QueryResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV1AccountsOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1AccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1AccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1AccountsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1AccountsOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1AccountsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
