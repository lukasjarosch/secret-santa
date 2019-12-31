package models

import "github.com/lukasjarosch/genki/types/nullable"

type Group struct {
	ID          int           `json:"-" db:"id" validate:"required"`
	GroupID     string        `json:"group_id" db:"group_id" validate:"required"`
	Name        string        `json:"name" db:"name" validate:"required"`
	CreatorName string        `json:"creator_name" db:"creator_name" validate:"required"`
	Members     []GroupMember `json:"members" db:"-"`
	Status      GroupStatus   `json:"status" db:"status" validate:"required"`
	CreatedAt   nullable.Time `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt   nullable.Time `json:"updated_at" db:"updated_at" validate:"required"`
}
