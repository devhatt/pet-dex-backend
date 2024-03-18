package interfaces

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	uniqueEntity "pet-dex-backend/v2/pkg/entity"
)

type BreedRepository interface {
	FindById(id uniqueEntity.ID) (*entity.Breed, error)
	List() (breeds []*dto.BreedList, err error)
	Save(breed entity.Breed) error
	Update(breedID string, updatePayload map[string]interface{}) error
}
