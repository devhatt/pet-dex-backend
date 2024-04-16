package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngRepository interface {
	Save(ong *entity.Ong, userId uniqueEntityId.ID) error
	SaveUser(user *entity.User) error
	AdressRepo
}
