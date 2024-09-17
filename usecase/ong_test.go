package usecase

import (
	"pet-dex-backend/v2/entity/dto"
	mockInterfaces "pet-dex-backend/v2/mocks/pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOngDelete(t *testing.T) {
	tcases := map[string]struct {
		repo         *mockInterfaces.MockOngRepository
		inputID      uniqueEntityId.ID
		findByIDResp *dto.OngListMapper
		findByIDErr  error
		deleteResp   error
		expectOutput error
	}{
		"success": {
			repo:         mockInterfaces.NewMockOngRepository(t),
			inputID:      uniqueEntityId.NewID(),
			findByIDResp: &dto.OngListMapper{},
			findByIDErr:  nil,
			deleteResp:   nil,
			expectOutput: nil,
		},
	}

	for name, tcase := range tcases {
		t.Run(name, func(t *testing.T) {
			tcase.repo.On("FindByID", tcase.inputID).Return(nil, tcase.expectOutput)
			tcase.repo.On("Delete", tcase.inputID).Return(tcase.expectOutput)

			usecase := NewOngUseCase(tcase.repo, nil, nil)
			err := usecase.Delete(tcase.inputID)

			assert.Equal(t, tcase.expectOutput, err, "expected error mismatch")
			tcase.repo.AssertExpectations(t)
		})
	}
}
