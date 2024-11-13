// Code generated by go-swagger; DO NOT EDIT.

package expected_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// PostV1betaExpectedTransactionsIDUnarchiveReader is a Reader for the PostV1betaExpectedTransactionsIDUnarchive structure.
type PostV1betaExpectedTransactionsIDUnarchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1betaExpectedTransactionsIDUnarchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1betaExpectedTransactionsIDUnarchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1beta/expected-transactions/{id}:unarchive] PostV1betaExpectedTransactionsIDUnarchive", response, response.Code())
	}
}

// NewPostV1betaExpectedTransactionsIDUnarchiveOK creates a PostV1betaExpectedTransactionsIDUnarchiveOK with default headers values
func NewPostV1betaExpectedTransactionsIDUnarchiveOK() *PostV1betaExpectedTransactionsIDUnarchiveOK {
	return &PostV1betaExpectedTransactionsIDUnarchiveOK{}
}

/*
PostV1betaExpectedTransactionsIDUnarchiveOK describes a response with status code 200, with default header values.

The Expected Transaction
*/
type PostV1betaExpectedTransactionsIDUnarchiveOK struct {
	Payload *models.ExpectedTransaction
}

// IsSuccess returns true when this post v1beta expected transactions Id unarchive o k response has a 2xx status code
func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1beta expected transactions Id unarchive o k response has a 3xx status code
func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1beta expected transactions Id unarchive o k response has a 4xx status code
func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1beta expected transactions Id unarchive o k response has a 5xx status code
func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1beta expected transactions Id unarchive o k response a status code equal to that given
func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1beta expected transactions Id unarchive o k response
func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) Code() int {
	return 200
}

func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) Error() string {
	return fmt.Sprintf("[POST /v1beta/expected-transactions/{id}:unarchive][%d] postV1betaExpectedTransactionsIdUnarchiveOK  %+v", 200, o.Payload)
}

func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) String() string {
	return fmt.Sprintf("[POST /v1beta/expected-transactions/{id}:unarchive][%d] postV1betaExpectedTransactionsIdUnarchiveOK  %+v", 200, o.Payload)
}

func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) GetPayload() *models.ExpectedTransaction {
	return o.Payload
}

func (o *PostV1betaExpectedTransactionsIDUnarchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpectedTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
