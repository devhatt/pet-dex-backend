package main

import (
	"fmt"
	"log"
	"net/http"
	ongcontroller "pet-dex-backend/v2/api/controllers"
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
	dbOngRepo := db.NewOngRepository(database)

	exampleUseCase := usecase.NewExampleUseCase(dbPetRepo)
	findPetUseCase := usecase.NewPetUseCase(dbPetRepo)
	createOngUseCase := usecase.NewOngUseCase(dbOngRepo)

	exampleController := petcontroller.NewExampleController(exampleUseCase)
	findPetController := petcontroller.NewFindPetController(findPetUseCase)
	createOngController := ongcontroller.NewOngcontroller(createOngUseCase)

	contrllers := routes.Controllers{
		FindPetController:   findPetController,
		ExampleController:   exampleController,
		CreateOngController: createOngController,
	}

	router := routes.InitializeRouter(contrllers)

	fmt.Printf("running on port %v", env.PORT)
	log.Fatal(http.ListenAndServe(env.PORT, router))
}
