package usecase

import (
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"

	"github.com/stretchr/testify/assert"

	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"
)

func TestBreedFindByID(t *testing.T) {
	ID := uniqueEntityId.NewID()

	expectedBreed := &entity.Breed{ID: ID, Name: "Pastor Alemão", Specie: "Dog"}

	mockRepo := mockInterfaces.NewMockBreedRepository(t)
	defer mockRepo.AssertExpectations(t)

	mockRepo.On("FindByID", ID).Return(expectedBreed, nil)
	usecase := NewBreedUseCase(mockRepo)

	resultPet, err := usecase.FindByID(ID)

	assert.NoError(t, err)
	assert.NotNil(t, resultPet)
	assert.Equal(t, expectedBreed, resultPet)
}
