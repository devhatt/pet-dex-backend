package usecase

import (
	"fmt"
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
	err := c.repo.Save(ong)
	if err != nil {
		fmt.Errorf("error saving ong: %w", err)
		return err
	}
	return nil
}
