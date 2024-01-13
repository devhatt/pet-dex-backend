package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type PetUseCase struct {
	repo interfaces.PetRepository
}

func NewPetUseCase(repo interfaces.PetRepository) *PetUseCase {
	return &PetUseCase{repo: repo}
}

func (c *PetUseCase) FindById(id int) (*entity.Pet, error) {
	pet, err := c.repo.FindById(id)
	if err != nil {
		fmt.Printf("failed")
		return nil, err
	}
	return pet, nil
}

func (c *PetUseCase) FindNoAuthPets() {
	json, err := c.repo.FindNoAuthPets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json)
}
