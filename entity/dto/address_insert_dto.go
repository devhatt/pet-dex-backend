package dto

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type AddressInsertDto struct {
	UserId    uniqueEntityId.ID `json:"userId" db:"userId"`
	Address   string            `json:"address" db:"address"`
	City      string            `json:"city" db:"city"`
	State     string            `json:"state" db:"state"`
	Latitude  float64           `json:"latitude" db:"latitude"`
	Longitude float64           `json:"longitude" db:"longitude"`
}
