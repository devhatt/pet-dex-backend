package routes

import (
	petcontroller "pet-dex-backend/v2/api/controllers/pet"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	FindPetController *petcontroller.FindPetController
	ExampleController *petcontroller.ExampleController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Route("/pets", func(r chi.Router) {
			r.Get("/", petcontroller.ExampleController)
			r.Patch("/{id}", controllers.FindPetController.FindPet)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {

		})
	})
}
