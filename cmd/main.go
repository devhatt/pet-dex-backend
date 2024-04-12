package main

import (
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase"

	"github.com/jmoiron/sqlx"
)

func main() {

	sqlxDb, err := sqlx.Connect("mysql", "dellis:@/shud")

	if err != nil {
		panic(err)
	}
	pr := db.NewPetRepository(sqlxDb)
	adoptUseCase := usecase.NewAdoptUseCase(pr)

	adoptUseCase.Do()
}
