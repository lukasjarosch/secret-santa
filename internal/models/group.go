package models


import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lukasjarosch/genki/types/nullable"
)

// Group represents a group of people, playing secret-santa.
type Group struct {
	ID          int             `json:"-" db:"id" validate:"required"`
	GroupID     uuid.UUID       `json:"group_id" db:"group_id" validate:"required"`
	Name        string          `json:"name" db:"name" validate:"required"`
	Description nullable.String `json:"description" db:"description"`
	Status      Status          `json:"status" db:"status"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at" validate:"required"`
}

type Status int

func (s Status) String() string {
	stat, ok := StatusString[s]
	if !ok {
		return ""
	}
	return stat
}

const (
	// UnknownStatus is the initial or any erroneous state of a group
	UnknownStatus Status = 0
	// Once the group has been actually created (saved to datastore)
	GroupCreatedStatus Status = 1
	// Members have been added to the group, email to them requesting the wishes is sent
	WishListEmailSentStatus Status = 2
	// All members saved their wishes, ready to play!
	AllWishesCollectedStatus Status = 3
	// A valid mapping permutation has been created
	SecretSantaMappedStatus Status = 4
	// Every group member received an email with the secret-santa information :)
	SecretSantaMailSent Status = 5
)

var StatusString = map[Status]string{
	UnknownStatus:            "unknown",
	GroupCreatedStatus:       "group_created",
	WishListEmailSentStatus:  "wishlist_email_sent",
	AllWishesCollectedStatus: "all_wishes_collected",
	SecretSantaMappedStatus:  "secret_santa_mapping_created",
	SecretSantaMailSent:      "secret_santa_mail_sent",
}

// String returns a printable representation of a group
func (g *Group) String() string {
	jsonGroup, _ := json.Marshal(g)
	return string(jsonGroup)
}
