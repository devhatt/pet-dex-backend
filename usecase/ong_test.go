package usecase

import (
	"pet-dex-backend/v2/entity"
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
	ID := uniqueEntityId.NewID()
	expectedOng := &entity.Ong{
		ID:   ID,
		Name: "Peludinhos",
	}

	mockRepo := new(MockOngRepository)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("FindByID", ID).Return(expectedOng, nil)
	usecase := NewOngUseCase(mockRepo)

	resultOng, err := usecase.FindByID(ID)

	assert.NoError(t, err)
	assert.NotNil(t, resultOng)
	assert.Equal(t, expectedOng, resultOng)

}
