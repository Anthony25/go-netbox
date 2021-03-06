// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualMachineInterface virtual machine interface
// swagger:model VirtualMachineInterface
type VirtualMachineInterface struct {

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// form factor
	FormFactor *VirtualMachineInterfaceFormFactor `json:"form_factor,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// MAC Address
	MacAddress *string `json:"mac_address,omitempty"`

	// mode
	Mode *VirtualMachineInterfaceMode `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []*NestedVLAN `json:"tagged_vlans"`

	// tags
	Tags []string `json:"tags"`

	// untagged vlan
	UntaggedVlan *NestedVLAN `json:"untagged_vlan,omitempty"`

	// virtual machine
	// Required: true
	VirtualMachine *NestedVirtualMachine `json:"virtual_machine"`
}

// Validate validates this virtual machine interface
func (m *VirtualMachineInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormFactor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUntaggedVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineInterface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineInterface) validateFormFactor(formats strfmt.Registry) error {

	if swag.IsZero(m.FormFactor) { // not required
		return nil
	}

	if m.FormFactor != nil {
		if err := m.FormFactor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("form_factor")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineInterface) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineInterface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(*m.Mtu), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(*m.Mtu), 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineInterface) validateTaggedVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	for i := 0; i < len(m.TaggedVlans); i++ {
		if swag.IsZero(m.TaggedVlans[i]) { // not required
			continue
		}

		if m.TaggedVlans[i] != nil {
			if err := m.TaggedVlans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagged_vlans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualMachineInterface) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *VirtualMachineInterface) validateUntaggedVlan(formats strfmt.Registry) error {

	if swag.IsZero(m.UntaggedVlan) { // not required
		return nil
	}

	if m.UntaggedVlan != nil {
		if err := m.UntaggedVlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("untagged_vlan")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineInterface) validateVirtualMachine(formats strfmt.Registry) error {

	if err := validate.Required("virtual_machine", "body", m.VirtualMachine); err != nil {
		return err
	}

	if m.VirtualMachine != nil {
		if err := m.VirtualMachine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_machine")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineInterface) UnmarshalBinary(b []byte) error {
	var res VirtualMachineInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VirtualMachineInterfaceFormFactor Form factor
// swagger:model VirtualMachineInterfaceFormFactor
type VirtualMachineInterfaceFormFactor struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *int64 `json:"value"`
}

func (m *VirtualMachineInterfaceFormFactor) UnmarshalJSON(b []byte) error {
	type VirtualMachineInterfaceFormFactorAlias VirtualMachineInterfaceFormFactor
	var t VirtualMachineInterfaceFormFactorAlias
	if err := json.Unmarshal([]byte("{\"label\":\"Virtual\",\"value\":0}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = VirtualMachineInterfaceFormFactor(t)
	return nil
}

// Validate validates this virtual machine interface form factor
func (m *VirtualMachineInterfaceFormFactor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineInterfaceFormFactor) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("form_factor"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineInterfaceFormFactor) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("form_factor"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineInterfaceFormFactor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineInterfaceFormFactor) UnmarshalBinary(b []byte) error {
	var res VirtualMachineInterfaceFormFactor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VirtualMachineInterfaceMode Mode
// swagger:model VirtualMachineInterfaceMode
type VirtualMachineInterfaceMode struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *int64 `json:"value"`
}

// Validate validates this virtual machine interface mode
func (m *VirtualMachineInterfaceMode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineInterfaceMode) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineInterfaceMode) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineInterfaceMode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineInterfaceMode) UnmarshalBinary(b []byte) error {
	var res VirtualMachineInterfaceMode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
