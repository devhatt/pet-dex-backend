package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngRepository interface {
	FindByID(ID uniqueEntityId.ID) (*entity.Ong, error)
	Save(ong *entity.Ong) error
	FindByID(ID uniqueEntityId.ID) (*entity.Ong, error)
}
