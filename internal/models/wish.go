package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lukasjarosch/genki/types/nullable"
)

type Wish struct {
	ID          int             `json:"-" db:"id" validate:"required"`
	WishID      uuid.UUID       `json:"wish_id" db:"wish_id" validate:"required"`
	Name        string          `json:"name" db:"name" validate:"required"`
	Description nullable.String `json:"description" db:"description"`
}

// String returns a printable representation of a wish
func (w Wish) String() string {
	wish, _ := json.Marshal(w)
	return string(wish)
}
