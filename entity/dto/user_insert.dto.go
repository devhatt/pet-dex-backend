package dto

import (
	"fmt"
	"net/mail"
	"pet-dex-backend/v2/pkg/utils"
	"slices"
	"time"
)

var userTypes = []string{"juridica", "fisica"}

type UserInsertDto struct {
	Name      string     `json:"name" example:"Claúdio"`
	Type      string     `json:"type" example:"fisica"`
	Document  string     `json:"document" example:"12345678900"`
	AvatarURL string     `json:"avatar_url" example:"https://example.com/avatar.jpg"`
	Email     string     `json:"email" example:"claudio@example.com"`
	Phone     string     `json:"phone" example:"21912345678"`
	Pass      string     `json:"pass" example:"Senhasegur@123"`
	BirthDate *time.Time `json:"birthdate" example:"2006-01-02T15:04:05Z"`
	City      string     `json:"city" example:"São Paulo"`
	State     string     `json:"state" example:"São Paulo"`
	Role      string     `json:"role" example:"developer"`
}

func (u *UserInsertDto) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("invalid name")
	}

	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return fmt.Errorf("invalid email")
	}

	if !slices.Contains(userTypes, u.Type) {
		return fmt.Errorf("type can only be 'juridica' or 'fisica'")
	}

	if u.Pass == "" {
		return fmt.Errorf("password cannot be empty")
	}

	if !utils.IsValidPassword(u.Pass) {
		return fmt.Errorf("invalid password format")
	}
	return nil
}
