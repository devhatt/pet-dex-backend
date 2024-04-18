package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
)

type OngUsecase struct {
	repo   interfaces.OngRepository
	hasher interfaces.Hasher
}

func NewOngUseCase(repo interfaces.OngRepository, hasher interfaces.Hasher) *OngUsecase {
	return &OngUsecase{
		repo:   repo,
		hasher: hasher,
	}
}

func (o *OngUsecase) Save(ong *entity.Ong) error {
	ongUser := entity.NewUser(ong.User.Name, ong.User.Type, ong.User.Document, ong.User.AvatarURL, ong.User.Email, ong.User.Phone, ong.User.Pass, ong.User.Adresses.City, ong.User.Adresses.State, ong.User.BirthDate)
	hashedPass, err := o.hasher.Hash(ong.User.Pass)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUsecase.Hash error: %w", err))
		return err
	}

	ongUser.Pass = hashedPass

	err = o.repo.SaveUser(ongUser)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.SaveUser error: %w", err))
		return err
	}

	err = o.repo.SaveAddress(&ongUser.Adresses)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.SaveAddress error: %w", err))
		return err
	}

	err = o.repo.Save(ong)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.Save error: %w", err))
		return err
	}

	return nil

}
