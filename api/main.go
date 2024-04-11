package main

import (
	"fmt"
	"log"
	"net/http"
	"pet-dex-backend/v2/api/controllers"
	"pet-dex-backend/v2/api/routes"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/pkg/encoder"
	"pet-dex-backend/v2/pkg/hasher"
	"pet-dex-backend/v2/usecase"

	"github.com/jmoiron/sqlx"
)

func main() {
	env, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	config.InitConfigs()
	sqlxDb, err := sqlx.Connect("mysql", env.DBUrl)

	if err != nil {
		panic(err)
	}

	dbPetRepo := db.NewPetRepository(sqlxDb)
	dbUserRepo := db.NewUserRepository(sqlxDb)
	bdBreedRepo := db.NewBreedRepository(sqlxDb)

	hash := hasher.NewHasher()
	encoder := encoder.NewEncoderAdapter(config.GetEnvConfig().JWT_SECRET)

	breedUsecase := usecase.NewBreedUseCase(bdBreedRepo)
	uusercase := usecase.NewUserUsecase(dbUserRepo, hash, encoder)
	petUsecase := usecase.NewPetUseCase(dbPetRepo)

	breedController := controllers.NewBreedController(breedUsecase)
	petController := controllers.NewPetController(petUsecase)
	userController := controllers.NewUserController(uusercase)

	controllers := routes.Controllers{
		PetController:   petController,
		BreedController: breedController,
		UserController:  userController,
	}
	router := routes.InitializeRouter(controllers)

	fmt.Printf("running on port %v \n", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}