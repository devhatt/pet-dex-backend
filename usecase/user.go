package usecase

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"time"

	"github.com/golang-jwt/jwt"
)

var loggerUser = config.GetLogger("user-usercase")

type UserUsecase struct {
	repo        interfaces.UserRepository
	hasher      interfaces.Hasher
	encoder     interfaces.Encoder
	addressRepo interfaces.AdressRepo
}

func NewUserUsecase(repo interfaces.UserRepository, hasher interfaces.Hasher, encoder interfaces.Encoder, addressRepo interfaces.AdressRepo) *UserUsecase {
	return &UserUsecase{
		repo:        repo,
		hasher:      hasher,
		encoder:     encoder,
		addressRepo: addressRepo,
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
		loggerUser.Error(fmt.Errorf("#UserUsecase.Update error: %w", err))
		return err
	}

	return nil

}

func (uc *UserUsecase) FindByID(ID uniqueEntityId.ID) (*entity.User, error) {
	user, err := uc.repo.FindByID(ID)

	if err != nil {
		loggerUser.Error(fmt.Errorf("#UserUsecase.Update error: %w", err))
		err = fmt.Errorf("failed to retrieve user: %w", err)
		return nil, err
	}

	address, err := uc.addressRepo.FindAddressByUserID(user.ID)

	if err != nil {
		loggerUser.Error(fmt.Errorf("#UserUsecase.Update error: %w", err))
		err = fmt.Errorf("failed to retrieve user: %w", err)
		return nil, err
	}

	user.Adresses = *address

	return user, nil
}
