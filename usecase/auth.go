package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type AuthUseCase struct {
	repo interfaces.UserClaims
}

func NewAuthUseCase(repo interfaces.PetRepository) (*PetUseCase) {
	return &PetUseCase{repo : repo}
}

func (c *PetUseCase) FindById(id int) (*entity.Pet, error) {
	pet, err := c.repo.FindById(id)
	if err != nil {
		fmt.Printf("failed")
		return nil, err
	}
	return pet, nil
}