package interfaces

import (
	"pet-dex-backend/v2/entity"
)

type BreedRepository interface {
	//FindById(id uniqueEntity.ID) (*entity.Breed, error)
	List() (breeds []*entity.Breed, err error)
	Save(breed entity.Breed) error
	Update(breedID string, updatePayload map[string]interface{}) error
}
