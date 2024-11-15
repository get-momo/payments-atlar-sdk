// Code generated by go-swagger; DO NOT EDIT.

package counterparties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// PutV1CounterpartiesIDReader is a Reader for the PutV1CounterpartiesID structure.
type PutV1CounterpartiesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1CounterpartiesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1CounterpartiesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutV1CounterpartiesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutV1CounterpartiesIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/counterparties/{id}] PutV1CounterpartiesID", response, response.Code())
	}
}

// NewPutV1CounterpartiesIDOK creates a PutV1CounterpartiesIDOK with default headers values
func NewPutV1CounterpartiesIDOK() *PutV1CounterpartiesIDOK {
	return &PutV1CounterpartiesIDOK{}
}

/*
PutV1CounterpartiesIDOK describes a response with status code 200, with default header values.

The updated counterparty
*/
type PutV1CounterpartiesIDOK struct {
	Payload *models.Counterparty
}

// IsSuccess returns true when this put v1 counterparties Id o k response has a 2xx status code
func (o *PutV1CounterpartiesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 counterparties Id o k response has a 3xx status code
func (o *PutV1CounterpartiesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 counterparties Id o k response has a 4xx status code
func (o *PutV1CounterpartiesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 counterparties Id o k response has a 5xx status code
func (o *PutV1CounterpartiesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 counterparties Id o k response a status code equal to that given
func (o *PutV1CounterpartiesIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put v1 counterparties Id o k response
func (o *PutV1CounterpartiesIDOK) Code() int {
	return 200
}

func (o *PutV1CounterpartiesIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/counterparties/{id}][%d] putV1CounterpartiesIdOK  %+v", 200, o.Payload)
}

func (o *PutV1CounterpartiesIDOK) String() string {
	return fmt.Sprintf("[PUT /v1/counterparties/{id}][%d] putV1CounterpartiesIdOK  %+v", 200, o.Payload)
}

func (o *PutV1CounterpartiesIDOK) GetPayload() *models.Counterparty {
	return o.Payload
}

func (o *PutV1CounterpartiesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Counterparty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1CounterpartiesIDBadRequest creates a PutV1CounterpartiesIDBadRequest with default headers values
func NewPutV1CounterpartiesIDBadRequest() *PutV1CounterpartiesIDBadRequest {
	return &PutV1CounterpartiesIDBadRequest{}
}

/*
PutV1CounterpartiesIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutV1CounterpartiesIDBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this put v1 counterparties Id bad request response has a 2xx status code
func (o *PutV1CounterpartiesIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 counterparties Id bad request response has a 3xx status code
func (o *PutV1CounterpartiesIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 counterparties Id bad request response has a 4xx status code
func (o *PutV1CounterpartiesIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 counterparties Id bad request response has a 5xx status code
func (o *PutV1CounterpartiesIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 counterparties Id bad request response a status code equal to that given
func (o *PutV1CounterpartiesIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put v1 counterparties Id bad request response
func (o *PutV1CounterpartiesIDBadRequest) Code() int {
	return 400
}

func (o *PutV1CounterpartiesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/counterparties/{id}][%d] putV1CounterpartiesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1CounterpartiesIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/counterparties/{id}][%d] putV1CounterpartiesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1CounterpartiesIDBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutV1CounterpartiesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1CounterpartiesIDConflict creates a PutV1CounterpartiesIDConflict with default headers values
func NewPutV1CounterpartiesIDConflict() *PutV1CounterpartiesIDConflict {
	return &PutV1CounterpartiesIDConflict{}
}

/*
PutV1CounterpartiesIDConflict describes a response with status code 409, with default header values.

Conflict with the current state of the target resource
*/
type PutV1CounterpartiesIDConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this put v1 counterparties Id conflict response has a 2xx status code
func (o *PutV1CounterpartiesIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 counterparties Id conflict response has a 3xx status code
func (o *PutV1CounterpartiesIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 counterparties Id conflict response has a 4xx status code
func (o *PutV1CounterpartiesIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 counterparties Id conflict response has a 5xx status code
func (o *PutV1CounterpartiesIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 counterparties Id conflict response a status code equal to that given
func (o *PutV1CounterpartiesIDConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the put v1 counterparties Id conflict response
func (o *PutV1CounterpartiesIDConflict) Code() int {
	return 409
}

func (o *PutV1CounterpartiesIDConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/counterparties/{id}][%d] putV1CounterpartiesIdConflict  %+v", 409, o.Payload)
}

func (o *PutV1CounterpartiesIDConflict) String() string {
	return fmt.Sprintf("[PUT /v1/counterparties/{id}][%d] putV1CounterpartiesIdConflict  %+v", 409, o.Payload)
}

func (o *PutV1CounterpartiesIDConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutV1CounterpartiesIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
