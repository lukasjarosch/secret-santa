package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type GroupMember struct {
	ID        int       `json:"-" db:"id" validate:"required"`
	MemberID  uuid.UUID `json:"member_id" db:"member_id" validate:"required"`
	GroupID   uuid.UUID `json:"group_id" db:"group_id" validate:"required"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validate:"required"`
}

// String returns a printable representation of a member
func (gm *GroupMember) String() string {
	groupMember, _ := json.Marshal(gm)
	return string(groupMember)
}
