package interfaces

import "pet-dex-backend/v2/entity"

type AdressRepo interface {
	SaveAddress(addr *entity.Address) error
}
