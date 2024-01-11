package routes

import (
	"os"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRouter() *chi.Mux {
	database := config.GetDB()
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))
	if os.Getenv("ENVIROMENT") != "DEVELOPMENT" {
		router.Use(middleware.Logger)
	}

	dbPetRepo := db.NewPetRepository(database)

	exampleUseCase := usecase.NewExampleUseCase(dbPetRepo)
	findPetUseCase := usecase.FindPetUseCase(dbPetRepo)

	exampleController := petcontroller.NewExampleController(exampleUseCase)
	findPetController := petcontroller.NewFindPetController(findPetUseCase)

	contrllers := Controllers{
		FindPetController: findPetController,
		ExampleController: exampleController,
	}

	InitRoutes(contrllers, router)

	return router
}
