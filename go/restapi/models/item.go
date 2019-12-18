// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Item item
// swagger:model item
type Item struct {

	// description
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// key value fields
	KeyValueFields []*KeyValue `json:"keyValueFields"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// picture
	// Required: true
	// Min Length: 1
	Picture *string `json:"picture"`
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyValueFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePicture(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *Item) validateKeyValueFields(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyValueFields) { // not required
		return nil
	}

	for i := 0; i < len(m.KeyValueFields); i++ {
		if swag.IsZero(m.KeyValueFields[i]) { // not required
			continue
		}

		if m.KeyValueFields[i] != nil {
			if err := m.KeyValueFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keyValueFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Item) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Item) validatePicture(formats strfmt.Registry) error {

	if err := validate.Required("picture", "body", m.Picture); err != nil {
		return err
	}

	if err := validate.MinLength("picture", "body", string(*m.Picture), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
