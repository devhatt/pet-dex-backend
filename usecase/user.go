package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/interfaces"
)

type UserUsecase struct {
	repo   interfaces.UserRepository
	hasher interfaces.Hasher
}

func NewUserUsecase(repo interfaces.UserRepository, hasher interfaces.Hasher) *UserUsecase {
	return &UserUsecase{
		repo:   repo,
		hasher: hasher,
	}
}

func (uc *UserUsecase) Save(userDto dto.UserInsertDto) error {
	user := entity.NewUser(userDto.Name, userDto.Type, userDto.Document, userDto.AvatarURL, userDto.Email, userDto.Phone, userDto.Pass, userDto.City, userDto.State, userDto.BirthDate)
	hashedPass, err := uc.hasher.Hash(user.Pass)

	if err != nil {
		fmt.Println(fmt.Errorf("#UserUsecase.Hash error: %w", err))
		return err
	}

	user.Pass = hashedPass

	err = uc.repo.Save(user)

	if err != nil {
		fmt.Println(fmt.Errorf("#UserUsecase.Save error: %w", err))
		return err
	}

	err = uc.repo.SaveAddress(&user.Adresses)

	if err != nil {
		fmt.Println(fmt.Errorf("#UserUsecase.SaveAddress error: %w", err))
		return err
	}

	return nil

}
