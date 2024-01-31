package interfaces

import (
	"pet-dex-backend/v2/entity"
)

type PetRepository interface {
	FindById(id string) (*entity.Pet, error)
	Save(pet entity.Pet) error
	Update(id string, updatePayload map[string]interface{}) error
}
