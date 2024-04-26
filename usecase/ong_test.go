package usecase

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOngRepository struct {
	mock.Mock
}

func (m *MockOngRepository) FindByID(ID uniqueEntityId.ID) (*entity.Ong, error) {
	args := m.Called(ID)
	return args.Get(0).(*entity.Ong), args.Error(1)
}

func TestFindOngByID(t *testing.T) {
	tcases := map[string]struct {
		repo           interfaces.OngRepository
		userRepo       interfaces.UserRepository
		hasher         interfaces.Hasher
		expectOutput   *OngUsecase
		expectedErrMsg string
	}{
		"success": {
			repo:           mockInterfaces.NewMockOngRepository(t),
			userRepo:       mockInterfaces.NewMockUserRepository(t),
			hasher:         mockInterfaces.NewMockHasher(t),
			expectOutput:   &OngUsecase{},
			expectedErrMsg: "",
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			usecase := NewOngUseCase(tcase.repo, tcase.userRepo, tcase.hasher)
			assert.IsTypef(t, tcase.expectOutput, usecase, "error: NewOngUsecase does not return expected type", nil)
		})
	}
}
