package pet

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type UpdateUseCase struct {
	repo interfaces.PetRepository
}

func NewUpdateUseCase(repo interfaces.PetRepository) *UpdateUseCase {
	return &UpdateUseCase{repo: repo}
}

func (c *UpdateUseCase) Do(id string, petToUpdate *entity.Pet) (err error) {

	petIsFound, err := c.repo.FindById(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve pet with ID %s: %w", id, err)
	}
	if petIsFound == nil {
		return fmt.Errorf("pet with ID %s not found", id)
	}

	if petIsFound.UserID != petToUpdate.UserID {
		return fmt.Errorf("unauthorized to update pet with ID %s", id)
	}

	updateValues := map[string]interface{}{}

	if &petToUpdate.Size != nil && petToUpdate.Size != "" && c.isValidSize(petToUpdate.Size) {
		updateValues["size"] = &petToUpdate.Size
	} else {
		return errors.New("Pet size is invalid")
	}

	err = c.repo.Update(id, updateValues)

	if err != nil {
		fmt.Errorf("failed to update size for pet with ID %s: %w", id, err)
		return err
	}

	return
}

func (c *UpdateUseCase) isValidSize(size string) bool {
	return size == "small" || size == "medium" || size == "large" || size == "giant"
}
