package main

import (
	"fmt"
	"net/http"
	"pet-dex-backend/v2/api/routes"
	"pet-dex-backend/v2/infra/config"
)

func main() {
	env, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}
	config.InitConfigs()

	router := routes.InitializeRouter()

	fmt.Printf("running on port %v", env.PORT)
	http.ListenAndServe(env.PORT, router)
}
