package usecase

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type OngUseCase struct {
	repo interfaces.OngRepository
}

func NewOngUseCase(repo interfaces.OngRepository) *OngUseCase {
	return &OngUseCase{repo: repo}
}

func (c *OngUseCase) Save(ong entity.Ong) error {
	c.repo.Save(ong)
	return nil
}
