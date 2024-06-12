package usecase

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type PetUseCase struct {
	repo   interfaces.PetRepository
	logger config.Logger
}

func NewPetUseCase(repo interfaces.PetRepository) *PetUseCase {
	return &PetUseCase{
		repo:   repo,
		logger: *config.GetLogger("pet-usecase"),
	}
}

func (c *PetUseCase) FindByID(ID uniqueEntityId.ID) (*entity.Pet, error) {
	pet, err := c.repo.FindByID(ID)
	if err != nil {
		c.logger.Error("failed to retrieve pet: ", err)
		return nil, err
	}
	return pet, nil
}

func (c *PetUseCase) Update(petID string, userID string, petUpdateDto dto.PetUpdateDto) (err error) {
	petToUpdate := entity.PetToEntity(&petUpdateDto)

	if !c.isValidPetSize(petToUpdate) {
		return errors.New("the animal size is invalid")
	}

	if !c.isValidSpecialCare(petToUpdate) {
		return errors.New("failed to update special care")
	}

	if !c.isValidWeight(petToUpdate) {
		return errors.New("the animal weight is invalid")
	}

	err = c.repo.Update(petID, userID, petToUpdate)
	if err != nil {
		c.logger.Error("error updating pet: ", err)
		return fmt.Errorf("failed to update pet with ID %s: %w", petID, err)
	}

	return nil
}

func (c *PetUseCase) isValidPetSize(petToUpdate *entity.Pet) bool {
	return (petToUpdate.Size == "small" || petToUpdate.Size == "medium" || petToUpdate.Size == "large" || petToUpdate.Size == "giant")
}

func (c *PetUseCase) isValidWeight(petToUpdate *entity.Pet) bool {
	return (petToUpdate.Weight > 0 &&
		(petToUpdate.WeightMeasure == "kg" || petToUpdate.WeightMeasure == "lb"))
}

func (c *PetUseCase) ListUserPets(userID uniqueEntityId.ID) ([]*entity.Pet, error) {
	pets, err := c.repo.ListByUser(userID)
	if err != nil {
		c.logger.Error("failed to retrieve all user pets: ", err)
		return nil, err
	}
	return pets, nil
}

func (c *PetUseCase) isValidSpecialCare(petToUpdate *entity.Pet) bool {
	var needed = petToUpdate.NeedSpecialCare.Needed
	var description = petToUpdate.NeedSpecialCare.Description

	if needed != nil {
		if *needed {
			return description != ""
		}
		if !*needed {
			return description == ""
		}
	}
	return true
}

func (c *PetUseCase) Save(petDto dto.PetInsertDto) error {
	pet := entity.NewPet(petDto.UserID, petDto.BreedID, petDto.Size, petDto.Name, petDto.Weight, petDto.AdoptionDate, petDto.Birthdate)

	err := c.repo.Save(*pet)
	if err != nil {
		c.logger.Error("failed to save pet: ", err)
		return err
	}
	return nil
}
