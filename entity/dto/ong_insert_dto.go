package dto

import (
	"encoding/json"
	"log"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type OngInsertDto struct {
	UserID         uniqueEntityId.ID
	Name           string           `json:"name"`
	Type           string           `json:"type"`
	Document       string           `json:"document"`
	AvatarURL      string           `json:"avatar_url"`
	Email          string           `json:"email"`
	Pass           string           `json:"pass"`
	City           string           `json:"city"`
	Phone          string           `json:"phone"`
	State          string           `json:"state"`
	OpeningHours   string           `json:"openingHours"`
	AdoptionPolicy string           `json:"adoptionPolicy"`
	BirthDate      *time.Time       `json:"birthdate"`
	Links          *json.RawMessage `json:"links"`

	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

func NewOng(ong OngInsertDto) *entity.Ong {
	ongId := uniqueEntityId.NewID()

	user := NewUser(ong.Name, ong.Type, ong.Document, ong.AvatarURL, ong.Email, ong.Phone, ong.Pass, ong.City, ong.State, ong.BirthDate)

	var socials *json.RawMessage
	err := json.Unmarshal(*ong.Links, &socials)
	if err != nil {
		log.Fatalln("error:", err)
	}

	return &entity.Ong{
		ID:             ongId,
		UserID:         user.ID,
		User:           *user,
		Phone:          user.Phone,
		Links:          socials,
		OpeningHours:   ong.OpeningHours,
		AdoptionPolicy: ong.AdoptionPolicy,
	}
}
