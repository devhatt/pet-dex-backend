package interfaces

import (
	"os"
	"pet-dex-backend/v2/entity"
)

type PetRepository interface {
	FindById(id int) (*entity.Pet, error)
	Save(pet entity.Pet) error
	FindNoAuthPets() (*os.File, error)
}
