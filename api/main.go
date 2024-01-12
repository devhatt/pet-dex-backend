package main

import (
	"fmt"
	"log"
	"net/http"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
	"pet-dex-backend/v2/api/routes"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase"
)

func main() {
	env, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	database := config.InitConfigs()
	dbPetRepo := db.NewPetRepository(database)

	exampleUseCase := usecase.NewExampleUseCase(dbPetRepo)
	findPetUseCase := usecase.NewPetUseCase(dbPetRepo)
	updateUseCase := usecase.NewUpdateUseCase(dbPetRepo)

	exampleController := petcontroller.NewExampleController(exampleUseCase)
	findPetController := petcontroller.NewFindPetController(findPetUseCase)
	updateController := petcontroller.NewUpdatePetController(updateUseCase)

	contrllers := routes.Controllers{
		FindPetController: findPetController,
		ExampleController:    exampleController,
		UpdateSizeController: updateController,
	}

	router := routes.InitializeRouter(contrllers)

	fmt.Printf("running on port %v", env.PORT)

	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}
