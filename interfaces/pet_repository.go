package interfaces

import (
	"pet-dex-backend/v2/entity"

	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type PetRepository interface {
	//FindById(id uniqueEntity.ID) (*entity.Pet, error)
	ListByUser(userID uniqueEntityId.ID) ([]*entity.Pet, error)
	Save(pet entity.Pet) error
	Update(petID string, userID string, updatePayload map[string]interface{}) error
}
