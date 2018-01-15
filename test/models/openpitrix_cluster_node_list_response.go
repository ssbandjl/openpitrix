// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixClusterNodeListResponse openpitrix cluster node list response
// swagger:model openpitrixClusterNodeListResponse
type OpenpitrixClusterNodeListResponse struct {

	// current page
	CurrentPage int32 `json:"current_page,omitempty"`

	// items
	Items OpenpitrixClusterNodeListResponseItems `json:"items"`

	// page size
	PageSize int32 `json:"page_size,omitempty"`

	// total items
	TotalItems int32 `json:"total_items,omitempty"`

	// total pages
	TotalPages int32 `json:"total_pages,omitempty"`
}

// Validate validates this openpitrix cluster node list response
func (m *OpenpitrixClusterNodeListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixClusterNodeListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixClusterNodeListResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixClusterNodeListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
