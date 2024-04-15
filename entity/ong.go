package entity

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"
)

type Ong struct {
	ID        uniqueEntityId.ID `json:"id" db:"id"`
	Name      string            `json:"name" db:"name"`
	Type      string            `json:"type" db:"type"`
	Document  string            `json:"document" db:"document"`
	AvatarURL string            `json:"avatar_url" db:"avatarUrl"`
	Email     string            `json:"email" db:"email"`
	Phone     string            `json:"phone" db:"phone"`
	Pass      string            `json:"pass" db:"pass"`
	CreatedAt *time.Time        `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time        `json:"updatedAt" db:"updated_at"`
	Adresses  Address           `json:"addresses"`
}
