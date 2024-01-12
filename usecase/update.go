package usecase

import (
	"fmt"
	"pet-dex-backend/v2/interfaces"
)

type UpdateUseCase struct {
	repo interfaces.PetRepository
}

func NewUpdateUseCase(repo interfaces.PetRepository) *UpdateUseCase {
	return &UpdateUseCase{repo: repo}
}

func (c *UpdateUseCase) Do(id int, newSize string) (err error) {

	if !c.isValidSize(newSize) {
		fmt.Errorf("invalid size: %s", newSize)
		return err
	}

	err = c.repo.UpdateSize(id, newSize)

	if err != nil {
		fmt.Errorf("failed to update size for pet with ID %d: %w", id, err)
		return err
	}

	return
}

func (c *UpdateUseCase) isValidSize(size string) bool {
	return size == "Pequeno" || size == "Médio" || size == "Grande"
}
