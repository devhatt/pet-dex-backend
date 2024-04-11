package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type OngUsecase struct {
	repo interfaces.OngRepository
}

func NewOngUseCase(repo interfaces.OngRepository) *OngUsecase {
	return &OngUsecase{repo: repo}
}

func (o *OngUsecase) Save(ong *entity.Ong) error {
	err := o.repo.Save(ong)

	if err != nil {
		fmt.Println(fmt.Errorf("error saving ong: %w", err))
		return err
	}

	return nil
}
