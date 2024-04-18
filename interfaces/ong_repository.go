package interfaces

import (
	"pet-dex-backend/v2/entity"
)

type OngRepository interface {
	Save(ong *entity.Ong) error
	SaveUser(user *entity.User) error
	AdressRepo
}
