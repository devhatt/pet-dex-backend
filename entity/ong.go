package entity

import (
	"encoding/json"
	"log"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type Ong struct {
	ID             uniqueEntityId.ID `json:"id" db:"id"`
	User           User
	OpeningHours   string           `json:"horario_funcionamento"`
	AdoptionPolicy string           `json:"politica_adocao"`
	SocialMedia    *json.RawMessage `json:"redes_sociais"`

	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

func NewOng(name, uType, document, avatar_url, email, phone, pass, city, state, openingHours, adoptionPolicy string, creationDate *time.Time, socialMedia json.RawMessage) *Ong {
	ongId := uniqueEntityId.NewID()

	user := NewUser(name, uType, document, avatar_url, email, phone, pass, city, state, creationDate)

	var socials *json.RawMessage
	err := json.Unmarshal(socialMedia, &socials)
	if err != nil {
		log.Fatalln("error:", err)
	}

	return &Ong{
		ID:             ongId,
		User:           *user,
		SocialMedia:    socials,
		OpeningHours:   openingHours,
		AdoptionPolicy: adoptionPolicy,
	}
}
