package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
)

type OngRepository interface {
	Save(ong *entity.Ong) error
	List(limit, offset int, sortBy, order string) (ongs []*dto.OngList, err error)
}
