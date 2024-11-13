// Code generated by go-swagger; DO NOT EDIT.

package credit_transfers

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

// GetV1CreditTransfersReader is a Reader for the GetV1CreditTransfers structure.
type GetV1CreditTransfersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1CreditTransfersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1CreditTransfersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1CreditTransfersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/credit-transfers] GetV1CreditTransfers", response, response.Code())
	}
}

// NewGetV1CreditTransfersOK creates a GetV1CreditTransfersOK with default headers values
func NewGetV1CreditTransfersOK() *GetV1CreditTransfersOK {
	return &GetV1CreditTransfersOK{}
}

/*
GetV1CreditTransfersOK describes a response with status code 200, with default header values.

QueryResponse with list of CreditTransfers
*/
type GetV1CreditTransfersOK struct {
	Payload *GetV1CreditTransfersOKBody
}

// IsSuccess returns true when this get v1 credit transfers o k response has a 2xx status code
func (o *GetV1CreditTransfersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 credit transfers o k response has a 3xx status code
func (o *GetV1CreditTransfersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 credit transfers o k response has a 4xx status code
func (o *GetV1CreditTransfersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 credit transfers o k response has a 5xx status code
func (o *GetV1CreditTransfersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 credit transfers o k response a status code equal to that given
func (o *GetV1CreditTransfersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 credit transfers o k response
func (o *GetV1CreditTransfersOK) Code() int {
	return 200
}

func (o *GetV1CreditTransfersOK) Error() string {
	return fmt.Sprintf("[GET /v1/credit-transfers][%d] getV1CreditTransfersOK  %+v", 200, o.Payload)
}

func (o *GetV1CreditTransfersOK) String() string {
	return fmt.Sprintf("[GET /v1/credit-transfers][%d] getV1CreditTransfersOK  %+v", 200, o.Payload)
}

func (o *GetV1CreditTransfersOK) GetPayload() *GetV1CreditTransfersOKBody {
	return o.Payload
}

func (o *GetV1CreditTransfersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1CreditTransfersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1CreditTransfersBadRequest creates a GetV1CreditTransfersBadRequest with default headers values
func NewGetV1CreditTransfersBadRequest() *GetV1CreditTransfersBadRequest {
	return &GetV1CreditTransfersBadRequest{}
}

/*
GetV1CreditTransfersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetV1CreditTransfersBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 credit transfers bad request response has a 2xx status code
func (o *GetV1CreditTransfersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 credit transfers bad request response has a 3xx status code
func (o *GetV1CreditTransfersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 credit transfers bad request response has a 4xx status code
func (o *GetV1CreditTransfersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 credit transfers bad request response has a 5xx status code
func (o *GetV1CreditTransfersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 credit transfers bad request response a status code equal to that given
func (o *GetV1CreditTransfersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 credit transfers bad request response
func (o *GetV1CreditTransfersBadRequest) Code() int {
	return 400
}

func (o *GetV1CreditTransfersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/credit-transfers][%d] getV1CreditTransfersBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1CreditTransfersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/credit-transfers][%d] getV1CreditTransfersBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1CreditTransfersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1CreditTransfersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1CreditTransfersOKBody get v1 credit transfers o k body
swagger:model GetV1CreditTransfersOKBody
*/
type GetV1CreditTransfersOKBody struct {
	models.QueryResponse

	// items
	Items []*models.Payment `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1CreditTransfersOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1CreditTransfersOKBodyAO0
	var getV1CreditTransfersOKBodyAO0 models.QueryResponse
	if err := swag.ReadJSON(raw, &getV1CreditTransfersOKBodyAO0); err != nil {
		return err
	}
	o.QueryResponse = getV1CreditTransfersOKBodyAO0

	// GetV1CreditTransfersOKBodyAO1
	var dataGetV1CreditTransfersOKBodyAO1 struct {
		Items []*models.Payment `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1CreditTransfersOKBodyAO1); err != nil {
		return err
	}

	o.Items = dataGetV1CreditTransfersOKBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1CreditTransfersOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1CreditTransfersOKBodyAO0, err := swag.WriteJSON(o.QueryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1CreditTransfersOKBodyAO0)
	var dataGetV1CreditTransfersOKBodyAO1 struct {
		Items []*models.Payment `json:"items"`
	}

	dataGetV1CreditTransfersOKBodyAO1.Items = o.Items

	jsonDataGetV1CreditTransfersOKBodyAO1, errGetV1CreditTransfersOKBodyAO1 := swag.WriteJSON(dataGetV1CreditTransfersOKBodyAO1)
	if errGetV1CreditTransfersOKBodyAO1 != nil {
		return nil, errGetV1CreditTransfersOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1CreditTransfersOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 credit transfers o k body
func (o *GetV1CreditTransfersOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetV1CreditTransfersOKBody) validateItems(formats strfmt.Registry) error {

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
					return ve.ValidateName("getV1CreditTransfersOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1CreditTransfersOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 credit transfers o k body based on the context it is used
func (o *GetV1CreditTransfersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetV1CreditTransfersOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1CreditTransfersOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1CreditTransfersOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1CreditTransfersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1CreditTransfersOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1CreditTransfersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
