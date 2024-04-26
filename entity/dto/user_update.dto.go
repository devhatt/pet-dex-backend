package dto

import (
	"pet-dex-backend/v2/entity"
	"time"
)

type UserUpdateDto struct {
	Name      string     `json:"name" db:"name"`
	Document  string     `json:"document" db:"document"`
	AvatarURL string     `json:"avatar_url" db:"avatarUrl"`
	Email     string     `json:"email" db:"email"`
	Phone     string     `json:"phone" db:"phone"`
	BirthDate *time.Time `json:"birthdate"`
}

func (dto *UserUpdateDto) ToEntity() entity.User {
	user := &entity.User{
		Name:      dto.Name,
		Document:  dto.Document,
		AvatarURL: dto.AvatarURL,
		Email:     dto.Email,
		Phone:     dto.Phone,
		BirthDate: dto.BirthDate,
	}

	if dto.BirthDate == nil {
		user.BirthDate = nil
	}

	return *user
}
