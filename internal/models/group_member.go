package models

import "github.com/lukasjarosch/genki/types/nullable"

type GroupMember struct {
	ID        int           `json:"-" db:"id" validate:"required"`
	MemberID  string        `json:"member_id" db:"member_id" validate:"required"`
	GroupID   string        `json:"group_id" db:"group_id" validate:"required"`
	Name      string        `json:"name" db:"name" validate:"required"`
	Email     string        `json:"email" db:"email" validate:"required,email"`
	CreatedAt nullable.Time `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt nullable.Time `json:"updated_at" db:"updated_at" validate:"required"`
}
