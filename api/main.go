package main

import (
	"fmt"
	"log"
	"net/http"
	"pet-dex-backend/v2/api/controllers"
	"pet-dex-backend/v2/api/routes"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/pkg/hasher"
	"pet-dex-backend/v2/usecase"

	"github.com/jmoiron/sqlx"
)

func main() {
	env, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	config.InitConfigs()
	database := config.GetDB()
	config.RunUpMigrations(database)
	sqlxDb, err := sqlx.Connect("mysql", env.DBUrl)

	if err != nil {
		panic(err)
	}
	dbPetRepo := db.NewPetRepository(sqlxDb)
	dbUserRepo := db.NewUserRepository(sqlxDb)
	hash := hasher.NewHasher()

	petUsecase := usecase.NewPetUseCase(dbPetRepo)
	uusercase := usecase.NewUserUsecase(dbUserRepo, hash)
	petController := controllers.NewPetController(petUsecase)
	userController := controllers.NewUserController(uusercase)

	contrllers := routes.Controllers{
		PetController:  petController,
		UserController: userController,
	}
	router := routes.InitializeRouter(contrllers)

	fmt.Printf("running on port %v \n", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}
