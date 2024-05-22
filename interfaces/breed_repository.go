package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type BreedRepository interface {
	List() (breeds []*dto.BreedList, err error)
	FindByID(ID uniqueEntityId.ID) (*entity.Breed, error)
}
