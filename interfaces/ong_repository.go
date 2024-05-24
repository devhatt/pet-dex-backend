package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngRepository interface {
	Save(ong *entity.Ong) error
	List(limit, offset int, sortBy, order string) (ongs []*dto.OngListDto, err error)
	Update(id uniqueEntityId.ID, ong entity.Ong) error
	FindByID(ID uniqueEntityId.ID) (*entity.Ong, error)
}
