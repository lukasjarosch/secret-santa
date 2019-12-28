package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// GroupUser is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type GroupUser struct {
    ID uuid.UUID `json:"id" db:"id"`
    GroupID uuid.UUID `json:"group_id" db:"group_id"`
    Name string `json:"name" db:"name"`
    Email string `json:"email" db:"email"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (g GroupUser) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// GroupUsers is not required by pop and may be deleted
type GroupUsers []GroupUser

// String is not required by pop and may be deleted
func (g GroupUsers) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (g *GroupUser) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: g.Name, Name: "Name"},
		&validators.StringIsPresent{Field: g.Email, Name: "Email"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (g *GroupUser) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (g *GroupUser) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
