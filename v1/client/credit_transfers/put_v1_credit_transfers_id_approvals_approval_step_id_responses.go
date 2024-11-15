// Code generated by go-swagger; DO NOT EDIT.

package credit_transfers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// PutV1CreditTransfersIDApprovalsApprovalStepIDReader is a Reader for the PutV1CreditTransfersIDApprovalsApprovalStepID structure.
type PutV1CreditTransfersIDApprovalsApprovalStepIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPutV1CreditTransfersIDApprovalsApprovalStepIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/credit-transfers/{id}/approvals/{approvalStepId}] PutV1CreditTransfersIDApprovalsApprovalStepID", response, response.Code())
	}
}

// NewPutV1CreditTransfersIDApprovalsApprovalStepIDAccepted creates a PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted with default headers values
func NewPutV1CreditTransfersIDApprovalsApprovalStepIDAccepted() *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted {
	return &PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted{}
}

/*
PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted describes a response with status code 202, with default header values.

The, now, approved identified CreditTransfer
*/
type PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted struct {
	Payload *models.Payment
}

// IsSuccess returns true when this put v1 credit transfers Id approvals approval step Id accepted response has a 2xx status code
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 credit transfers Id approvals approval step Id accepted response has a 3xx status code
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 credit transfers Id approvals approval step Id accepted response has a 4xx status code
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 credit transfers Id approvals approval step Id accepted response has a 5xx status code
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 credit transfers Id approvals approval step Id accepted response a status code equal to that given
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the put v1 credit transfers Id approvals approval step Id accepted response
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) Code() int {
	return 202
}

func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/credit-transfers/{id}/approvals/{approvalStepId}][%d] putV1CreditTransfersIdApprovalsApprovalStepIdAccepted  %+v", 202, o.Payload)
}

func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/credit-transfers/{id}/approvals/{approvalStepId}][%d] putV1CreditTransfersIdApprovalsApprovalStepIdAccepted  %+v", 202, o.Payload)
}

func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) GetPayload() *models.Payment {
	return o.Payload
}

func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest creates a PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest with default headers values
func NewPutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest() *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest {
	return &PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest{}
}

/*
PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this put v1 credit transfers Id approvals approval step Id bad request response has a 2xx status code
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 credit transfers Id approvals approval step Id bad request response has a 3xx status code
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 credit transfers Id approvals approval step Id bad request response has a 4xx status code
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 credit transfers Id approvals approval step Id bad request response has a 5xx status code
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 credit transfers Id approvals approval step Id bad request response a status code equal to that given
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put v1 credit transfers Id approvals approval step Id bad request response
func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) Code() int {
	return 400
}

func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/credit-transfers/{id}/approvals/{approvalStepId}][%d] putV1CreditTransfersIdApprovalsApprovalStepIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/credit-transfers/{id}/approvals/{approvalStepId}][%d] putV1CreditTransfersIdApprovalsApprovalStepIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutV1CreditTransfersIDApprovalsApprovalStepIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
