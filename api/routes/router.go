package routes

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRouter(contrllers Controllers) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))
	if os.Getenv("ENVIROMENT") != "DEVELOPMENT" {
		router.Use(middleware.Logger)
	}

	InitRoutes(contrllers, router)
	return router
}
