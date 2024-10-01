package entity

import (
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type Ong struct {
	ID             uniqueEntityId.ID `json:"id" db:"id"`
	UserID         uniqueEntityId.ID `db:"userId"`
	User           User              `json:"user"`
	Phone          string            `json:"phone" db:"phone"`
	OpeningHours   string            `json:"openingHours" db:"openingHours"`
	AdoptionPolicy string            `json:"adoptionPolicy" db:"adoptionPolicy"`
	Links          []dto.LinkDto     `json:"links"`

	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

func NewOng(ong dto.OngInsertDto) *Ong {
	ongId := uniqueEntityId.NewID()

	user := NewUser(ong.User)

	return &Ong{
		ID:             ongId,
		UserID:         user.ID,
		User:           *user,
		Phone:          user.Phone,
		Links:          []dto.LinkDto{},
		OpeningHours:   ong.OpeningHours,
		AdoptionPolicy: ong.AdoptionPolicy,
	}
}

func OngToUpdate(ong dto.OngUpdateDto) (*Ong, error) {
	user := User{
		Name:      ong.User.Name,
		Document:  ong.User.Document,
		AvatarURL: ong.User.AvatarURL,
		Email:     ong.User.Email,
		Phone:     ong.User.Phone,
		BirthDate: ong.User.BirthDate,
	}

	return &Ong{
		User:           user,
		Phone:          user.Phone,
		Links:          ong.Links,
		OpeningHours:   ong.OpeningHours,
		AdoptionPolicy: ong.AdoptionPolicy,
	}, nil
}
