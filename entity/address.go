package entity

import (
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type Address struct {
	ID        uniqueEntityId.ID `json:"id" db:"id"`
	UserId    uniqueEntityId.ID `json:"userId" db:"userId"`
	Address   string            `json:"address" db:"address"`
	City      string            `json:"city" db:"city"`
	State     string            `json:"state" db:"state"`
	Latitude  float64           `json:"latitude" db:"latitude"`
	Longitude float64           `json:"longitude" db:"longitude"`
}

func NewAddress(address dto.AddressInsertDto) *Address {
	return &Address{
		ID:        uniqueEntityId.NewID(),
		UserId:    address.UserId,
		Address:   "",
		City:      address.City,
		State:     address.State,
		Latitude:  address.Latitude,
		Longitude: address.Longitude,
	}
}
