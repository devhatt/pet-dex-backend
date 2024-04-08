package usecase

import (
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngUseCase struct {
	repo interfaces.OngRepository
}

func (c *OngUseCase) FindByID(ID uniqueEntityId.ID) *entity.ong
