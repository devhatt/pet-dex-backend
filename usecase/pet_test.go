package usecase

import (
	"errors"
	"time"

	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"

	"pet-dex-backend/v2/entity"
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUseCaseDo(t *testing.T) {
	id := "123"
	Data, _ := time.Parse(time.DateTime, "2023-09-20")
	Birthdate, _ := time.Parse(time.DateTime, "2023-09-20")
	userId := uniqueEntityId.NewID()
	petUpdateDto := dto.PetUpdateDto{Size: "small", AdoptionDate: Data, Birthdate: Birthdate, Weight: 4.53, WeightMeasure: "kg"}

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        dto.PetUpdateDto
		petId        string
		userId       string
		expectOutput error
	}{
		"success": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        petUpdateDto,
			petId:        id,
			userId:       userId.String(),
			expectOutput: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Update", tcase.petId, tcase.userId, entity.PetToEntity((&tcase.input))).Return(tcase.expectOutput)

			usecase := NewPetUseCase(tcase.repo)
			err := usecase.Update(tcase.petId, tcase.userId, tcase.input)

			assert.NoError(t, err)
			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestUseCaseDoInvalidSize(t *testing.T) {
	id := "123"
	Data, _ := time.Parse(time.DateTime, "2023-09-20")
	Birthdate, _ := time.Parse(time.DateTime, "2023-09-20")
	userId := uniqueEntityId.NewID()
	petUpdateDto := dto.PetUpdateDto{Size: "Invalid Size", AdoptionDate: Data, Birthdate: Birthdate, Weight: 4.53, WeightMeasure: "kg"}

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        dto.PetUpdateDto
		petId        string
		userId       string
		expectOutput error
	}{
		"error": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        petUpdateDto,
			petId:        id,
			userId:       userId.String(),
			expectOutput: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {

			usecase := NewPetUseCase(tcase.repo)
			err := usecase.Update(tcase.petId, tcase.userId, tcase.input)

			assert.EqualError(t, err, "the animal size is invalid")
			tcase.repo.AssertNotCalled(t, "Update")
		})
	}
}

func TestUpdateUseCaseDoRepositoryError(t *testing.T) {
	id := "123"
	Data, _ := time.Parse(time.DateTime, "2023-09-20")
	Birthdate, _ := time.Parse(time.DateTime, "2023-09-20")
	userId := uniqueEntityId.NewID()
	petUpdateDto := dto.PetUpdateDto{Size: "small", AdoptionDate: Data, Birthdate: Birthdate, Weight: 4.53, WeightMeasure: "kg"}
	repoError := errors.New("error updating pet")

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        dto.PetUpdateDto
		petId        string
		userId       string
		expectOutput error
	}{
		"error": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        petUpdateDto,
			petId:        id,
			userId:       userId.String(),
			expectOutput: repoError,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Update", tcase.petId, tcase.userId, entity.PetToEntity((&tcase.input))).Return(tcase.expectOutput)

			usecase := NewPetUseCase(tcase.repo)
			err := usecase.Update(tcase.petId, tcase.userId, tcase.input)

			assert.EqualError(t, err, "failed to update pet with ID 123: error updating pet")
		})
	}
}

func TestUpdateUseCaseisValidSize(t *testing.T) {
	usecase := NewPetUseCase(nil)

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        *entity.Pet
		expectOutput bool
	}{
		"small": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Size: "small"},
			expectOutput: true,
		},
		"medium": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Size: "medium"},
			expectOutput: true,
		},
		"large": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Size: "large"},
			expectOutput: true,
		},
		"giant": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Size: "giant"},
			expectOutput: true,
		},
		"Invalid Size": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Size: "Invalid Size"},
			expectOutput: false,
		},
		"empty": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Size: ""},
			expectOutput: false,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			value := usecase.isValidPetSize(tcase.input)
			assert.Equal(t, tcase.expectOutput, value)
		})
	}
}

func TestUpdateUseCaseDoVaccines(t *testing.T) {
	id := "123"
	userId := uniqueEntityId.NewID()
	vaccines := []dto.VaccinesDto{
		{Name: "Rabies", Date: time.Now(), DoctorCRM: "123456"},
		{Name: "Distemper", Date: time.Now(), DoctorCRM: "123456"},
	}
	petUpdateDto := dto.PetUpdateDto{Size: "medium", Vaccines: vaccines, Weight: 4.53, WeightMeasure: "kg"}

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        dto.PetUpdateDto
		petId        string
		userId       string
		expectOutput error
	}{
		"success": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        petUpdateDto,
			petId:        id,
			userId:       userId.String(),
			expectOutput: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Update", tcase.petId, tcase.userId, entity.PetToEntity((&tcase.input))).Return(tcase.expectOutput)

			usecase := NewPetUseCase(tcase.repo)
			err := usecase.Update(tcase.petId, tcase.userId, tcase.input)

			assert.NoError(t, err)
			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
		})
	}
}

func TestUpdateUseCaseDoVaccinesError(t *testing.T) {
	id := "123"
	userId := uniqueEntityId.NewID()
	vaccines := []dto.VaccinesDto{
		{Name: "Rabies", Date: time.Now(), DoctorCRM: "123456"},
		{Name: "Distemper", Date: time.Now(), DoctorCRM: "123456"},
	}
	petUpdateDto := dto.PetUpdateDto{Size: "medium", Vaccines: vaccines, Weight: 4.53, WeightMeasure: "kg"}
	repoError := errors.New("error updating vaccines")

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        dto.PetUpdateDto
		petId        string
		userId       string
		expectOutput error
	}{
		"success": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        petUpdateDto,
			petId:        id,
			userId:       userId.String(),
			expectOutput: repoError,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Update", tcase.petId, tcase.userId, entity.PetToEntity((&tcase.input))).Return(tcase.expectOutput)

			usecase := NewPetUseCase(tcase.repo)
			err := usecase.Update(tcase.petId, tcase.userId, tcase.input)

			assert.EqualError(t, err, "failed to update pet with ID 123: error updating vaccines")
			tcase.repo.AssertExpectations(t)
		})
	}
}

func TestUpdateUseCaseValidWeight(t *testing.T) {
	usecase := NewPetUseCase(nil)

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        *entity.Pet
		expectOutput bool
	}{
		"success kg": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Weight: 1, WeightMeasure: "kg"},
			expectOutput: true,
		},
		"success lb": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Weight: 1, WeightMeasure: "lb"},
			expectOutput: true,
		},
		"error kg": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Weight: 0, WeightMeasure: "kg"},
			expectOutput: false,
		},
		"error lb": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Weight: 0, WeightMeasure: "lb"},
			expectOutput: false,
		},
		"error invalid": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        &entity.Pet{Weight: 1, WeightMeasure: "invalid"},
			expectOutput: false,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			value := usecase.isValidWeight(tcase.input)
			assert.Equal(t, tcase.expectOutput, value)
		})
	}
}

func TestListUserPets(t *testing.T) {
	userId := uniqueEntityId.NewID()

	var availableToAdoption = true
	expectedPets := []*entity.Pet{
		{ID: uniqueEntityId.NewID(), UserID: userId, Name: "Rex", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userId, Name: "Thor", AvailableToAdoption: &availableToAdoption},
	}

	tcases := map[string]struct {
		repo          *mockInterfaces.MockPetRepository
		userId        uuid.UUID
		expectOutput  []*entity.Pet
		expectedError error
	}{
		"success": {
			repo:          mockInterfaces.NewMockPetRepository(t),
			userId:        userId,
			expectOutput:  expectedPets,
			expectedError: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListByUser", tcase.userId).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListUserPets(userId)

			assert.NoError(t, err)
			assert.Len(t, pets, 2)
			assert.Equal(t, tcase.expectOutput, pets, "expected error mismatch")
		})
	}
}

func TestListUserPetsNoPetsFound(t *testing.T) {
	userId := uniqueEntityId.NewID()

	tcases := map[string]struct {
		repo          *mockInterfaces.MockPetRepository
		petId         uuid.UUID
		userId        uuid.UUID
		expectOutput  []*entity.Pet
		expectedError error
	}{
		"success": {
			repo:          mockInterfaces.NewMockPetRepository(t),
			userId:        userId,
			expectOutput:  []*entity.Pet{},
			expectedError: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListByUser", tcase.userId).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListUserPets(tcase.userId)

			assert.NoError(t, err)
			assert.Len(t, pets, 0)
			assert.Equal(t, tcase.expectOutput, pets, "expected error mismatch")
		})
	}
}

func TestListUserPetsErrorOnRepo(t *testing.T) {
	userId := uniqueEntityId.NewID()

	tcases := map[string]struct {
		repo          *mockInterfaces.MockPetRepository
		petId         string
		userId        uuid.UUID
		expectOutput  []*entity.Pet
		expectedError error
	}{
		"success": {
			repo:          mockInterfaces.NewMockPetRepository(t),
			userId:        userId,
			expectOutput:  nil,
			expectedError: errors.New("this is a repository error"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListByUser", tcase.userId).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListUserPets(tcase.userId)

			assert.Error(t, err)
			assert.Nil(t, pets)
			assert.EqualError(t, err, "failed to retrieve all user pets: this is a repository error")
		})
	}
}

func TestFindByID(t *testing.T) {
	petId := uniqueEntityId.NewID()

	var availabelToAdoption = true
	expectedPet := &entity.Pet{ID: petId, UserID: uniqueEntityId.NewID(), Name: "Rex", AvailableToAdoption: &availabelToAdoption}

	tcases := map[string]struct {
		repo          *mockInterfaces.MockPetRepository
		petId         uuid.UUID
		expectOutput  *entity.Pet
		expectedError error
	}{
		"success": {
			repo:          mockInterfaces.NewMockPetRepository(t),
			petId:         petId,
			expectOutput:  expectedPet,
			expectedError: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("FindByID", tcase.petId).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pet, err := usecase.FindByID(tcase.petId)

			assert.NoError(t, err)
			assert.NotNil(t, pet)
			assert.Equal(t, tcase.expectOutput, pet, "expected error mismatch")
		})
	}
}

func TestFindByIDNilResult(t *testing.T) {
	petId := uniqueEntityId.NewID()

	expectedPet := &entity.Pet{ID: petId}

	tcases := map[string]struct {
		repo          *mockInterfaces.MockPetRepository
		petId         uuid.UUID
		expectOutput  *entity.Pet
		expectedError error
	}{
		"success": {
			repo:          mockInterfaces.NewMockPetRepository(t),
			petId:         petId,
			expectOutput:  expectedPet,
			expectedError: errors.New("sql: no rows in result set"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("FindByID", tcase.petId).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pet, err := usecase.FindByID(tcase.petId)

			assert.Error(t, err)
			assert.Nil(t, pet)
			assert.EqualError(t, err, "failed to retrieve pet: sql: no rows in result set")
		})
	}
}

func TestFindByIDErrorOnRepo(t *testing.T) {
	petId := uniqueEntityId.NewID()

	expectedPet := &entity.Pet{ID: petId}

	tcases := map[string]struct {
		repo          *mockInterfaces.MockPetRepository
		petId         uuid.UUID
		expectOutput  *entity.Pet
		expectedError error
	}{
		"success": {
			repo:          mockInterfaces.NewMockPetRepository(t),
			petId:         petId,
			expectOutput:  expectedPet,
			expectedError: errors.New("this is a repository error"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("FindByID", petId).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pet, err := usecase.FindByID(tcase.petId)

			assert.Error(t, err)
			assert.Nil(t, pet)
			assert.EqualError(t, err, "failed to retrieve pet: this is a repository error")
		})
	}
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

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        dto.PetInsertDto
		expectOutput error
	}{
		"success": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        petToSave,
			expectOutput: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Save", mock.AnythingOfType("*entity.Pet")).Return(tcase.expectOutput)

			usecase := NewPetUseCase(tcase.repo)
			err := usecase.Save(tcase.input)

			assert.NoError(t, err)

			tcase.repo.AssertExpectations(t)
		})
	}
}

func TestPetUseCase_SaveErrorOnRepo(t *testing.T) {
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

	tcases := map[string]struct {
		repo         *mockInterfaces.MockPetRepository
		input        dto.PetInsertDto
		expectOutput error
	}{
		"success": {
			repo:         mockInterfaces.NewMockPetRepository(t),
			input:        petToSave,
			expectOutput: errors.New("error saving pet"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("Save", mock.AnythingOfType("*entity.Pet")).Return(tcase.expectOutput)

			usecase := NewPetUseCase(tcase.repo)
			err := usecase.Save(tcase.input)

			assert.EqualError(t, err, "failed to save pet: error saving pet")
			tcase.repo.AssertExpectations(t)
		})
	}
}

func TestListPetsUnauthenticated(t *testing.T) {
	userID := uniqueEntityId.NewID()

	isUnauthorized := true
	var availableToAdoption = true
	allPets := []*entity.Pet{
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Rex", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Thor", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Pedi", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Bob", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Bidu", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Mafu", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Pilão", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Chocolate", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Docinho", AvailableToAdoption: &availableToAdoption},
	}

	expectedPets := allPets[:6]

	tcases := map[string]struct {
		repo           *mockInterfaces.MockPetRepository
		expectOutput   []*entity.Pet
		input          int
		isUnauthorized bool
		expectedError  error
	}{
		"success": {
			repo:           mockInterfaces.NewMockPetRepository(t),
			input:          1,
			isUnauthorized: isUnauthorized,
			expectOutput:   expectedPets,
			expectedError:  nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListAllByPage", tcase.input).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListPetsByPage(tcase.input, tcase.isUnauthorized)

			assert.NoError(t, err)
			assert.Len(t, pets, 6)
			assert.Equal(t, tcase.expectOutput, pets)
		})
	}
}

func TestListPetsUnauthenticatedNoPets(t *testing.T) {

	isUnauthorized := true
	expectedPets := []*entity.Pet{}

	tcases := map[string]struct {
		repo           *mockInterfaces.MockPetRepository
		expectOutput   []*entity.Pet
		input          int
		isUnauthorized bool
		expectedError  error
	}{
		"success": {
			repo:           mockInterfaces.NewMockPetRepository(t),
			input:          1,
			isUnauthorized: isUnauthorized,
			expectOutput:   expectedPets,
			expectedError:  nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListAllByPage", tcase.input).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListPetsByPage(tcase.input, tcase.isUnauthorized)

			assert.NoError(t, err)
			assert.Len(t, pets, 0)
			assert.Equal(t, tcase.expectOutput, pets)
		})
	}
}

func TestListPetsUnauthenticatedErrorOnRepo(t *testing.T) {

	isUnauthorized := true
	expectedPets := []*entity.Pet{}

	tcases := map[string]struct {
		repo           *mockInterfaces.MockPetRepository
		expectOutput   []*entity.Pet
		input          int
		isUnauthorized bool
		expectedError  error
	}{
		"success": {
			repo:           mockInterfaces.NewMockPetRepository(t),
			input:          1,
			isUnauthorized: isUnauthorized,
			expectOutput:   expectedPets,
			expectedError:  errors.New("repository error"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListAllByPage", tcase.input).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListPetsByPage(tcase.input, tcase.isUnauthorized)

			assert.Error(t, err)
			assert.Len(t, pets, 0)
			assert.EqualError(t, err, "failed to retrieve all user pets: repository error")
		})
	}
}

func TestListPetsAuthenticated(t *testing.T) {
	userID := uniqueEntityId.NewID()

	isUnauthorized := false
	var availableToAdoption = true
	allPets := []*entity.Pet{
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Rex", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Thor", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Pedi", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Bob", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Bidu", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Mafu", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Pilão", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Chocolate", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Docinho", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Amarelo", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Fofão", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Caramelo", AvailableToAdoption: &availableToAdoption},
		{ID: uniqueEntityId.NewID(), UserID: userID, Name: "Dogão", AvailableToAdoption: &availableToAdoption},
	}

	expectedPets := allPets[:12]

	tcases := map[string]struct {
		repo           *mockInterfaces.MockPetRepository
		expectOutput   []*entity.Pet
		input          int
		isUnauthorized bool
		expectedError  error
	}{
		"success": {
			repo:           mockInterfaces.NewMockPetRepository(t),
			input:          1,
			isUnauthorized: isUnauthorized,
			expectOutput:   expectedPets,
			expectedError:  nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListAllByPage", tcase.input).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListPetsByPage(tcase.input, tcase.isUnauthorized)

			assert.NoError(t, err)
			assert.Len(t, pets, 12)
			assert.Equal(t, tcase.expectOutput, pets)
		})
	}
}

func TestListPetsAuthenticatedNoPets(t *testing.T) {

	isUnauthorized := false

	expectedPets := []*entity.Pet{}

	tcases := map[string]struct {
		repo           *mockInterfaces.MockPetRepository
		expectOutput   []*entity.Pet
		input          int
		isUnauthorized bool
		expectedError  error
	}{
		"success": {
			repo:           mockInterfaces.NewMockPetRepository(t),
			input:          1,
			isUnauthorized: isUnauthorized,
			expectOutput:   expectedPets,
			expectedError:  nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListAllByPage", tcase.input).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListPetsByPage(tcase.input, tcase.isUnauthorized)

			assert.NoError(t, err)
			assert.Len(t, pets, 0)
			assert.Equal(t, tcase.expectOutput, pets)
		})
	}
}

func TestListPetsAuthenticatedErrorOnRepo(t *testing.T) {

	isUnauthorized := false

	expectedPets := []*entity.Pet{}

	tcases := map[string]struct {
		repo           *mockInterfaces.MockPetRepository
		expectOutput   []*entity.Pet
		input          int
		isUnauthorized bool
		expectedError  error
	}{
		"success": {
			repo:           mockInterfaces.NewMockPetRepository(t),
			input:          1,
			isUnauthorized: isUnauthorized,
			expectOutput:   expectedPets,
			expectedError:  errors.New("repository error"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("ListAllByPage", tcase.input).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewPetUseCase(tcase.repo)
			pets, err := usecase.ListPetsByPage(tcase.input, tcase.isUnauthorized)

			assert.Error(t, err)
			assert.Len(t, pets, 0)
			assert.EqualError(t, err, "failed to retrieve pets page: repository error")
		})
	}
}
