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

func (c *UpdateUseCase) Do(id string, userID string, petToUpdate *entity.Pet) (err error) {
	// Validação de ID e usuário
	//petIsFound, err := c.repo.FindById(id)
	//if err != nil {
	//	return fmt.Errorf("falha ao recuperar o animal de estimação com o ID %s: %w", id, err)
	//}
	//
	//if petIsFound == nil {
	//	return fmt.Errorf("animal de estimação com o ID %s não encontrado", id)
	//}
	//
	//if petIsFound.UserID != petToUpdate.UserID {
	//	return fmt.Errorf("não autorizado para atualizar o animal de estimação com o ID %s", id)
	//}

	// Atributos a serem atualizados
	updateValues := map[string]interface{}{}

	if &petToUpdate.Size != nil && petToUpdate.Size != "" && c.IsValidSize(petToUpdate.Size) {
		updateValues["size"] = &petToUpdate.Size
	} else {
		return errors.New("Tamanho do animal de estimação é inválido")
	}

	// Atualização do pet
	err = c.repo.Update(id, updateValues)
	if err != nil {
		return fmt.Errorf("falha ao atualizar o tamanho do animal de estimação com o ID %s: %w", id, err)
	}

	return nil
}

func (c *UpdateUseCase) IsValidSize(size string) bool {
	return size == "small" || size == "medium" || size == "large" || size == "giant"
}
