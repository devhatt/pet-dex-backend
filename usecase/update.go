package usecase

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type UpdateUseCase struct {
	petRepository interfaces.PetRepository
}

func NewUpdateUseCase(p interfaces.PetRepository) *UpdateUseCase {
	return &UpdateUseCase{
		petRepository: p,
	}
}

func (auc *UpdateUseCase) Do(petID string, userID string, petToUpdate *entity.Pet) error {
	// petIsFound, err := auc.petRepository.FindById(petID)

	// if err != nil {
	// 	return fmt.Errorf("failed to retrieve pet with ID %d: %v", petID, err)
	// }

	// if petIsFound.UserID != petToUpdate.UserID {
	// 	return fmt.Errorf("unauthorized to update pet with ID %d", petID)
	// }

	updateValues := map[string]interface{}{}
	if petToUpdate.Weight > 0.0 {
		updateValues["weight"] = petToUpdate.Weight
	} else {
		return errors.New("Pet size is invalid")
	}

	err := auc.petRepository.Update(petID, userID, updateValues)

	if err != nil {
		return fmt.Errorf("failed to update size for pet with ID %s: %w", petID, err)
	}

	return nil
}
