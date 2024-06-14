package dto

import (
	"time"
)

type UserUpdateDto struct {
	Name      string     `json:"name" db:"name"`
	Document  string     `json:"document" db:"document"`
	AvatarURL string     `json:"avatar_url" db:"avatarUrl"`
	Email     string     `json:"email" db:"email"`
	Phone     string     `json:"phone" db:"phone"`
	Role      string     `json:"role"`
	BirthDate *time.Time `json:"birthdate"`
}
