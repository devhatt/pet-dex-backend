package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

var loggerBreed = config.GetLogger("breed-usecase")

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
		loggerBreed.Error("error listing breeds", err)
		err = fmt.Errorf("error listing breeds: %w", err)
		return nil, err
	}
	return breed, nil
}

func (useCase *BreedUseCase) FindByID(ID uniqueEntityId.ID) (*entity.Breed, error) {
	breed, err := useCase.repo.FindByID(ID)
	if err != nil {
		err = fmt.Errorf("failed to retrieve breed: %w", err)
		return nil, err
	}
	return breed, nil
}