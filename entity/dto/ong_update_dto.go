package dto

import (
	"encoding/json"
)

type OngUpdateDto struct {
	Phone          string `json:"phone" db:"phone"`
	User           UserUpdateDto
	OpeningHours   string           `json:"openingHours"`
	AdoptionPolicy string           `json:"adoptionPolicy"`
	Links          *json.RawMessage `json:"links"`
}
