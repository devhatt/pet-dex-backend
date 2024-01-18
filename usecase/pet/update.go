package pet

import (
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

	updateValues := map[string]interface{}{}

	if &petToUpdate.Size != nil && petToUpdate.Size != "" {
		updateValues["size"] = &petToUpdate.Size
	}

	err = c.repo.Update(id, updateValues)

	if err != nil {
		fmt.Errorf("failed to update size for pet with ID %d: %w", id, err)
		return err
	}

	return
}

func (c *UpdateUseCase) isValidSize(size string) bool {
	return size == "Pequeno" || size == "Médio" || size == "Grande"
}
