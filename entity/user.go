package entity

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type User struct {
	ID uniqueEntityId.ID  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Document string `json:"document"`
	AvatarURL string `json:"avatar_url"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Pass string `json:"pass"`
	Adresses addresses `json:"addresses"`
	Person Persons `json:"person"`
}

type Persons struct {
	ID uniqueEntityId.ID `json:"id"`
	BirthDate *time.Time `json:"birthdate"`
}

type addresses struct {
	ID uniqueEntityId.ID `json:"id"`
	Address string `json:"addresses"`
	City string `json:"city"`
	State string `json:"state"`
	Latitude float64 `json:"latitude"`
	Logintude float64 `json:"longitude"`
}