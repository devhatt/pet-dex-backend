package main

import (
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	dbc, err := sqlx.Connect("mysql", "dellis:@/shud")
	if err != nil {
		panic(err)
	}
	defer dbc.Close()
	pr := db.NewPetRepository(dbc)
	adoptUseCase := usecase.NewAdoptUseCase(pr)

	adoptUseCase.Do()
}
