package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngRepository interface {
	Save(ong *entity.Ong) error
	List(limit, offset int, sortBy, order string) (ongs []*dto.OngListMapper, err error)
	FindByID(ID uniqueEntityId.ID) (*dto.OngListMapper, error)
	Update(id uniqueEntityId.ID, ong entity.Ong) error
	Delete(id uniqueEntityId.ID) error
}
