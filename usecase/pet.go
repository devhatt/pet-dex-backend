package usecase

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/interfaces"

	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type PetUseCase struct {
	repo interfaces.PetRepository
}

func NewPetUseCase(repo interfaces.PetRepository) *PetUseCase {
	return &PetUseCase{repo: repo}
}

func (c *PetUseCase) FindByID(ID uniqueEntityId.ID) (*entity.Pet, error) {
	pet, err := c.repo.FindByID(ID)
	if err != nil {
		err = fmt.Errorf("failed to retrieve pet: %w", err)
		return nil, err
	}
	return pet, nil
}

func (c *PetUseCase) Update(petID string, userID string, petUpdateDto dto.PetUpdatetDto) (err error) {
	petToUpdate := petUpdateDto.ToEntity()

	if !c.isValidPetSize(petToUpdate) {
		return errors.New("the animal size is invalid")
	}

	if !c.isValideSpecialCare(petToUpdate) {
		return errors.New("failed to update special care")
	}

	err = c.repo.Update(petID, userID, petToUpdate)
	if err != nil {
		return fmt.Errorf("failed to update for pet with ID %s: %w", petID, err)
	}

	return nil
}

func (c *PetUseCase) isValidPetSize(petToUpdate *entity.Pet) bool {
	return &petToUpdate.Size != nil && petToUpdate.Size != "" &&
		(petToUpdate.Size == "small" || petToUpdate.Size == "medium" || petToUpdate.Size == "large" || petToUpdate.Size == "giant")
}

func (c *PetUseCase) ListUserPets(userID uniqueEntityId.ID) ([]*entity.Pet, error) {
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

	if needed {
		return description != ""
	}
	if !needed {
		return description == ""
	}
	return false
}
