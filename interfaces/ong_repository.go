package interfaces

import (
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngRepository interface {
	FindByID(ID uniqueEntityId.ID)
}
