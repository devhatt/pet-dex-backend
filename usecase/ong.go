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
	user := entity.NewUser(ong.Name, ong.Type, ong.Document, ong.AvatarURL, ong.Email, ong.Phone, ong.Pass, ong.Address.City, ong.Address.State, ong.CreationDate)
	hashedPass, err := o.hasher.Hash(ong.Pass)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUsecase.Hash error: %w", err))
		return err
	}

	user.Pass = hashedPass

	err = o.repo.SaveUser(user)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.SaveUser error: %w", err))
		return err
	}

	err = o.repo.SaveAddress(&user.Adresses)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.SaveAddress error: %w", err))
		return err
	}

	err = o.repo.Save(ong, user.ID)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.Save error: %w", err))
		return err
	}

	return nil

}
