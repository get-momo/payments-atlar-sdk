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

// PostV1betaExpectedTransactionsIDArchiveReader is a Reader for the PostV1betaExpectedTransactionsIDArchive structure.
type PostV1betaExpectedTransactionsIDArchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1betaExpectedTransactionsIDArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1betaExpectedTransactionsIDArchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1beta/expected-transactions/{id}:archive] PostV1betaExpectedTransactionsIDArchive", response, response.Code())
	}
}

// NewPostV1betaExpectedTransactionsIDArchiveOK creates a PostV1betaExpectedTransactionsIDArchiveOK with default headers values
func NewPostV1betaExpectedTransactionsIDArchiveOK() *PostV1betaExpectedTransactionsIDArchiveOK {
	return &PostV1betaExpectedTransactionsIDArchiveOK{}
}

/*
PostV1betaExpectedTransactionsIDArchiveOK describes a response with status code 200, with default header values.

The Expected Transaction
*/
type PostV1betaExpectedTransactionsIDArchiveOK struct {
	Payload *models.ExpectedTransaction
}

// IsSuccess returns true when this post v1beta expected transactions Id archive o k response has a 2xx status code
func (o *PostV1betaExpectedTransactionsIDArchiveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1beta expected transactions Id archive o k response has a 3xx status code
func (o *PostV1betaExpectedTransactionsIDArchiveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1beta expected transactions Id archive o k response has a 4xx status code
func (o *PostV1betaExpectedTransactionsIDArchiveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1beta expected transactions Id archive o k response has a 5xx status code
func (o *PostV1betaExpectedTransactionsIDArchiveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1beta expected transactions Id archive o k response a status code equal to that given
func (o *PostV1betaExpectedTransactionsIDArchiveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1beta expected transactions Id archive o k response
func (o *PostV1betaExpectedTransactionsIDArchiveOK) Code() int {
	return 200
}

func (o *PostV1betaExpectedTransactionsIDArchiveOK) Error() string {
	return fmt.Sprintf("[POST /v1beta/expected-transactions/{id}:archive][%d] postV1betaExpectedTransactionsIdArchiveOK  %+v", 200, o.Payload)
}

func (o *PostV1betaExpectedTransactionsIDArchiveOK) String() string {
	return fmt.Sprintf("[POST /v1beta/expected-transactions/{id}:archive][%d] postV1betaExpectedTransactionsIdArchiveOK  %+v", 200, o.Payload)
}

func (o *PostV1betaExpectedTransactionsIDArchiveOK) GetPayload() *models.ExpectedTransaction {
	return o.Payload
}

func (o *PostV1betaExpectedTransactionsIDArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpectedTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
