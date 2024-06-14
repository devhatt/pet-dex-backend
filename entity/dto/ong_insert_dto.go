package dto

import (
	"encoding/json"
	"time"
)

type OngInsertDto struct {
	User           UserInsertDto
	OpeningHours   string           `json:"openingHours"`
	AdoptionPolicy string           `json:"adoptionPolicy"`
	BirthDate      *time.Time       `json:"birthdate"`
	Links          *json.RawMessage `json:"links"`
	CreatedAt      *time.Time       `json:"createdAt" db:"created_at"`
	UpdatedAt      *time.Time       `json:"updatedAt" db:"updated_at"`
}
