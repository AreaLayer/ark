// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RegisterOutputsForNextRoundRequest v1 register outputs for next round request
//
// swagger:model v1RegisterOutputsForNextRoundRequest
type V1RegisterOutputsForNextRoundRequest struct {

	// Mocks wabisabi's blinded credentials.
	ID string `json:"id,omitempty"`

	// List of receivers for a registered payment.
	Outputs []*V1Output `json:"outputs"`
}

// Validate validates this v1 register outputs for next round request
func (m *V1RegisterOutputsForNextRoundRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RegisterOutputsForNextRoundRequest) validateOutputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Outputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Outputs); i++ {
		if swag.IsZero(m.Outputs[i]) { // not required
			continue
		}

		if m.Outputs[i] != nil {
			if err := m.Outputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 register outputs for next round request based on the context it is used
func (m *V1RegisterOutputsForNextRoundRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOutputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RegisterOutputsForNextRoundRequest) contextValidateOutputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Outputs); i++ {

		if m.Outputs[i] != nil {

			if swag.IsZero(m.Outputs[i]) { // not required
				return nil
			}

			if err := m.Outputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RegisterOutputsForNextRoundRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RegisterOutputsForNextRoundRequest) UnmarshalBinary(b []byte) error {
	var res V1RegisterOutputsForNextRoundRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}