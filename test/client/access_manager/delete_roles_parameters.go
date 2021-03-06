// Code generated by go-swagger; DO NOT EDIT.

package access_manager

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

	"openpitrix.io/openpitrix/test/models"
)

// NewDeleteRolesParams creates a new DeleteRolesParams object
// with the default values initialized.
func NewDeleteRolesParams() *DeleteRolesParams {
	var ()
	return &DeleteRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRolesParamsWithTimeout creates a new DeleteRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRolesParamsWithTimeout(timeout time.Duration) *DeleteRolesParams {
	var ()
	return &DeleteRolesParams{

		timeout: timeout,
	}
}

// NewDeleteRolesParamsWithContext creates a new DeleteRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRolesParamsWithContext(ctx context.Context) *DeleteRolesParams {
	var ()
	return &DeleteRolesParams{

		Context: ctx,
	}
}

// NewDeleteRolesParamsWithHTTPClient creates a new DeleteRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRolesParamsWithHTTPClient(client *http.Client) *DeleteRolesParams {
	var ()
	return &DeleteRolesParams{
		HTTPClient: client,
	}
}

/*DeleteRolesParams contains all the parameters to send to the API endpoint
for the delete roles operation typically these are written to a http.Request
*/
type DeleteRolesParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteRolesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete roles params
func (o *DeleteRolesParams) WithTimeout(timeout time.Duration) *DeleteRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete roles params
func (o *DeleteRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete roles params
func (o *DeleteRolesParams) WithContext(ctx context.Context) *DeleteRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete roles params
func (o *DeleteRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete roles params
func (o *DeleteRolesParams) WithHTTPClient(client *http.Client) *DeleteRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete roles params
func (o *DeleteRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete roles params
func (o *DeleteRolesParams) WithBody(body *models.OpenpitrixDeleteRolesRequest) *DeleteRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete roles params
func (o *DeleteRolesParams) SetBody(body *models.OpenpitrixDeleteRolesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
