package dto

import (
	"fmt"
	"net/mail"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"regexp"
	"time"
)

var regex = regexp.MustCompile(`^[A-Za-z\d\W]{6,}$`)

type UserInsertDto struct {
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Document  string     `json:"document"`
	AvatarURL string     `json:"avatar_url"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Pass      string     `json:"pass"`
	BirthDate *time.Time `json:"birthdate"`
	City      string     `json:"city"`
	State     string     `json:"state"`
}

func (u *UserInsertDto) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("invalid name")
	}

	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return fmt.Errorf("invalid email")
	}

	if !regex.MatchString(u.Pass) {
		return fmt.Errorf("invalid password format")
	}
	return nil
}

func NewUser(name, uType, document, avatar_url, email, phone, pass, city, state string, birthdate *time.Time) *entity.User {
	userId := uniqueEntityId.NewID()

	var addressDto AddressInsertDto
	addressDto.UserId = userId
	addressDto.City = city
	addressDto.State = state

	address := NewAddress(addressDto)

	return &entity.User{
		ID:        userId,
		Name:      name,
		Type:      uType,
		Document:  document,
		AvatarURL: avatar_url,
		Email:     email,
		Phone:     phone,
		Pass:      pass,
		BirthDate: birthdate,
		Adresses:  *address,
	}
}
