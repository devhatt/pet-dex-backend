package dto

import (
	"fmt"
	"net/mail"
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
