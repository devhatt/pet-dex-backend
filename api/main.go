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
	dbOnRepo := db.NewOngRepository(sqlxDb)
	hash := hasher.NewHasher()
	bdBreedRepo := db.NewBreedRepository(sqlxDb)

	encoder := encoder.NewEncoderAdapter(config.GetEnvConfig().JWT_SECRET)

	breedUsecase := usecase.NewBreedUseCase(bdBreedRepo)
	uusercase := usecase.NewUserUsecase(dbUserRepo, hash, encoder)
	petUsecase := usecase.NewPetUseCase(dbPetRepo)
	ongUsecase := usecase.NewOngUseCase(dbOnRepo)
	breedController := controllers.NewBreedController(breedUsecase)
	petController := controllers.NewPetController(petUsecase)
	userController := controllers.NewUserController(uusercase)
	ongController := controllers.NewOngcontroller(ongUsecase)
	controllers := routes.Controllers{
		PetController:   petController,
		UserController:  userController,
		BreedController: breedController,
		OngController:   ongController,
	}
	router := routes.InitializeRouter(controllers)

	fmt.Printf("running on port %v \n", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}
