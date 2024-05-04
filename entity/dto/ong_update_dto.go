package dto

import (
	"encoding/json"
)

type OngUpdateDto struct {
	/* Name           string           `json:"name" db:"name"`
	Document       string           `json:"document" db:"document"`
	AvatarURL      string           `json:"avatar_url" db:"avatarUrl"`
	Email          string           `json:"email" db:"email"`*/
	Phone          string `json:"phone" db:"phone"`
	User           UserUpdateDto
	OpeningHours   string           `json:"openingHours"`
	AdoptionPolicy string           `json:"adoptionPolicy"`
	Links          *json.RawMessage `json:"links"`
}
