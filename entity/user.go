package entity

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type User struct {
	ID        uniqueEntityId.ID `json:"id" db:"id"`
	Name      string            `json:"name" db:"name"`
	Type      string            `json:"type" db:"type"`
	Document  string            `json:"document" db:"document"`
	AvatarURL string            `json:"avatar_url" db:"avatarUrl"`
	Email     string            `json:"email" db:"email"`
	Phone     string            `json:"phone" db:"phone"`
	Pass      string            `json:"pass" db:"pass"`
	BirthDate *time.Time        `json:"birthdate"`
	CreatedAt *time.Time        `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time        `json:"updatedAt" db:"updated_at"`
	Adresses  Address         `json:"addresses"`
}

func NewUser(name, uType, document, avatar_url, email, phone, pass, city, state string, birthdate *time.Time) *User {
	userId := uniqueEntityId.NewID()

	address := NewAddress(userId, city, state)

	return &User{
		ID:        userId,
		Name:      name,
		Type:      uType,
		Document:  document,
		AvatarURL: avatar_url,
		Email:     email,
		Phone:     phone,
		Pass:      pass,
		BirthDate: birthdate,
		Adresses:  address,
	}
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

type Address struct {
	ID        uniqueEntityId.ID `json:"id" db:"id"`
	UserId    uniqueEntityId.ID `json:"userId" db:"userId"`
	Address   string            `json:"address" db:"address"`
	City      string            `json:"city" db:"city"`
	State     string            `json:"state" db:"state"`
	Latitude  float64           `json:"latitude" db:"latitude"`
	Logintute float64           `json:"longitute" db:"longitute"`
}
