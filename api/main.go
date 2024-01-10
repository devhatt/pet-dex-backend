package main

import (
	"fmt"
	"net/http"
	"os"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
	"pet-dex-backend/v2/api/routes"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database := config.InitConfigs()
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))
	if os.Getenv("ENVIROMENT") != "DEVELOPMENT" {
		router.Use(middleware.Logger)
	}

	dbPetRepo := db.NewPetRepository(database)

	findPetUseCase := usecase.FindPetUseCase(dbPetRepo)

	findPetController := petcontroller.NewFindPetController(findPetUseCase)

	contrllers := routes.Controllers{
		FindPetController: findPetController,
	}
	routes.InitRoutes(contrllers, router)

	fmt.Println("running on port 8000")
	http.ListenAndServe(":8000", router)
}
