package interfaces

import (
	"pet-dex-backend/v2/entity"
)

type PetRepository interface {
	//FindById(id int) (*entity.Pet, error)
	Save(pet entity.Pet) error
	Update(petID string, userID string, updatePayload map[string]interface{}) error
	ListUserPets(userID int) ([]*entity.Pet, error)
}
