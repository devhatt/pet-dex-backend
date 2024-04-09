package interfaces

import (
	"pet-dex-backend/v2/entity/dto"
)

type BreedRepository interface {
	List() (breeds []*dto.BreedList, err error)
}
