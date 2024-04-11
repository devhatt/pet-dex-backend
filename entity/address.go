package entity

import "pet-dex-backend/v2/pkg/uniqueEntityId"

type Address struct {
	ID        uniqueEntityId.ID `json:"id" db:"id"`
	UserId    uniqueEntityId.ID `json:"userId" db:"userId"`
	Address   string            `json:"address" db:"address"`
	City      string            `json:"city" db:"city"`
	State     string            `json:"state" db:"state"`
	Latitude  float64           `json:"latitude" db:"latitude"`
	Logintute float64           `json:"longitute" db:"longitute"`
}

func NewAddress(userId uniqueEntityId.ID, city, state string) Address {
	return Address{
		ID:      uniqueEntityId.NewID(),
		UserId:  userId,
		Address: "",
		City:    city,
		State:   state,
	}
}
