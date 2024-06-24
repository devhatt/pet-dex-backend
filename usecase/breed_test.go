package usecase

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBreedUseCase(t *testing.T) {
	tcases := map[string]struct {
		repo         interfaces.BreedRepository
		expectOutput *BreedUseCase
	}{
		"success": {
			repo:         mockInterfaces.NewMockBreedRepository(t),
			expectOutput: &BreedUseCase{},
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			usecase := NewBreedUseCase(tcase.repo)

			assert.IsTypef(t, tcase.expectOutput, usecase, "error: New Hasher not returns a *Hasher{} struct", nil)
		})
	}
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
