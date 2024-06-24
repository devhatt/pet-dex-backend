package interfaces

import (
	"pet-dex-backend/v2/entity"

	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type PetRepository interface {
	ListByUser(userID uniqueEntityId.ID) ([]*entity.Pet, error)
	FindByID(ID uniqueEntityId.ID) (*entity.Pet, error)
	Save(pet *entity.Pet) error
	Update(petID string, userID string, petToUpdate *entity.Pet) error
	ListAllByPage(page int) ([]*entity.Pet, error)
}
