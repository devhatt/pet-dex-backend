package interfaces

import (
	"pet-dex-backend/v2/entity"
)

type PetRepository interface {
	FindById(id int) (*entity.Pet, error)
	Save(pet entity.Pet) error
	Update(id int, updatePayload map[string]interface{}) error
}
