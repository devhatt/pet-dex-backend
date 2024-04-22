package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
)

type OngUsecase struct {
	repo   interfaces.OngRepository
	hasher interfaces.Hasher
	logger config.Logger
}

func NewOngUseCase(repo interfaces.OngRepository, hasher interfaces.Hasher) *OngUsecase {
	return &OngUsecase{
		repo:   repo,
		hasher: hasher,
		logger: *config.NewLogger("ong-usecase"),
	}
}

func (o *OngUsecase) Save(ongDto *dto.OngInsertDto) error {
	ong := entity.NewOng(ongDto.Name, ongDto.Type, ongDto.Document, ongDto.AvatarURL, ongDto.Email, ongDto.Phone, ongDto.Pass, ongDto.City, ongDto.State, ongDto.OpeningHours, ongDto.AdoptionPolicy, ongDto.BirthDate, *ongDto.Links)
	hashedPass, err := o.hasher.Hash(ong.User.Pass)

	if err != nil {
		logger.Error("error on hashing: ", err)
		return err
	}

	ong.User.Pass = hashedPass

	err = o.repo.SaveUser(&ong.User)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.SaveUser error: %w", err))
		return err
	}

	err = o.repo.SaveAddress(&ong.User.Adresses)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.SaveAddress error: %w", err))
		return err
	}

	err = o.repo.Save(ong)

	if err != nil {
		logger.Error("error on Save: ", err)
		return err
	}

	return nil

}
