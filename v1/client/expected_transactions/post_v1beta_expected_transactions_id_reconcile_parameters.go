// Code generated by go-swagger; DO NOT EDIT.

package expected_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// NewPostV1betaExpectedTransactionsIDReconcileParams creates a new PostV1betaExpectedTransactionsIDReconcileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1betaExpectedTransactionsIDReconcileParams() *PostV1betaExpectedTransactionsIDReconcileParams {
	return &PostV1betaExpectedTransactionsIDReconcileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1betaExpectedTransactionsIDReconcileParamsWithTimeout creates a new PostV1betaExpectedTransactionsIDReconcileParams object
// with the ability to set a timeout on a request.
func NewPostV1betaExpectedTransactionsIDReconcileParamsWithTimeout(timeout time.Duration) *PostV1betaExpectedTransactionsIDReconcileParams {
	return &PostV1betaExpectedTransactionsIDReconcileParams{
		timeout: timeout,
	}
}

// NewPostV1betaExpectedTransactionsIDReconcileParamsWithContext creates a new PostV1betaExpectedTransactionsIDReconcileParams object
// with the ability to set a context for a request.
func NewPostV1betaExpectedTransactionsIDReconcileParamsWithContext(ctx context.Context) *PostV1betaExpectedTransactionsIDReconcileParams {
	return &PostV1betaExpectedTransactionsIDReconcileParams{
		Context: ctx,
	}
}

// NewPostV1betaExpectedTransactionsIDReconcileParamsWithHTTPClient creates a new PostV1betaExpectedTransactionsIDReconcileParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1betaExpectedTransactionsIDReconcileParamsWithHTTPClient(client *http.Client) *PostV1betaExpectedTransactionsIDReconcileParams {
	return &PostV1betaExpectedTransactionsIDReconcileParams{
		HTTPClient: client,
	}
}

/*
PostV1betaExpectedTransactionsIDReconcileParams contains all the parameters to send to the API endpoint

	for the post v1beta expected transactions ID reconcile operation.

	Typically these are written to a http.Request.
*/
type PostV1betaExpectedTransactionsIDReconcileParams struct {

	/* ID.

	   Expected Transaction ID
	*/
	ID string

	/* Transaction.

	   Expected Transaction
	*/
	Transaction *models.ReconcileExpectedTransactionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1beta expected transactions ID reconcile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1betaExpectedTransactionsIDReconcileParams) WithDefaults() *PostV1betaExpectedTransactionsIDReconcileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1beta expected transactions ID reconcile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1betaExpectedTransactionsIDReconcileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) WithTimeout(timeout time.Duration) *PostV1betaExpectedTransactionsIDReconcileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) WithContext(ctx context.Context) *PostV1betaExpectedTransactionsIDReconcileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) WithHTTPClient(client *http.Client) *PostV1betaExpectedTransactionsIDReconcileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) WithID(id string) *PostV1betaExpectedTransactionsIDReconcileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) SetID(id string) {
	o.ID = id
}

// WithTransaction adds the transaction to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) WithTransaction(transaction *models.ReconcileExpectedTransactionRequest) *PostV1betaExpectedTransactionsIDReconcileParams {
	o.SetTransaction(transaction)
	return o
}

// SetTransaction adds the transaction to the post v1beta expected transactions ID reconcile params
func (o *PostV1betaExpectedTransactionsIDReconcileParams) SetTransaction(transaction *models.ReconcileExpectedTransactionRequest) {
	o.Transaction = transaction
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1betaExpectedTransactionsIDReconcileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Transaction != nil {
		if err := r.SetBodyParam(o.Transaction); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
