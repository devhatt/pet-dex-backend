package entity

import (
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type User struct {
	ID                       uniqueEntityId.ID `json:"id" db:"id"`
	Name                     string            `json:"name" db:"name"`
	Type                     string            `json:"type" db:"type"`
	Document                 string            `json:"document" db:"document"`
	AvatarURL                string            `json:"avatar_url" db:"avatarUrl"`
	Email                    string            `json:"email" db:"email"`
	Phone                    string            `json:"phone" db:"phone"`
	Pass                     string            `json:"pass" db:"pass"`
	PushNotificationsEnabled *bool             `json:"pushNotificationsEnabled" db:"pushNotificationsEnabled"`
	BirthDate                *time.Time        `json:"birthdate"`
	Role                     string            `json:"role" db:"role"`
	CreatedAt                *time.Time        `json:"createdAt" db:"created_at"`
	UpdatedAt                *time.Time        `json:"updatedAt" db:"updated_at"`
	Adresses                 Address           `json:"addresses"`
}

func NewUser(user dto.UserInsertDto) *User {
	userId := uniqueEntityId.NewID()

	var addressDto dto.AddressInsertDto
	addressDto.UserId = userId
	addressDto.City = user.City
	addressDto.State = user.State

	address := NewAddress(addressDto)

	return &User{
		ID:        userId,
		Name:      user.Name,
		Type:      user.Type,
		Document:  user.Document,
		AvatarURL: user.AvatarURL,
		Email:     user.Email,
		Phone:     user.Phone,
		Pass:      user.Pass,
		BirthDate: user.BirthDate,
		Role:      user.Role,
		Adresses:  *address,
	}
}

func UserToUpdate(dto *dto.UserUpdateDto) User {
	user := &User{
		Name:      dto.Name,
		Document:  dto.Document,
		AvatarURL: dto.AvatarURL,
		Email:     dto.Email,
		Phone:     dto.Phone,
		BirthDate: dto.BirthDate,
		Role:      dto.Role,
	}

	if dtoUpdate.BirthDate == nil {
		user.BirthDate = nil
	}

	return user
}
