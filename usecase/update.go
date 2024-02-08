package usecase

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

func (c *UpdateUseCase) Do(petID string, userID string, petToUpdate *entity.Pet) (err error) {
	//petIsFound, err := c.repo.FindById(petIDid)
	//if err != nil {
	//	return fmt.Errorf("falha ao recuperar o animal de estimação com o ID %s: %w", petIDid, err)
	//}
	//if petIsFound.UserID != userID {
	//	return fmt.Errorf("não autorizado para atualizar o animal de estimação com o ID %s", petIDid)
	//}

	updateValues := map[string]interface{}{}

	if &petToUpdate.Size != nil && petToUpdate.Size != "" && c.IsValidSize(petToUpdate.Size) {
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

func (c *UpdateUseCase) IsValidSize(size string) bool {
	return size == "small" || size == "medium" || size == "large" || size == "giant"
}
