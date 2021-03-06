// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateActionsDeleteParams creates a new WeaviateActionsDeleteParams object
// with the default values initialized.
func NewWeaviateActionsDeleteParams() *WeaviateActionsDeleteParams {
	var ()
	return &WeaviateActionsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionsDeleteParamsWithTimeout creates a new WeaviateActionsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionsDeleteParamsWithTimeout(timeout time.Duration) *WeaviateActionsDeleteParams {
	var ()
	return &WeaviateActionsDeleteParams{

		timeout: timeout,
	}
}

// NewWeaviateActionsDeleteParamsWithContext creates a new WeaviateActionsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionsDeleteParamsWithContext(ctx context.Context) *WeaviateActionsDeleteParams {
	var ()
	return &WeaviateActionsDeleteParams{

		Context: ctx,
	}
}

// NewWeaviateActionsDeleteParamsWithHTTPClient creates a new WeaviateActionsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionsDeleteParamsWithHTTPClient(client *http.Client) *WeaviateActionsDeleteParams {
	var ()
	return &WeaviateActionsDeleteParams{
		HTTPClient: client,
	}
}

/*WeaviateActionsDeleteParams contains all the parameters to send to the API endpoint
for the weaviate actions delete operation typically these are written to a http.Request
*/
type WeaviateActionsDeleteParams struct {

	/*ActionID
	  Unique ID of the thing.

	*/
	ActionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate actions delete params
func (o *WeaviateActionsDeleteParams) WithTimeout(timeout time.Duration) *WeaviateActionsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate actions delete params
func (o *WeaviateActionsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate actions delete params
func (o *WeaviateActionsDeleteParams) WithContext(ctx context.Context) *WeaviateActionsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate actions delete params
func (o *WeaviateActionsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate actions delete params
func (o *WeaviateActionsDeleteParams) WithHTTPClient(client *http.Client) *WeaviateActionsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate actions delete params
func (o *WeaviateActionsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the weaviate actions delete params
func (o *WeaviateActionsDeleteParams) WithActionID(actionID strfmt.UUID) *WeaviateActionsDeleteParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the weaviate actions delete params
func (o *WeaviateActionsDeleteParams) SetActionID(actionID strfmt.UUID) {
	o.ActionID = actionID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
