package models

import "github.com/lukasjarosch/genki/types/nullable"

type WishList []Wish

type Wish struct {
	ID        int             `json:"-" db:"id" validate:"required"`
	MemberID  string          `json:"member_id" db:"member_id" validate:"required"`
	Name      string          `json:"name" db:"name" validate:"required"`
	Notes     nullable.String `json:"notes" db:"notes"`
	CreatedAt nullable.Time   `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt nullable.Time   `json:"updated_at" db:"updated_at" validate:"required"`
}
