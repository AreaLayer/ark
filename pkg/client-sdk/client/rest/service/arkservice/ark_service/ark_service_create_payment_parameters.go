// Code generated by go-swagger; DO NOT EDIT.

package ark_service

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

	"github.com/ark-network/ark/pkg/client-sdk/client/rest/service/models"
)

// NewArkServiceCreatePaymentParams creates a new ArkServiceCreatePaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewArkServiceCreatePaymentParams() *ArkServiceCreatePaymentParams {
	return &ArkServiceCreatePaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewArkServiceCreatePaymentParamsWithTimeout creates a new ArkServiceCreatePaymentParams object
// with the ability to set a timeout on a request.
func NewArkServiceCreatePaymentParamsWithTimeout(timeout time.Duration) *ArkServiceCreatePaymentParams {
	return &ArkServiceCreatePaymentParams{
		timeout: timeout,
	}
}

// NewArkServiceCreatePaymentParamsWithContext creates a new ArkServiceCreatePaymentParams object
// with the ability to set a context for a request.
func NewArkServiceCreatePaymentParamsWithContext(ctx context.Context) *ArkServiceCreatePaymentParams {
	return &ArkServiceCreatePaymentParams{
		Context: ctx,
	}
}

// NewArkServiceCreatePaymentParamsWithHTTPClient creates a new ArkServiceCreatePaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewArkServiceCreatePaymentParamsWithHTTPClient(client *http.Client) *ArkServiceCreatePaymentParams {
	return &ArkServiceCreatePaymentParams{
		HTTPClient: client,
	}
}

/*
ArkServiceCreatePaymentParams contains all the parameters to send to the API endpoint

	for the ark service create payment operation.

	Typically these are written to a http.Request.
*/
type ArkServiceCreatePaymentParams struct {

	// Body.
	Body *models.V1CreatePaymentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ark service create payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArkServiceCreatePaymentParams) WithDefaults() *ArkServiceCreatePaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ark service create payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArkServiceCreatePaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ark service create payment params
func (o *ArkServiceCreatePaymentParams) WithTimeout(timeout time.Duration) *ArkServiceCreatePaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ark service create payment params
func (o *ArkServiceCreatePaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ark service create payment params
func (o *ArkServiceCreatePaymentParams) WithContext(ctx context.Context) *ArkServiceCreatePaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ark service create payment params
func (o *ArkServiceCreatePaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ark service create payment params
func (o *ArkServiceCreatePaymentParams) WithHTTPClient(client *http.Client) *ArkServiceCreatePaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ark service create payment params
func (o *ArkServiceCreatePaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ark service create payment params
func (o *ArkServiceCreatePaymentParams) WithBody(body *models.V1CreatePaymentRequest) *ArkServiceCreatePaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ark service create payment params
func (o *ArkServiceCreatePaymentParams) SetBody(body *models.V1CreatePaymentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ArkServiceCreatePaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
