package usecase

import (
	"fmt"
	"pet-dex-backend/v2/interfaces"
)

type AdoptUseCase struct {
	petRepository interfaces.PetRepository
}

func NewAdoptUseCase(p interfaces.PetRepository) *AdoptUseCase {
	return &AdoptUseCase{
		petRepository: p,
	}
}

func (auc *AdoptUseCase) Do() {
	fmt.Println("Hello, World!")
}
