package dto

import (
	"encoding/json"
)

type OngUpdateDto struct {
	/* 	Name           string           `json:"name"`
	   	Document       string           `json:"document"`
	   	AvatarURL      string           `json:"avatar_url"`
	*/
	User           UserUpdateDto
	Phone          string           `json:"phone"`
	OpeningHours   string           `json:"openingHours"`
	AdoptionPolicy string           `json:"adoptionPolicy"`
	Links          *json.RawMessage `json:"links"`
}
