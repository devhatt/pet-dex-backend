package dto

import (
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

func (u *UserInsertDto) Validate() bool {
	return regex.MatchString(u.Pass)
}
