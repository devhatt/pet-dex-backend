package usecase

import (
	"errors"
	"pet-dex-backend/v2/entity/dto"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"
)

type MockPetRepository struct {
	mock.Mock
}

func (m *MockPetRepository) Save(entity.Pet) error {
	return nil
}

func (m *MockPetRepository) FindByID(ID uniqueEntityId.ID) (*entity.Pet, error) {
	args := m.Called(ID)
	return args.Get(0).(*entity.Pet), args.Error(1)
}

func (m *MockPetRepository) Update(petID string, userID string, petToUpdate *entity.Pet) error {
	args := m.Called(petID, userID, petToUpdate)
	return args.Error(0)
}

func (m *MockPetRepository) ListByUser(userID uniqueEntityId.ID) ([]*entity.Pet, error) {
	args := m.Called(userID)
	return args.Get(0).([]*entity.Pet), args.Error(1)
}

func TestUpdateUseCaseDo(t *testing.T) {
	id := "123"
	Data, _ := time.Parse(time.DateTime, "2023-09-20")
	Birthdate, _ := time.Parse(time.DateTime, "2023-09-20")
	userID := uniqueEntityId.NewID()
	petUpdateDto := dto.PetUpdateDto{Size: "small", AdoptionDate: Data, Birthdate: Birthdate, Weight: 4.53, WeightMeasure: "kg"}
	mockRepo := new(MockPetRepository)
	mockRepo.On("Update", id, userID.String(), entity.ToEntity(&petUpdateDto)).Return(nil)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Update(id, userID.String(), petUpdateDto)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUseCaseDoInvalidSize(t *testing.T) {
	id := "123"
	userID := uniqueEntityId.NewID()
	petUpdateDto := dto.PetUpdateDto{Size: "Invalid Size"}
	mockRepo := new(MockPetRepository)
	mockRepo.On("Update", id, userID.String(), entity.ToEntity(&petUpdateDto)).Return(nil)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Update(id, userID.String(), petUpdateDto)

	assert.EqualError(t, err, "the animal size is invalid")
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdateUseCaseDoRepositoryError(t *testing.T) {
	id := "123"
	userID := "321"
	petUpdateDto := dto.PetUpdateDto{Size: "small", Weight: 4.53, WeightMeasure: "kg"}
	repoError := errors.New("error updating pet")
	mockRepo := new(MockPetRepository)
	mockRepo.On("Update", id, userID, entity.ToEntity(&petUpdateDto)).Return(repoError)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Update(id, userID, petUpdateDto)

	assert.EqualError(t, err, "failed to update pet with ID 123: error updating pet")
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

func TestUpdateUseCaseDoVaccines(t *testing.T) {
	id := "123"
	userID := uniqueEntityId.NewID().String()
	vaccines := []dto.VaccinesDto{
		{Name: "Rabies", Date: time.Now(), DoctorCRM: "123456"},
		{Name: "Distemper", Date: time.Now(), DoctorCRM: "123456"},
	}
	petUpdateDto := dto.PetUpdateDto{Size: "medium", Vaccines: vaccines, Weight: 4.53, WeightMeasure: "kg"}
	mockRepo := new(MockPetRepository)
	mockRepo.On("Update", id, userID, entity.ToEntity(&petUpdateDto)).Return(nil)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Update(id, userID, petUpdateDto)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCaseDoVaccinesError(t *testing.T) {
	id := "123"
	userID := "321"
	vaccines := []dto.VaccinesDto{
		{Name: "Rabies", Date: time.Now(), DoctorCRM: "123456"},
		{Name: "Distemper", Date: time.Now(), DoctorCRM: "123456"},
	}
	petUpdateDto := dto.PetUpdateDto{Size: "small", Vaccines: vaccines, Weight: 4.53, WeightMeasure: "kg"}
	repoError := errors.New("error updating vaccines")
	mockRepo := new(MockPetRepository)
	mockRepo.On("Update", id, userID, entity.ToEntity(&petUpdateDto)).Return(repoError)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Update(id, userID, petUpdateDto)

	assert.EqualError(t, err, "failed to update pet with ID 123: error updating vaccines")
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCaseValidWeight(t *testing.T) {
	usecase := PetUseCase{}

	assert.True(t, usecase.isValidWeight(&entity.Pet{Weight: 1, WeightMeasure: "kg"}))
	assert.True(t, usecase.isValidWeight(&entity.Pet{Weight: 1, WeightMeasure: "lb"}))
	assert.False(t, usecase.isValidWeight(&entity.Pet{Weight: 0, WeightMeasure: "kg"}))
	assert.False(t, usecase.isValidWeight(&entity.Pet{Weight: 1, WeightMeasure: "invalid"}))
}

func TestListUserPets(t *testing.T) {
	userID := uniqueEntityId.NewID()

	var availableToAdoption = true
	expectedPets := []*entity.Pet{
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Rex", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Thor", AvailableToAdoption: &availableToAdoption},
	}

	mockRepo := new(MockPetRepository)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("ListByUser", userID).Return(expectedPets, nil)
	usecase := NewPetUseCase(mockRepo)

	pets, err := usecase.ListUserPets(userID)

	assert.NoError(t, err)
	assert.Len(t, pets, 2)
}

func TestListUserPetsNoPetsFound(t *testing.T) {
	userID := uniqueEntityId.NewID()

	mockRepo := new(MockPetRepository)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("ListByUser", userID).Return([]*entity.Pet{}, nil)
	usecase := NewPetUseCase(mockRepo)

	pets, err := usecase.ListUserPets(userID)

	assert.NoError(t, err)
	assert.Len(t, pets, 0)
}

func TestListUserPetsErrorOnRepo(t *testing.T) {
	userID := uniqueEntityId.NewID()

	mockRepo := new(MockPetRepository)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("ListByUser", userID).Return([]*entity.Pet{}, errors.New("this is a repository error"))
	usecase := NewPetUseCase(mockRepo)

	pets, err := usecase.ListUserPets(userID)

	assert.Error(t, err)
	assert.Nil(t, pets)
	assert.EqualError(t, err, "failed to retrieve all user pets: this is a repository error")
}

func TestFindByID(t *testing.T) {
	ID := uniqueEntityId.NewID()

	var availabelToAdoption = true
	expectedPet := &entity.Pet{ID: ID, UserID: uniqueEntityId.NewID(), Name: "Rex", AvailableToAdoption: &availabelToAdoption}

	mockRepo := new(MockPetRepository)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("FindByID", ID).Return(expectedPet, nil)
	usecase := NewPetUseCase(mockRepo)

	resultPet, err := usecase.FindByID(ID)

	assert.NoError(t, err)
	assert.NotNil(t, resultPet)
	assert.Equal(t, expectedPet, resultPet)
}

func TestFindByIDNilResult(t *testing.T) {
	petID := uniqueEntityId.NewID()
	var pet *entity.Pet

	mockRepo := new(MockPetRepository)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("FindByID", petID).Return(pet, errors.New("sql: no rows in result set"))
	usecase := NewPetUseCase(mockRepo)

	resultPet, err := usecase.FindByID(petID)

	assert.Error(t, err)
	assert.Nil(t, resultPet)
	assert.EqualError(t, err, "failed to retrieve pet: sql: no rows in result set")
}

func TestFindByIDErrorOnRepo(t *testing.T) {
	petID := uniqueEntityId.NewID()
	var pet *entity.Pet

	mockRepo := new(MockPetRepository)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("FindByID", petID).Return(pet, errors.New("this is a repository error"))
	usecase := NewPetUseCase(mockRepo)

	resultPet, err := usecase.FindByID(petID)

	assert.Error(t, err)
	assert.Nil(t, resultPet)
	assert.EqualError(t, err, "failed to retrieve pet: this is a repository error")
}
func TestPetUseCase_Save(t *testing.T) {
	birthdateString := "2016/10/21"
	adoptDateString := "2018/07/29"

	birthDate, _ := time.Parse(config.StandardDateLayout, birthdateString)
	adoptDate, _ := time.Parse(config.StandardDateLayout, adoptDateString)

	petToSave := dto.PetInsertDto{
		Name:         "Felpudo",
		UserID:       uniqueEntityId.NewID(),
		BreedID:      uniqueEntityId.NewID(),
		Weight:       4,
		Size:         "Médio",
		Birthdate:    &birthDate,
		AdoptionDate: &adoptDate,
	}

	mockRepo := new(MockPetRepository)
	mockRepo.On("Save", mock.AnythingOfType("entity.Pet")).Return(nil)

	usecase := NewPetUseCase(mockRepo)
	err := usecase.Save(petToSave)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestPetUseCase_SaveErrorOnRepo(t *testing.T) {
	birthdateString := "2016/10/21"
	adoptDateString := "2018/07/29"

	birthDate, _ := time.Parse(config.StandardDateLayout, birthdateString)
	adoptDate, _ := time.Parse(config.StandardDateLayout, adoptDateString)

	petToSave := dto.PetInsertDto{
		Name:         "",
		UserID:       uniqueEntityId.NewID(),
		BreedID:      uniqueEntityId.NewID(),
		Weight:       4,
		Size:         "Médio",
		Birthdate:    &birthDate,
		AdoptionDate: &adoptDate,
	}

	repoError := errors.New("error saving pet")
	mockRepo := new(MockPetRepository)
	mockRepo.On("Save", mock.AnythingOfType("entity.Pet")).Return(repoError)
	usecase := NewPetUseCase(mockRepo)

	err := usecase.Save(petToSave)

	assert.EqualError(t, err, "failed to save pet: error saving pet")
	mockRepo.AssertExpectations(t)
}
