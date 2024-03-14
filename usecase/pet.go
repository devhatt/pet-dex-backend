package usecase

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"

	uniqueEntity "pet-dex-backend/v2/pkg/entity"
)

type PetUseCase struct {
	repo interfaces.PetRepository
}

func NewPetUseCase(repo interfaces.PetRepository) *PetUseCase {
	return &PetUseCase{repo: repo}
}

func (c *PetUseCase) FindById(userID uniqueEntity.ID) (*entity.Pet, error) {
	return nil, nil
}

func (petUseCase *PetUseCase) Update(petID string, userID string, petToUpdate *entity.Pet) (err error) {

	updateValues := map[string]interface{}{}

	if petUseCase.isValidPetSize(petToUpdate) {
		updateValues["size"] = petToUpdate.Size
	} else {
		return errors.New("the animal size is invalid")
	}

	if petUseCase.isValidWeight(petToUpdate) {
		updateValues["weight"] = petToUpdate.Weight
		updateValues["weight_measure"] = petToUpdate.WeightMeasure
	} else {
		return errors.New("the animal weight is invalid")
	}

	if len(updateValues) == 0 {
		return errors.New("invalid update payload")
	}

	err = petUseCase.repo.Update(petID, userID, updateValues)
	if err != nil {
		return fmt.Errorf("failed to update size for pet with ID %s: %w", petID, err)
	}

	return nil
}

func (c *PetUseCase) isValidPetSize(petToUpdate *entity.Pet) bool {
	return &petToUpdate.Size != nil && petToUpdate.Size != "" &&
		(petToUpdate.Size == "small" || petToUpdate.Size == "medium" || petToUpdate.Size == "large" || petToUpdate.Size == "giant")
}

func (c *PetUseCase) isValidWeight(petToUpdate *entity.Pet) bool {
	return petToUpdate.Weight > 0 &&
		(petToUpdate.WeightMeasure == "kg" || petToUpdate.WeightMeasure == "lb")
}

func (c *PetUseCase) ListUserPets(userID uniqueEntity.ID) ([]*entity.Pet, error) {
	pets, err := c.repo.ListByUser(userID)
	if err != nil {
		err = fmt.Errorf("failed to retrieve all user pets: %w", err)
		return nil, err
	}
	return pets, nil
}
