package dto

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngDeleteDto struct {
	ID uniqueEntityId.ID `json:"id"`
}
