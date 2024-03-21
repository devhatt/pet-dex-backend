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

func (c *PetUseCase) FindByID(ID uniqueEntity.ID) (*entity.Pet, error) {
	pet, err := c.repo.FindByID(ID)
	if err != nil {
		err = fmt.Errorf("failed to retrieve pet: %w", err)
		return nil, err
	}
	return pet, nil
}

func (c *PetUseCase) Update(petID string, userID string, petToUpdate *entity.Pet) (err error) {

	updateValues := map[string]interface{}{}

	if c.isValidPetSize(petToUpdate) {
		updateValues["size"] = &petToUpdate.Size
	} else {
		return errors.New("the animal size is invalid")
	}

	err = c.repo.Update(petID, userID, updateValues)
	if err != nil {
		return fmt.Errorf("failed to update size for pet with ID %s: %w", petID, err)
	}

	return nil
}

func (c *PetUseCase) isValidPetSize(petToUpdate *entity.Pet) bool {
	return &petToUpdate.Size != nil && petToUpdate.Size != "" &&
		(petToUpdate.Size == "small" || petToUpdate.Size == "medium" || petToUpdate.Size == "large" || petToUpdate.Size == "giant")
}

func (c *PetUseCase) ListUserPets(userID uniqueEntity.ID) ([]*entity.Pet, error) {
	pets, err := c.repo.ListByUser(userID)
	if err != nil {
		err = fmt.Errorf("failed to retrieve all user pets: %w", err)
		return nil, err
	}
	return pets, nil
}
