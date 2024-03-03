package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pet-dex-backend/v2/entity"
	pkgEntity "pet-dex-backend/v2/pkg/entity"

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
	userID := pkgEntity.NewID()
	petToUpdate := &entity.Pet{Size: "medium", UserID: userID}
	mockRepo := new(MockPetRepository)
	//mockRepo.On("FindById", id).Return(&entity.Pet{ID: "123", UserID: "321"}, nil)
	mockRepo.On("Update", id, userID.String(), map[string]interface{}{"size": &petToUpdate.Size}).Return(nil)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Update(id, userID.String(), petToUpdate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUseCaseDoInvalidSize(t *testing.T) {
	id := "123"
	userID := pkgEntity.NewID()
	petToUpdate := &entity.Pet{Size: "Invalid Size"}
	mockRepo := new(MockPetRepository)
	//mockRepo.On("FindById", id).Return(&entity.Pet{ID: "123", UserID: "321"}, nil)
	mockRepo.On("Update", id, userID.String(), map[string]interface{}{"size": &petToUpdate.Size}).Return(nil)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Update(id, userID.String(), petToUpdate)

	assert.EqualError(t, err, "The animal size is invalid")
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdateUseCaseDoRepositoryError(t *testing.T) {
	id := "123"
	userID := "321"
	petToUpdate := &entity.Pet{Size: "small"}
	repoError := errors.New("error updating pet")
	mockRepo := new(MockPetRepository)
	mockRepo.On("Update", id, userID, mock.Anything).Return(repoError)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Update(id, userID, petToUpdate)

	assert.EqualError(t, err, "failed to update size for pet with ID 123: error updating pet")
	mockRepo.AssertExpectations(t)
}
func TestUpdateUseCaseisValidSize(t *testing.T) {
	usecase := PetUseCase{}

	assert.True(t, usecase.isValidPetSize(&entity.Pet{Size: "small"}))
	assert.True(t, usecase.isValidPetSize(&entity.Pet{Size: "medium"}))
	assert.True(t, usecase.isValidPetSize(&entity.Pet{Size: "large"}))
	assert.True(t, usecase.isValidPetSize(&entity.Pet{Size: "giant"}))
	assert.False(t, usecase.isValidPetSize(&entity.Pet{Size: "Invalid Size"}))
	assert.False(t, usecase.isValidPetSize(&entity.Pet{Size: ""}))
}
