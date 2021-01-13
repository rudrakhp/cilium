// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPAMResponse IPAM configuration of an endpoint
//
// swagger:model IPAMResponse
type IPAMResponse struct {

	// address
	// Required: true
	Address *AddressPair `json:"address"`

	// host addressing
	// Required: true
	HostAddressing *NodeAddressing `json:"host-addressing"`

	// ipv4
	IPV4 *IPAMAddressResponse `json:"ipv4,omitempty"`

	// ipv6
	IPV6 *IPAMAddressResponse `json:"ipv6,omitempty"`
}

// Validate validates this IP a m response
func (m *IPAMResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostAddressing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAMResponse) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *IPAMResponse) validateHostAddressing(formats strfmt.Registry) error {

	if err := validate.Required("host-addressing", "body", m.HostAddressing); err != nil {
		return err
	}

	if m.HostAddressing != nil {
		if err := m.HostAddressing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host-addressing")
			}
			return err
		}
	}

	return nil
}

func (m *IPAMResponse) validateIPV4(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4) { // not required
		return nil
	}

	if m.IPV4 != nil {
		if err := m.IPV4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv4")
			}
			return err
		}
	}

	return nil
}

func (m *IPAMResponse) validateIPV6(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6) { // not required
		return nil
	}

	if m.IPV6 != nil {
		if err := m.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAMResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAMResponse) UnmarshalBinary(b []byte) error {
	var res IPAMResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}