package dto

import "time"

type UserInsertDto struct {
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Document  string            `json:"document"`
	AvatarURL string            `json:"avatar_url"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Pass      string            `json:"pass"`
	BirthDate *time.Time        `json:"birthdate"`
	City      string            `json:"city"`
	State     string            `json:"state"`
}

