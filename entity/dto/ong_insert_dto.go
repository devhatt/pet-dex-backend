package dto

import (
	"encoding/json"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type OngInsertDto struct {
	UserID         uniqueEntityId.ID
	Name           string           `json:"name"`
	Type           string           `json:"type"`
	Document       string           `json:"document"`
	AvatarURL      string           `json:"avatar_url"`
	Email          string           `json:"email"`
	Pass           string           `json:"pass"`
	City           string           `json:"city"`
	Phone          string           `json:"phone"`
	State          string           `json:"state"`
	OpeningHours   string           `json:"openingHours"`
	AdoptionPolicy string           `json:"adoptionPolicy"`
	BirthDate      *time.Time       `json:"birthdate"`
	Links          *json.RawMessage `json:"links"`

	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
