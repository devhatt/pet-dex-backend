package usecase

import (
	"pet-dex-backend/v2/entity"
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

func (uc *UserUsecase) Save(user entity.User) error {
	//TODO: mandar o create de User para o controller recebendo dele um DTO
	//TODO: talvez melhorar os tratamentos de erros aqui
	//TODO: Fazer testes
	hashedPass, err := uc.hasher.Hash(user.Pass)

	if err != nil {
		return err
	}

	user.Pass = hashedPass

	err = uc.repo.Save(user)

	if err != nil {
		return err
	}

	err = uc.repo.SaveAddress(user)

	if err != nil {
		return err
	}

	return nil

}
