package usecase

import (
	"errors"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/interfaces"
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/hasher"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserUseCase(t *testing.T) {
	tcases := map[string]struct {
		repo         interfaces.UserRepository
		hasher       interfaces.Hasher
		expectOutput *UserUsecase
		encoder      interfaces.Encoder
	}{
		"success": {
			repo:         mockInterfaces.NewMockUserRepository(t),
			hasher:       mockInterfaces.NewMockHasher(t),
			expectOutput: &UserUsecase{},
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			usecase := NewUserUsecase(tcase.repo, tcase.hasher, tcase.encoder)

			assert.IsTypef(t, tcase.expectOutput, usecase, "error: New Hasher not returns a *Hasher{} struct", nil)
		})
	}
}

func TestSave(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		hasher       *mockInterfaces.MockHasher
		input        dto.UserInsertDto
		encoder      interfaces.Encoder
		expectOutput error
	}{
		"success": {
			repo:   mockInterfaces.NewMockUserRepository(t),
			hasher: mockInterfaces.NewMockHasher(t),
			input: dto.UserInsertDto{
				Name:      "teste",
				Type:      "teste",
				Document:  "teste",
				AvatarURL: "teste",
				Email:     "teste",
				Phone:     "teste",
				Pass:      "hashedPass",
				BirthDate: &time.Time{},
				City:      "teste",
				State:     "teste",
			},
			expectOutput: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.hasher.On("Hash", tcase.input.Pass).Return("hashedPass", tcase.expectOutput)
			tcase.repo.On("Save", mock.Anything).Return(tcase.expectOutput)
			tcase.repo.On("SaveAddress", mock.Anything).Return(tcase.expectOutput)

			usecase := NewUserUsecase(tcase.repo, tcase.hasher, tcase.encoder)
			err := usecase.Save(tcase.input)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestErrorSave(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		hasher       *mockInterfaces.MockHasher
		input        dto.UserInsertDto
		encoder      interfaces.Encoder
		expectOutput error
	}{
		"errorSave": {
			repo:   mockInterfaces.NewMockUserRepository(t),
			hasher: mockInterfaces.NewMockHasher(t),
			input: dto.UserInsertDto{
				Name:      "teste",
				Type:      "teste",
				Document:  "teste",
				AvatarURL: "teste",
				Email:     "teste",
				Phone:     "teste",
				Pass:      "hashedPass",
				BirthDate: &time.Time{},
				City:      "teste",
				State:     "teste",
			},
			expectOutput: fmt.Errorf("error on save"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.hasher.On("Hash", tcase.input.Pass).Return("hashedPass", nil)
			tcase.repo.On("Save", mock.Anything).Return(tcase.expectOutput)

			usecase := NewUserUsecase(tcase.repo, tcase.hasher, tcase.encoder)
			err := usecase.Save(tcase.input)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestErrorHash(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		hasher       *mockInterfaces.MockHasher
		input        dto.UserInsertDto
		encoder      interfaces.Encoder
		expectOutput error
	}{
		"errorHash": {
			repo:   mockInterfaces.NewMockUserRepository(t),
			hasher: mockInterfaces.NewMockHasher(t),
			input: dto.UserInsertDto{
				Name:      "teste",
				Type:      "teste",
				Document:  "teste",
				AvatarURL: "teste",
				Email:     "teste",
				Phone:     "teste",
				Pass:      "hashedPass",
				BirthDate: &time.Time{},
				City:      "teste",
				State:     "teste",
			},
			expectOutput: fmt.Errorf("error on hash"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.hasher.On("Hash", tcase.input.Pass).Return("hashedPass", tcase.expectOutput)

			usecase := NewUserUsecase(tcase.repo, tcase.hasher, tcase.encoder)
			err := usecase.Save(tcase.input)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestErrorSaveAddress(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		hasher       *mockInterfaces.MockHasher
		input        dto.UserInsertDto
		encoder      interfaces.Encoder
		expectOutput error
	}{
		"errorSaveAddress": {
			repo:   mockInterfaces.NewMockUserRepository(t),
			hasher: mockInterfaces.NewMockHasher(t),
			input: dto.UserInsertDto{
				Name:      "teste",
				Type:      "teste",
				Document:  "teste",
				AvatarURL: "teste",
				Email:     "teste",
				Phone:     "teste",
				Pass:      "hashedPass",
				BirthDate: &time.Time{},
				City:      "teste",
				State:     "teste",
			},
			expectOutput: fmt.Errorf("error on save addresse"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.hasher.On("Hash", tcase.input.Pass).Return("hashedPass", nil)
			tcase.repo.On("Save", mock.Anything).Return(nil)
			tcase.repo.On("SaveAddress", mock.Anything).Return(tcase.expectOutput)

			usecase := NewUserUsecase(tcase.repo, tcase.hasher, tcase.encoder)
			err := usecase.Save(tcase.input)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestUpdate(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		hasher       *mockInterfaces.MockHasher
		inputID      uniqueEntityId.ID
		inputDto     dto.UserUpdateDto
		encoder      interfaces.Encoder
		expectOutput error
	}{
		"success": {
			repo:    mockInterfaces.NewMockUserRepository(t),
			hasher:  mockInterfaces.NewMockHasher(t),
			inputID: uniqueEntityId.NewID(),
			inputDto: dto.UserUpdateDto{
				Name: "teste",
			},
			expectOutput: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Update", mock.Anything, mock.Anything).Return(tcase.expectOutput)
			usecase := NewUserUsecase(tcase.repo, tcase.hasher, tcase.encoder)
			err := usecase.Update(tcase.inputID, tcase.inputDto)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestErrorUpdate(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		hasher       *mockInterfaces.MockHasher
		inputID      uniqueEntityId.ID
		inputDto     dto.UserUpdateDto
		encoder      interfaces.Encoder
		expectOutput error
	}{
		"errorSave": {
			repo:    mockInterfaces.NewMockUserRepository(t),
			hasher:  mockInterfaces.NewMockHasher(t),
			inputID: uniqueEntityId.NewID(),
			inputDto: dto.UserUpdateDto{
				Name: "teste",
			},
			expectOutput: fmt.Errorf("error on update user"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Update", mock.Anything, mock.Anything).Return(tcase.expectOutput)

			usecase := NewUserUsecase(tcase.repo, tcase.hasher, tcase.encoder)
			err := usecase.Update(tcase.inputID, tcase.inputDto)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestDelete(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		inputID      uniqueEntityId.ID
		expectOutput error
	}{
		"success": {
			repo:         mockInterfaces.NewMockUserRepository(t),
			inputID:      uniqueEntityId.NewID(),
			expectOutput: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Delete", tcase.inputID).Return(tcase.expectOutput)

			usecase := NewUserUsecase(tcase.repo, nil, nil)
			err := usecase.Delete(tcase.inputID)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestErrorDelete(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		inputID      uniqueEntityId.ID
		expectOutput error
	}{
		"errorDelete": {
			repo:         mockInterfaces.NewMockUserRepository(t),
			inputID:      uniqueEntityId.NewID(),
			expectOutput: fmt.Errorf("error on delete user"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Delete", tcase.inputID).Return(tcase.expectOutput)

			usecase := NewUserUsecase(tcase.repo, nil, nil)
			err := usecase.Delete(tcase.inputID)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestChangePassword(t *testing.T) {
	hash := hasher.NewHasher()
	oldHashPassword, _ := hash.Hash("oldPassword123!")
	userId := uniqueEntityId.NewID()
	mockedRepo := mockInterfaces.NewMockUserRepository(t)
	mockedHasher := mockInterfaces.NewMockHasher(t)
	tcases := map[string]struct {
		inputDto                   dto.UserChangePasswordDto
		encoder                    interfaces.Encoder
		expectedCompareReturn      bool
		expectOutputFindById       *entity.User
		expectOutputChangePassword error
	}{
		"success": {
			inputDto: dto.UserChangePasswordDto{
				OldPassword:      "oldPassword123!",
				NewPassword:      "NewPassword123!",
				NewPasswordAgain: "NewPassword123!",
			},
			expectedCompareReturn: true,
			expectOutputFindById: &entity.User{
				ID:   userId,
				Pass: oldHashPassword,
			},
			expectOutputChangePassword: nil,
		},
		"Wrong old Password": {
			expectedCompareReturn: false,
			inputDto: dto.UserChangePasswordDto{
				OldPassword:      "wrongOldPassword123!",
				NewPassword:      "NewPassword123!",
				NewPasswordAgain: "NewPassword123!",
			},
			expectOutputFindById: &entity.User{
				ID:   userId,
				Pass: oldHashPassword,
			},
			expectOutputChangePassword: errors.New("old password does not match"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			newHashPassword, _ := hash.Hash(tcase.inputDto.NewPassword)
			mockedHasher.On("Compare", tcase.inputDto.OldPassword, tcase.expectOutputFindById.Pass).Return(tcase.expectedCompareReturn)
			mockedHasher.On("Hash", tcase.inputDto.NewPassword).Return(newHashPassword, nil)
			mockedRepo.On("FindByID", userId).Return(tcase.expectOutputFindById, nil)
			mockedRepo.On("ChangePassword", mock.Anything, mock.Anything).Return(tcase.expectOutputChangePassword)
			usecase := NewUserUsecase(mockedRepo, mockedHasher, tcase.encoder)
			err := usecase.ChangePassword(tcase.inputDto, userId)

			assert.Equal(t, tcase.expectOutputChangePassword, err, "expected error mismatch")
		})
	}
}
