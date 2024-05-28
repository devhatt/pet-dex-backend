package dto

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngListMapper struct {
	ID             uniqueEntityId.ID   `json:"id" db:"id"`
	UserId         uniqueEntityId.ID   `json:"userId" db:"userId"`
	Name           string              `json:"name" db:"name"`
	Address        string              `json:"address" db:"address"`
	City           string              `json:"city" db:"city"`
	State          string              `json:"state" db:"state"`
	Phone          string              `json:"phone" db:"phone"`
	OpeningHours   string              `json:"openingHours" db:"openingHours"`
	AdoptionPolicy string              `json:"adoptionPolicy" db:"adoptionPolicy"`
	Links          string              `json:"links" db:"links"`
}