package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/interfaces"
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"
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
	}{
		"success": {
			repo:         mockInterfaces.NewMockUserRepository(t),
			hasher:       mockInterfaces.NewMockHasher(t),
			expectOutput: &UserUsecase{},
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			usecase := NewUserUsecase(tcase.repo, tcase.hasher)

			assert.IsTypef(t, tcase.expectOutput, usecase, "error: New Hasher not returns a *Hasher{} struct", nil)
		})
	}
}

func TestSave(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockUserRepository
		hasher       *mockInterfaces.MockHasher
		input        dto.UserInsertDto
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

			usecase := NewUserUsecase(tcase.repo, tcase.hasher)
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

			usecase := NewUserUsecase(tcase.repo, tcase.hasher)
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

			usecase := NewUserUsecase(tcase.repo, tcase.hasher)
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

			usecase := NewUserUsecase(tcase.repo, tcase.hasher)
			err := usecase.Save(tcase.input)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}
