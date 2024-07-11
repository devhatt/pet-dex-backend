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

type UserUsecase struct {
	repo        interfaces.UserRepository
	hasher      interfaces.Hasher
	encoder     interfaces.Encoder
	logger      config.Logger
	ssoProvider interfaces.SingleSignOnProvider
}

func NewUserUsecase(repo interfaces.UserRepository, hasher interfaces.Hasher, encoder interfaces.Encoder, ssoProvider interfaces.SingleSignOnProvider) *UserUsecase {
	return &UserUsecase{
		repo:        repo,
		hasher:      hasher,
		encoder:     encoder,
		logger:      *config.GetLogger("user-usecase"),
		ssoProvider: ssoProvider,
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

func (uc *UserUsecase) Login(loginDto *dto.UserLoginDto) (string, error) {
	user, err := uc.FindByEmail(loginDto.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if user.Name == "" {
		return "", errors.New("invalid credentials")
	}
	if !uc.hasher.Compare(loginDto.Password, user.Pass) {
		return "", errors.New("invalid credentials")
	}
	token, _ := uc.encoder.NewAccessToken(interfaces.UserClaims{
		Id:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	})
	return token, nil
}

func (uc *UserUsecase) Update(userID uniqueEntityId.ID, userDto dto.UserUpdateDto) error {
	user := entity.UserToUpdate(userDto)

	err := uc.repo.Update(userID, user)
	if err != nil {
		uc.logger.Error("error updating user: ", err)
		return err
	}

	return nil
}

func (uc *UserUsecase) FindByEmail(email string) (*entity.User, error) {
	user, err := uc.repo.FindByEmail(email)

	if err != nil {
		uc.logger.Error("error finding user by email:", err)
		return nil, err
	}

	return user, nil
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

func (uc *UserUsecase) Delete(userID uniqueEntityId.ID) error {
	err := uc.repo.Delete(userID)

	if err != nil {
		uc.logger.Error(fmt.Errorf("#UserUsecase.Delete error: %w", err))
		return err
	}

	return nil
}

func (uc *UserUsecase) ChangePassword(userChangePasswordDto dto.UserChangePasswordDto, userId uniqueEntityId.ID) error {
	user, err := uc.repo.FindByID(userId)
	if err != nil {
		uc.logger.Error("error finding user by id: ", err)
		return errors.New("") // Don't show the user the reason for secure reasons
	}

	if !uc.hasher.Compare(userChangePasswordDto.OldPassword, user.Pass) {
		uc.logger.Error("old password does not match")
		return errors.New("old password does not match")
	}

	newPassword, err := uc.hasher.Hash(userChangePasswordDto.NewPassword)
	if err != nil {
		uc.logger.Error("error hashing: ", err)
		return err
	}
	err = uc.repo.ChangePassword(userId, newPassword)
	if err != nil {
		uc.logger.Error(err)
		return err
	}
	return nil
}

func (uc *UserUsecase) UpdatePushNotificationSettings(userID uniqueEntityId.ID, userPushNotificationEnabled dto.UserPushNotificationEnabled) error {
	user, err := uc.repo.FindByID(userID)

	if err != nil {
		uc.logger.Error("error finding user by id: ", err)
		return errors.New("user dont exists")
	}

	user.PushNotificationsEnabled = &userPushNotificationEnabled.PushNotificationEnabled

	err = uc.repo.Update(userID, *user)

	if err != nil {
		uc.logger.Error("error updating user by id: ", err)
		return errors.New("error on updating push notification")
	}

	return nil

}

func (uc *UserUsecase) ProviderLogin(accessToken string, provider string) (*entity.User, bool, error) {
	userInfo, err := uc.ssoProvider.GetUserDetails(provider, accessToken)
	if err != nil {
		return nil, false, err
	}

	user, _ := uc.FindByEmail(userInfo.Email)

	return user, user == nil, nil

}

func (uc *UserUsecase) NewAccessToken(id string, name string, email string) (string, error) {
	return uc.encoder.NewAccessToken(interfaces.UserClaims{
		Id:    id,
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	})
}
