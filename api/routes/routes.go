package routes

import (
	"pet-dex-backend/v2/api/controllers"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	PetController          *controllers.PetController
	ExampleController      *petcontroller.ExampleController
	FindPetController      *petcontroller.FindPetController
	ListUserPetsController *petcontroller.ListUserPetsController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Route("/pets", func(r chi.Router) {
			r.Get("/", controllers.ExampleController.ExampleHandler)
			r.Get("/{id}", controllers.FindPetController.FindPet)
			r.Get("/my-pets/{userID}", controllers.ListUserPetsController.ListUserPets)
			r.Patch("/{petID}", controllers.PetController.Update)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {

		})
	})
}
