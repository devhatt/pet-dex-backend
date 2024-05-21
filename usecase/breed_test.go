package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"
)

type MockBreedRepository struct {
	mock.Mock
}

func (m *MockBreedRepository) FindByID(ID uniqueEntityId.ID) (*entity.Breed, error) {
	args := m.Called(ID)
	return args.Get(0).(*entity.Breed), args.Error(1)
}

func (m *MockBreedRepository) List() ([]*dto.BreedList, error) {
	args := m.Called()
	return args.Get(0).([]*dto.BreedList), args.Error(1)
}

func TestBreedFindByID(t *testing.T) {
	ID := uniqueEntityId.NewID()

	expectedBreed := &entity.Breed{ID: ID, Name: "Pastor Alemão", Specie: "Dog"}

	mockRepo := new(MockBreedRepository)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("FindByID", ID).Return(expectedBreed, nil)
	usecase := NewBreedUseCase(mockRepo)

	resultPet, err := usecase.FindByID(ID)

	assert.NoError(t, err)
	assert.NotNil(t, resultPet)
	assert.Equal(t, expectedBreed, resultPet)
}