package interfaces

import (
	"pet-dex-backend/v2/entity"

	"github.com/google/uuid"
)

type PetRepository interface {
	//FindById(id int) (*entity.Pet, error)
	ListPetsByUserID(userID uuid.UUID) ([]*entity.Pet, error)
	Save(pet entity.Pet) error
	Update(petID string, userID string, updatePayload map[string]interface{}) error
}
