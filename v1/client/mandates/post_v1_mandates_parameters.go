// Code generated by go-swagger; DO NOT EDIT.

package mandates

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

// NewPostV1MandatesParams creates a new PostV1MandatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1MandatesParams() *PostV1MandatesParams {
	return &PostV1MandatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1MandatesParamsWithTimeout creates a new PostV1MandatesParams object
// with the ability to set a timeout on a request.
func NewPostV1MandatesParamsWithTimeout(timeout time.Duration) *PostV1MandatesParams {
	return &PostV1MandatesParams{
		timeout: timeout,
	}
}

// NewPostV1MandatesParamsWithContext creates a new PostV1MandatesParams object
// with the ability to set a context for a request.
func NewPostV1MandatesParamsWithContext(ctx context.Context) *PostV1MandatesParams {
	return &PostV1MandatesParams{
		Context: ctx,
	}
}

// NewPostV1MandatesParamsWithHTTPClient creates a new PostV1MandatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1MandatesParamsWithHTTPClient(client *http.Client) *PostV1MandatesParams {
	return &PostV1MandatesParams{
		HTTPClient: client,
	}
}

/*
PostV1MandatesParams contains all the parameters to send to the API endpoint

	for the post v1 mandates operation.

	Typically these are written to a http.Request.
*/
type PostV1MandatesParams struct {

	/* Mandate.

	   Mandate
	*/
	Mandate *models.CreateMandateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 mandates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1MandatesParams) WithDefaults() *PostV1MandatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 mandates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1MandatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 mandates params
func (o *PostV1MandatesParams) WithTimeout(timeout time.Duration) *PostV1MandatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 mandates params
func (o *PostV1MandatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 mandates params
func (o *PostV1MandatesParams) WithContext(ctx context.Context) *PostV1MandatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 mandates params
func (o *PostV1MandatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 mandates params
func (o *PostV1MandatesParams) WithHTTPClient(client *http.Client) *PostV1MandatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 mandates params
func (o *PostV1MandatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMandate adds the mandate to the post v1 mandates params
func (o *PostV1MandatesParams) WithMandate(mandate *models.CreateMandateRequest) *PostV1MandatesParams {
	o.SetMandate(mandate)
	return o
}

// SetMandate adds the mandate to the post v1 mandates params
func (o *PostV1MandatesParams) SetMandate(mandate *models.CreateMandateRequest) {
	o.Mandate = mandate
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1MandatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Mandate != nil {
		if err := r.SetBodyParam(o.Mandate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}