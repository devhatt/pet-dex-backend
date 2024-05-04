package entity

import (
	"pet-dex-backend/v2/entity/dto"
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
	Adresses  Address           `json:"addresses"`
}

func NewUser(name, uType, document, avatar_url, email, phone, pass, city, state string, birthdate *time.Time) *User {
	userId := uniqueEntityId.NewID()

	var addressDto dto.AddressInsertDto
	addressDto.UserId = userId
	addressDto.City = city
	addressDto.State = state

	address := NewAddress(addressDto)

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
		Adresses:  *address,
	}
}

func UserToUpdate(dtoUpdate dto.UserUpdateDto) User {
	user := User{
		Name:      dtoUpdate.Name,
		Document:  dtoUpdate.Document,
		AvatarURL: dtoUpdate.AvatarURL,
		Email:     dtoUpdate.Email,
		Phone:     dtoUpdate.Phone,
		BirthDate: dtoUpdate.BirthDate,
	}

	if dtoUpdate.BirthDate == nil {
		user.BirthDate = nil
	}

	return user
}
