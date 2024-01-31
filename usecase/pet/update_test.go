package pet

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pet-dex-backend/v2/entity"
	"testing"
)

type MockPetRepository struct {
	mock.Mock
}

func (m *MockPetRepository) FindById(id string) (*entity.Pet, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Pet), args.Error(1)
}

func (m *MockPetRepository) Save(entity.Pet) error {
	return nil
}

func (m *MockPetRepository) Update(id string, updateValues map[string]interface{}) error {
	args := m.Called(id, updateValues)
	return args.Error(0)
}

func TestUpdateUseCaseDoPetNotFound(t *testing.T) {
	id := "123"
	petToUpdate := &entity.Pet{Size: "medium"}
	mockRepo := new(MockPetRepository)
	mockRepo.On("FindById", id).Return(nil, errors.New("Not founded"))
	usecase := NewUpdateUseCase(mockRepo)

	err := usecase.Do(id, petToUpdate)

	assert.EqualError(t, err, "pet with ID 123 not found")
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdateUseCaseDo(t *testing.T) {
	id := "123"
	petToUpdate := &entity.Pet{Size: "medium", UserID: "321"}
	mockRepo := new(MockPetRepository)
	mockRepo.On("FindById", id).Return(&entity.Pet{ID: "123", UserID: "321"}, nil)
	mockRepo.On("Update", id, map[string]interface{}{"size": &petToUpdate.Size}).Return(nil)
	usecase := NewUpdateUseCase(mockRepo)

	err := usecase.Do(id, petToUpdate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUseCaseDoInvalidSize(t *testing.T) {
	id := "123"
	petToUpdate := &entity.Pet{Size: "Invalid Size", UserID: "321"}
	mockRepo := new(MockPetRepository)
	usecase := NewUpdateUseCase(mockRepo)
	mockRepo.On("FindById", id).Return(&entity.Pet{ID: "123", UserID: "321"}, nil)
	err := usecase.Do(id, petToUpdate)

	assert.EqualError(t, err, "Pet size is invalid")
}
func TestUpdateUseCaseDoRepositoryError(t *testing.T) {
	id := "123"
	petToUpdate := &entity.Pet{Size: "small"}
	repoError := errors.New("error updating pet")
	mockRepo := new(MockPetRepository)
	mockRepo.On("Update", id, mock.Anything).Return(repoError)
	usecase := NewUpdateUseCase(mockRepo)

	err := usecase.Do(id, petToUpdate)

	assert.EqualError(t, err, "failed to update size for pet with ID 123: error updating pet")
	mockRepo.AssertExpectations(t)
}
func TestUpdateUseCaseisValidSize(t *testing.T) {
	usecase := UpdateUseCase{}

	assert.True(t, usecase.isValidSize("small"))
	assert.True(t, usecase.isValidSize("medium"))
	assert.True(t, usecase.isValidSize("large"))
	assert.True(t, usecase.isValidSize("giant"))
	assert.False(t, usecase.isValidSize("Invalid Size"))
}
