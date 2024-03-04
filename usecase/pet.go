package usecase

import (
	"errors"
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
	return nil, nil
}

func (c *PetUseCase) Update(petID string, userID string, petToUpdate *entity.Pet) (err error) {

	updateValues := map[string]interface{}{}

	if c.isValidPetSize(petToUpdate) {
		updateValues["size"] = &petToUpdate.Size
	} else {
		return errors.New("The animal size is invalid")
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
