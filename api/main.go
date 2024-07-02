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
	"pet-dex-backend/v2/pkg/sso"
	"pet-dex-backend/v2/usecase"

	"github.com/jmoiron/sqlx"
)

func main() {
	env, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	// config.InitConfigs()
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_DATABASE)
	sqlxDb := sqlx.MustConnect("mysql", databaseUrl)

	dbPetRepo := db.NewPetRepository(sqlxDb)
	dbUserRepo := db.NewUserRepository(sqlxDb)
	dbOngRepo := db.NewOngRepository(sqlxDb)
	hash := hasher.NewHasher()
	bdBreedRepo := db.NewBreedRepository(sqlxDb)

	encoder := encoder.NewEncoderAdapter(config.GetEnvConfig().JWT_SECRET)

	googleSsoGt := sso.NewGoogleGateway()
	facebookSsoGt := sso.NewFacebookGateway()

	ssoProvider := sso.NewProvider(googleSsoGt, facebookSsoGt)

	breedUsecase := usecase.NewBreedUseCase(bdBreedRepo)
	uusercase := usecase.NewUserUsecase(dbUserRepo, hash, encoder, ssoProvider)
	petUsecase := usecase.NewPetUseCase(dbPetRepo)
	ongUsecase := usecase.NewOngUseCase(dbOngRepo, dbUserRepo, hash)
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

	fmt.Printf("running on port %v \n", env.API_PORT)
	log.Fatal(http.ListenAndServe(":"+env.API_PORT, router))
}
