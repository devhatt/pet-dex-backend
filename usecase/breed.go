package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/interfaces"
)

type BreedUseCase struct {
	repo interfaces.BreedRepository
}

func NewBreedUseCase(repo interfaces.BreedRepository) *BreedUseCase {
	return &BreedUseCase{
		repo: repo,
	}
}

func (useCase *BreedUseCase) List() ([]*dto.BreedList, error) {
	breed, err := useCase.repo.List()
	if err != nil {
		err = fmt.Errorf("error listing breeds: %w", err)
		return nil, err
	}
	return breed, nil
}
