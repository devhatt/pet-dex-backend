package routes

import (
	"pet-dex-backend/v2/api/controllers"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	FindPetController *petcontroller.FindPetController
	ExampleController *petcontroller.ExampleController
	PetController     *controllers.PetController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", controllers.ExampleController.ExampleHandler)
			//r.Get("/{userID}", controllers.FindPetController.FindPet)
		})

		r.Route("/users/{userID}/pets", func(r chi.Router) {
			r.Get("/", controllers.ExampleController.ExampleHandler)
			r.Patch("/{petID}", controllers.PetController.Update)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {

		})
	})
}
