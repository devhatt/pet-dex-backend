package interfaces

import (
	"pet-dex-backend/v2/entity"

	uniqueEntity "pet-dex-backend/v2/pkg/entity"
)

type PetRepository interface {
	//FindById(id uniqueEntity.ID) (*entity.Pet, error)
	ListByUser(userID uniqueEntity.ID) ([]*entity.Pet, error)
	Save(pet entity.Pet) error
	Update(petID string, userID string, updatePayload map[string]interface{}) error
	ListByUserNoAuth()([]*entity.Pet, error)
}