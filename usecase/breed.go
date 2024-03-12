package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type BreedUseCase struct {
	repo interfaces.BreedRepository
}

func NewBreedUseCase(repo interfaces.BreedRepository) *BreedUseCase {
	return &BreedUseCase{repo: repo}
}

func (useCase *BreedUseCase) List() ([]*entity.Breed, error) {
	fmt.Println("List breeds")
	breed, err := useCase.repo.List()
	if err != nil {
		fmt.Println("Error listing breeds")
		fmt.Errorf("error listing breeds: %w", err)
		return nil, err
	}
	return breed, nil

}
