package entity

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type Ong struct {
	UserId      uniqueEntityId.ID `json:"user_id"`
	Cnpj        string            `json:"cnpj"`
	Email       string            `json:"email"`
	Address     Address           `json:"localizacao"`
	Image       string            `json:"image"`
	SocialMedia SocialMedia       `json:"social_media"`
}

type SocialMedia struct {
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
}
