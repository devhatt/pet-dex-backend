package dto

import (
	"pet-dex-backend/v2/entity"
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

func NewAddress(address AddressInsertDto) *entity.Address {
	return &entity.Address{
		ID:        uniqueEntityId.NewID(),
		UserId:    address.UserId,
		Address:   "",
		City:      address.City,
		State:     address.State,
		Latitude:  address.Latitude,
		Longitude: address.Longitude,
	}
}
