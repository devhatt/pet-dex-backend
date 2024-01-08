package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type PetUseCase struct {
	repo interfaces.PetRepository
}

func NewPetUseCase(repo interfaces.PetRepository) (*PetUseCase) {
	return &PetUseCase{repo : repo}
}

func (c *PetUseCase) Find(id int) (*entity.Pet, error) {
	pet, err := c.repo.Find(id)
	if err != nil {
		fmt.Printf("failed")
		return nil, err
	}
	return pet, nil
}