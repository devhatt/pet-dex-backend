package entity

import (
	"encoding/json"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type Ong struct {
	ID             uniqueEntityId.ID `json:"id" db:"id"`
	UserID         uniqueEntityId.ID `db:"userId"`
	User           User              `json:"user"`
	Phone          string            `json:"phone" db:"phone"`
	OpeningHours   string            `json:"openingHours" db:"openingHours"`
	AdoptionPolicy string            `json:"adoptionPolicy" db:"adoptionPolicy"`
	Links          *json.RawMessage  `json:"links"`

	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
