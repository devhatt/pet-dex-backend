package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type AdressRepo interface {
	SaveAddress(addr *entity.Address) error
	FindAddressByUserID(ID uniqueEntityId.ID) (*entity.Address, error)
}
