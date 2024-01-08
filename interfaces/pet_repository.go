package interfaces

import (
	"pet-dex-backend/v2/entity"
)

type PetRepository interface {
	Find(id int) (*entity.Pet, error)
	Save(pet entity.Pet) error
}