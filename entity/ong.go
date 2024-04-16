package entity

import (
	"encoding/json"
	"log"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type Ong struct {
	ID             uniqueEntityId.ID `json:"id" db:"id"`
	Name           string            `json:"nome" db:"name"`
	Type           string            `json:"type"`
	Document       string            `json:"cnpj" db:"document"`
	AvatarURL      string            `json:"imagem" db:"avatarUrl"`
	Email          string            `json:"email"`
	Pass           string            `json:"senha" db:"pass"`
	Phone          string            `json:"telefone" db:"phone"`
	Address        Address           `json:"localizacao"`
	OpeningHours   string            `json:"horario_funcionamento"`
	AdoptionPolicy string            `json:"politica_adocao"`
	CreationDate   *time.Time        `json:"data_criacao"`
	SocialMedia    *json.RawMessage  `json:"redes_sociais"`

	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}

func NewOng(name, uType, document, avatar_url, email, phone, pass, city, state, openingHours, adoptionPolicy string, creationDate *time.Time, socialMedia json.RawMessage) *Ong {
	ongId := uniqueEntityId.NewID()

	address := NewAddress(ongId, city, state)

	var socials *json.RawMessage
	err := json.Unmarshal(socialMedia, &socials)
	if err != nil {
		log.Fatalln("error:", err)
	}

	return &Ong{
		ID:             ongId,
		Name:           name,
		Type:           uType,
		CreationDate:   creationDate,
		Document:       document,
		AvatarURL:      avatar_url,
		Email:          email,
		Pass:           pass,
		Phone:          phone,
		SocialMedia:    socials,
		Address:        address,
		OpeningHours:   openingHours,
		AdoptionPolicy: adoptionPolicy,
	}
}
