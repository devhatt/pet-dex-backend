package main

import (
	"fmt"
	"log"
	"net/http"
	"pet-dex-backend/v2/api/controllers"
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
	config.RunMigrations(database)
	dbPetRepo := db.NewPetRepository(database)

	petUsecase := usecase.NewPetUseCase(dbPetRepo)

	petController := controllers.NewPetController(petUsecase)

	contrllers := routes.Controllers{
		PetController: petController,
	}
	router := routes.InitializeRouter(contrllers)

	fmt.Printf("running on port %v \n", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}
