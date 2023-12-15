package interfaces

import "pet-dex-backend/v2/entity"

type PetRepository interface {
	Save(entity.Pet) (error)
}