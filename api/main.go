package main

import (
	"fmt"
	"net/http"
	"os"
	"pet-dex-backend/v2/api/routes"
	"pet-dex-backend/v2/infra/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config.InitConfigs()
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))
	if os.Getenv("ENVIROMENT") != "DEVELOPMENT" {
		router.Use(middleware.Logger)
	}

	routes.InitRoutes(router)

	fmt.Println("running on port 8000")
	http.ListenAndServe(":8000", router)
}
