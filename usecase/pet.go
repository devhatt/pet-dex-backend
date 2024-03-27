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

func (c *PetUseCase) Update(petID string, userID string, petToUpdate *entity.Pet) (err error) {

	updateValues := map[string]interface{}{}

	if c.isValidPetSize(petToUpdate) {
		updateValues["size"] = &petToUpdate.Size
	} else {
		return errors.New("the animal size is invalid")
	}

	if c.isValideSpecialCare(petToUpdate) {
		updateValues["cuidados_especiais"] = &petToUpdate.NeedSpecialCare
	} else {
		return errors.New("Failled to update special care")
	}

	err = c.repo.Update(petID, userID, updateValues)
	if err != nil {
		return fmt.Errorf("failed to update for pet with ID %s: %w", petID, err)
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

func (c *PetUseCase) isValideSpecialCare(petToUpdate *entity.Pet) bool {
	var needed = petToUpdate.NeedSpecialCare.Needed
	var description = petToUpdate.NeedSpecialCare.Description

	if needed == true {
		if &description != nil || description != "" {
			return true
		}
	}
	if needed == false {
		if &description == nil || description == "" {
			return true
		}
	}
	return false
}
