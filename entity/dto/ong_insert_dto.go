package dto

import (
	"encoding/json"
	"time"
)

type OngInsertDto struct {
	Name           string           `json:"name"`
	Type           string           `json:"type"`
	Phone          string           `json:"phone"`
	Document       string           `json:"document"`
	AvatarURL      string           `json:"avatar_url"`
	Email          string           `json:"email"`
	Pass           string           `json:"pass"`
	City           string           `json:"city"`
	State          string           `json:"state"`
	OpeningHours   string           `json:"openingHours"`
	AdoptionPolicy string           `json:"adoptionPolicy"`
	BirthDate      *time.Time       `json:"birthdate"`
	Links          *json.RawMessage `json:"links"`

	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
