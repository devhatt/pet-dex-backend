package dto

import (
	"time"
)

type OngInsertDto struct {
	User           UserInsertDto
	OpeningHours   string     `json:"openingHours"`
	AdoptionPolicy string     `json:"adoptionPolicy"`
	BirthDate      *time.Time `json:"birthdate"`
	Links          []Link     `json:"links"`
	CreatedAt      *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      *time.Time `json:"updatedAt" db:"updated_at"`
}

type Link struct {
	URL  string
	Text string
}
