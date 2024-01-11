package main

import (
	"fmt"
	"net/http"
	"pet-dex-backend/v2/api/routes"
	"pet-dex-backend/v2/infra/config"
)

func main() {
	config.InitConfigs()

	router := routes.InitializeRouter()

	fmt.Println("running on port 8000")
	http.ListenAndServe(":8000", router)
}
