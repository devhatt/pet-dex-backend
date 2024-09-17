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

func (o *OngUsecase) List(limit, offset int, sortBy, order string) ([]*dto.OngListMapper, error) {
	ong, err := o.repo.List(limit, offset, sortBy, order)

	if err != nil {
		err = fmt.Errorf("error listing ongs: %w", err)
		return nil, err
	}
	return ong, nil
}

func (c *OngUsecase) FindByID(ID uniqueEntityId.ID) (*dto.OngListMapper, error) {

	ong, err := c.repo.FindByID(ID)

	if err != nil {
		c.logger.Error("error on ong repository: ", err)
		err = fmt.Errorf("failed to retrieve ong: %w", err)
		return nil, err
	}

	return ong, nil
}

func (o *OngUsecase) Update(ongId uniqueEntityId.ID, ongDto *dto.OngUpdateDto) error {
	ongToUpdate, _ := entity.OngToUpdate(*ongDto)

	ong, err := o.repo.FindByID(ongId)
	if err != nil {
		o.logger.Error("error on ong usecase: ", err)
		return err
	}

	err = o.userRepo.Update(ong.UserID, ongToUpdate.User)
	if err != nil {
		o.logger.Error("error on ong usecase: ", err)
		return err
	}

	err = o.repo.Update(ongId, *ongToUpdate)
	if err != nil {
		o.logger.Error("error on ong usecase: ", err)
		return err
	}

	return nil

}

func (o *OngUsecase) Delete(ongId uniqueEntityId.ID) error {
	_, err := o.repo.FindByID(ongId)
	if err != nil {
		o.logger.Error("error on ong usecase: ", err)
		return err
	}
	err = o.repo.Delete(ongId)
	if err != nil {
		o.logger.Error("error on ong usecase: ", err)
		return err
	}
	return nil
}
