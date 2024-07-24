package usecase

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	uuidList := []uniqueEntityId.ID{
		uuid.MustParse("f3768895-d8cc-40d7-b8ae-8b7eb0eac26c"),
		uuid.MustParse("db6ba220-19dc-4f6c-845f-0fbf84c275b9"),
		uuid.MustParse("eb90009f-dfcc-4568-95b9-3f393ef9a9c2"),
		uuid.MustParse("43c7b32a-3e31-4894-8c93-0e8b29415caa"),
	}

	tcases := map[string]struct {
		repo          *mockInterfaces.MockBreedRepository
		expectOutput  []*dto.BreedList
		expectedError error
	}{
		"success": {
			repo: mockInterfaces.NewMockBreedRepository(t),
			expectOutput: []*dto.BreedList{
				{ID: uuidList[0], Name: "Amarelo", ImgUrl: "image url 1"},
				{ID: uuidList[1], Name: "Caramela", ImgUrl: "image url 2"},
				{ID: uuidList[2], Name: "Nuvem", ImgUrl: "image url 3"},
				{ID: uuidList[3], Name: "Thor", ImgUrl: "image url 4"},
			},
			expectedError: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("List").Return(tcase.expectOutput, nil)

			usecase := NewBreedUseCase(tcase.repo)
			list, err := usecase.List()

			assert.Equal(t, tcase.expectOutput, list, "expected output mismatch")
			assert.Equal(t, tcase.expectedError, err, "expected error mismatch")
		})
	}
}

func TestListErrorOnRepo(t *testing.T) {

	tcases := map[string]struct {
		repo          *mockInterfaces.MockBreedRepository
		expectOutput  []*dto.BreedList
		mockError     error
		expectedError error
	}{
		"errorList": {
			repo:          mockInterfaces.NewMockBreedRepository(t),
			expectOutput:  nil,
			mockError:     fmt.Errorf("error listing breeds"),
			expectedError: fmt.Errorf("error listing breeds: error listing breeds"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("List").Return(tcase.expectOutput, tcase.mockError)

			usecase := NewBreedUseCase(tcase.repo)
			_, err := usecase.List()

			assert.Equal(t, tcase.expectedError, err, "expected error mismatch")
		})
	}
}

func TestBreedFindByID(t *testing.T) {
	ID := uniqueEntityId.NewID()

	expectedBreed := &entity.Breed{ID: ID, Name: "Pastor Alemão", Specie: "Dog"}

	tcases := map[string]struct {
		repo          *mockInterfaces.MockBreedRepository
		expectOutput  *entity.Breed
		expectedError error
	}{
		"success": {
			repo:          mockInterfaces.NewMockBreedRepository(t),
			expectOutput:  expectedBreed,
			expectedError: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("FindByID", expectedBreed.ID).Return(tcase.expectOutput, tcase.expectedError)

			usecase := NewBreedUseCase(tcase.repo)
			list, err := usecase.FindByID(expectedBreed.ID)

			assert.Equal(t, tcase.expectOutput, list, "expected output mismatch")
			assert.Equal(t, tcase.expectedError, err, "expected error mismatch")
		})
	}
}

func TestBreedFindByIDErrorOnRepo(t *testing.T) {
	ID := uniqueEntityId.NewID()

	tcases := map[string]struct {
		repo          *mockInterfaces.MockBreedRepository
		expectOutput  *entity.Breed
		mockInput     error
		expectedError error
	}{
		"error": {
			repo:          mockInterfaces.NewMockBreedRepository(t),
			expectOutput:  nil,
			mockInput:     fmt.Errorf("error retrieving breed"),
			expectedError: fmt.Errorf("failed to retrieve breed:"),
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("FindByID", ID).Return(tcase.expectOutput, tcase.mockInput)

			usecase := NewBreedUseCase(tcase.repo)
			list, err := usecase.FindByID(ID)

			assert.EqualError(t, err, "failed to retrieve breed: error retrieving breed")
			assert.Equal(t, list, tcase.expectOutput)
			assert.Error(t, err, tcase.expectedError)
		})
	}
}
