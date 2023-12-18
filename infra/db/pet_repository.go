package db

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type PetRepository struct {

}

func NewPetRepository () interfaces.PetRepository{
	return &PetRepository{}
}

func (pr *PetRepository) Save(entity.Pet) error {
	return nil
}