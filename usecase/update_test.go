package usecase

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

func (m MockPetRepository) FindById(petID string) (entity.Pet, error) {
	args := m.Called(petID)
	return args.Get(0).(entity.Pet), args.Error(1)
}

func (MockPetRepository) Save(entity.Pet) error {
	return nil
}

func (m MockPetRepository) Update(petID string, userID string, updateValues map[string]interface{}) error {
	args := m.Called(petID, userID, updateValues)
	return args.Error(0)
}

//func TestUpdateUseCaseDoPetNotFound(t *testing.T) {
//	id := "123"
//	userID := "321"
//	petToUpdate := &entity.Pet{Size: "medium"}
//	mockRepo := new(MockPetRepository)
//	mockRepo.On("FindById", id).Return(entity.Pet{}, errors.New("pet with ID 123 not found"))
//	usecase := NewUpdateUseCase(mockRepo)
//
//	err := usecase.Do(id, userID, petToUpdate)
//
//	assert.EqualError(t, err, "pet with ID 123 not found")
//	mockRepo.AssertExpectations(t)
//	mockRepo.AssertNotCalled(t, "Update")
//}

func TestUpdateUseCaseDo(t *testing.T) {
	id := "123"
	userID := "321"
	petToUpdate := &entity.Pet{Size: "medium", UserID: "321"}
	mockRepo := new(MockPetRepository)
	//mockRepo.On("FindById", id).Return(&entity.Pet{ID: "123", UserID: "321"}, nil)
	mockRepo.On("Update", id, userID, map[string]interface{}{"size": &petToUpdate.Size}).Return(nil)
	usecase := NewUpdateUseCase(mockRepo)

	err := usecase.Do(id, userID, petToUpdate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

//	func TestUseCaseDoInvalidSize(t *testing.T) {
//		id := "123"
//		petToUpdate := &entity.Pet{Size: "Invalid Size"}
//		mockRepo := new(MockPetRepository)
//		usecase := usecase.NewUpdateUseCase(mockRepo)
//
//		err := usecase(id, petToUpdate)
//
//		assert.EqualError(t,err, "failed to update size for pet with ID 123: Size is invalid")
//		mockRepo.AssertNotCalled(t, "Update")
//	}

func TestUpdateUseCaseDoRepositoryError(t *testing.T) {
	id := "123"
	userID := "321"
	petToUpdate := &entity.Pet{Size: "small"}
	repoError := errors.New("error updating pet")
	mockRepo := new(MockPetRepository)
	mockRepo.On("Update", id, userID, mock.Anything).Return(repoError)
	usecase := NewUpdateUseCase(mockRepo)

	err := usecase.Do(id, userID, petToUpdate)

	assert.EqualError(t, err, "failed to update size for pet with ID 123: error updating pet")
	mockRepo.AssertExpectations(t)
}
func TestUpdateUseCaseisValidSize(t *testing.T) {
	usecase := UpdateUseCase{}

	assert.True(t, usecase.IsValidSize("small"))
	assert.True(t, usecase.IsValidSize("medium"))
	assert.True(t, usecase.IsValidSize("large"))
	assert.True(t, usecase.IsValidSize("giant"))
	assert.False(t, usecase.IsValidSize("Invalid Size"))
}
