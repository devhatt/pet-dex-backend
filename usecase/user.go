package usecase

import (
	"errors"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserUsecase struct {
	repo    interfaces.UserRepository
	hasher  interfaces.Hasher
	encoder interfaces.Encoder
	logger  config.Logger
}

func NewUserUsecase(repo interfaces.UserRepository, hasher interfaces.Hasher, encoder interfaces.Encoder) *UserUsecase {
	return &UserUsecase{
		repo:    repo,
		hasher:  hasher,
		encoder: encoder,
		logger:  *config.GetLogger("user-usecase"),
	}
}

func (uc *UserUsecase) Save(userDto dto.UserInsertDto) error {
	user := entity.NewUser(userDto.Name, userDto.Type, userDto.Document, userDto.AvatarURL, userDto.Email, userDto.Phone, userDto.Pass, userDto.City, userDto.State, userDto.BirthDate)

	hashedPass, err := uc.hasher.Hash(user.Pass)
	if err != nil {
		uc.logger.Error("error hashing: ", err)
		return err
	}

	user.Pass = hashedPass

	err = uc.repo.Save(user)
	if err != nil {
		uc.logger.Error("error saving user: ", err)
		return err
	}

	err = uc.repo.SaveAddress(&user.Adresses)
	if err != nil {
		uc.logger.Error("error saving user adress: ", err)
		return err
	}

	return nil

}

func (uc *UserUsecase) GenerateToken(loginDto *dto.UserLoginDto) (string, error) {
	user := uc.repo.FindByEmail(loginDto.Email)
	if user.Name == "" {
		return "", errors.New("invalid credentials")
	}
	if !uc.hasher.Compare(loginDto.Password, user.Pass) {
		return "", errors.New("invalid credentials")
	}
	token, _ := uc.encoder.NewAccessToken(interfaces.UserClaims{
		Id:    user.ID.String(),
		Name:  user.Email,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	})
	return token, nil
}

func (uc *UserUsecase) Update(userID uniqueEntityId.ID, userDto dto.UserUpdateDto) error {
	user := entity.UserToUpdate(&userDto)

	err := uc.repo.Update(userID, user)
	if err != nil {
		uc.logger.Error("error updating user: ", err)
		return err
	}

	return nil

}

func (uc *UserUsecase) FindByID(ID uniqueEntityId.ID) (*entity.User, error) {
	user, err := uc.repo.FindByID(ID)

	if err != nil {
		uc.logger.Error("error finding user by id:", err)
		return nil, err
	}

	address, err := uc.repo.FindAddressByUserID(user.ID)

	if err != nil {
		uc.logger.Error("error finding user address:", err)
		return nil, err
	}

	user.Adresses = *address

	return user, nil
}
