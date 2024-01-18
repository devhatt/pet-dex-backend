package main

import (
	"database/sql"
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase/pet"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbc, err := sql.Open("mysql", "dellis:@/shud")
	if err != nil {
		panic(err)
	}
	defer dbc.Close()
	pr := db.NewPetRepository(dbc)
	adoptUseCase := pet.NewAdoptUseCase(pr)

	adoptUseCase.Do()
}
