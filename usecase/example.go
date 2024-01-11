package usecase

import (
	"pet-dex-backend/v2/interfaces"
)

type ExampleUsecase struct {
	petRepository interfaces.PetRepository
}

func NewExampleUseCase(p interfaces.PetRepository) *ExampleUsecase {
	return &ExampleUsecase{
		petRepository: p,
	}
}

func (auc *ExampleUsecase) Do() string {
	return "hello world"
}
