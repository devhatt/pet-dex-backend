package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type BreedUseCase struct {
	repo   interfaces.BreedRepository
	logger config.Logger
}

func NewBreedUseCase(repo interfaces.BreedRepository) *BreedUseCase {
	return &BreedUseCase{
		repo:   repo,
		logger: *config.GetLogger("breed-usecase"),
	}
}

func (useCase *BreedUseCase) List() ([]*dto.BreedList, error) {
	breed, err := useCase.repo.List()
	if err != nil {
		useCase.logger.Error("error listing breeds: ", err)
		err = fmt.Errorf("error listing breeds: %w", err)
		return nil, err
	}
	return breed, nil
}

func (useCase *BreedUseCase) FindByID(ID uniqueEntityId.ID) (*entity.Breed, error) {
	breed, err := useCase.repo.FindByID(ID)
	if err != nil {
		useCase.logger.Error("failed to retrieve breed: ", err)
		err = fmt.Errorf("failed to retrieve breed: %w", err)
		return nil, err
	}
	return breed, nil
}
