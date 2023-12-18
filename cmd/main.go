package main

import (
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase"
)

func main() {
	pr := db.NewPetRepository()
	adoptUseCase := usecase.NewAdoptUseCase(pr)

	adoptUseCase.Do()
}
