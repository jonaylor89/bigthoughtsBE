package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type Class struct {
	class_id   int    `db:"class_id"`
	class_num  string `db:"class_num"`
	class_name string `db:"class_name"`
}

// String is not required by pop and may be deleted
func (c Class) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Classes is not required by pop and may be deleted
type Classes []Class

// String is not required by pop and may be deleted
func (c Classes) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Class) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Class) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Class) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
