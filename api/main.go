package main

import (
	"fmt"
	"log"
	"net/http"
	"pet-dex-backend/v2/api/controllers"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
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

	database := config.InitConfigs()
	config.RunMigrations(database)
	sqlxDb, err := sqlx.Connect("mysql", env.DBUrl)

	if err != nil {
		panic(err)
	}
	dbPetRepo := db.NewPetRepository(database)
	dbUserRepo := db.NewUserRepository(sqlxDb)
	hash := hasher.NewHasher()

	exampleUseCase := usecase.NewExampleUseCase(dbPetRepo)
	findPetUseCase := usecase.NewPetUseCase(dbPetRepo)
	uusercase := usecase.NewUserUsecase(dbUserRepo, hash)

	exampleController := petcontroller.NewExampleController(exampleUseCase)
	findPetController := petcontroller.NewFindPetController(findPetUseCase)
	userController := controllers.NewUserController(uusercase)

	contrllers := routes.Controllers{
		FindPetController: findPetController,
		ExampleController: exampleController,
		UserController: userController,
	}
	router := routes.InitializeRouter(contrllers)

	fmt.Printf("running on port %v \n", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}
