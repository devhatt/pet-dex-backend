package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngUseCase struct {
	repo interfaces.OngRepository
}

func (c *OngUseCase) FindByID(ID uniqueEntityId.ID) (*entity.Ong, error) {
	ong, err := c.repo.FindByID(ID)
	if err != nil {
		err = fmt.Errorf("Failed to retrieve ONG: %w ", err)
	}
	return ong, nil
}
