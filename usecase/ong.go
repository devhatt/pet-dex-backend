package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)

type OngUsecase struct {
	repo     interfaces.OngRepository
	userRepo interfaces.UserRepository
	hasher   interfaces.Hasher
	logger   config.Logger
}

func NewOngUseCase(repo interfaces.OngRepository, userRepo interfaces.UserRepository, hasher interfaces.Hasher) *OngUsecase {
	return &OngUsecase{
		repo:     repo,
		userRepo: userRepo,
		hasher:   hasher,
		logger:   *config.NewLogger("ong-usecase"),
	}
}

func (o *OngUsecase) Save(ongDto *dto.OngInsertDto) error {
	ong := entity.NewOng(*ongDto)
	hashedPass, err := o.hasher.Hash(ong.User.Pass)

	if err != nil {
		o.logger.Error("error on ong usecase: ", err)
		return err
	}

	ong.User.Pass = hashedPass

	err = o.userRepo.Save(&ong.User)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.SaveUser error: %w", err))
		return err
	}

	err = o.userRepo.SaveAddress(&ong.User.Adresses)

	if err != nil {
		fmt.Println(fmt.Errorf("#OngUseCase.SaveAddress error: %w", err))
		return err
	}

	err = o.repo.Save(ong)

	if err != nil {
		o.logger.Error("error on ong Save: ", err)
		return err
	}

	return nil

}

func (o *OngUsecase) FindByID(ID uniqueEntityId.ID) (*entity.Ong, error) {
	// Buscar a ONG pelo ID no repositório.
	ong, err := o.repo.FindByID(ID)
	if err != nil {
		o.logger.Error("Error find Ong by ID: ", err)
		return nil, fmt.Errorf("Error find Ong by ID: %w", err)
	}
	return ong, nil
}
