package dto

import (
	"encoding/json"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngList struct {
	ID uniqueEntityId.ID `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Phone string `json:"phone" db:"phone"`
	State string `json:"state"`
	OpeningHours string `json:"openingHours" db:"openingHours"`
	AdoptionPolicy string `json:"adoptionPolicy" db:"adoptionPolicy"`
	Links *json.RawMessage `json:"links"`

}