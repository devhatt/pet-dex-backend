package entity

import (
	"encoding/json"
	"log"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type Ong struct {
	ID             uniqueEntityId.ID `json:"id" db:"id"`
	UserID         uniqueEntityId.ID `json:"userId" db:"userId"`
	User           User              `json:"user"`
	Phone          string            `json:"phone" db:"phone"`
	OpeningHours   string            `json:"openingHours" db:"openingHours"`
	AdoptionPolicy string            `json:"adoptionPolicy" db:"adoptionPolicy"`
	Links          *json.RawMessage  `json:"links"`

	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

func NewOng(name, uType, document, avatar_url, email, phone, pass, city, state, openingHours, adoptionPolicy string, creationDate *time.Time, links json.RawMessage) *Ong {
	ongId := uniqueEntityId.NewID()

	user := NewUser(name, uType, document, avatar_url, email, phone, pass, city, state, creationDate)

	var socials *json.RawMessage
	err := json.Unmarshal(links, &socials)
	if err != nil {
		log.Fatalln("error:", err)
	}

	return &Ong{
		ID:             ongId,
		UserID:         user.ID,
		User:           *user,
		Phone:          phone,
		Links:          socials,
		OpeningHours:   openingHours,
		AdoptionPolicy: adoptionPolicy,
	}
}
