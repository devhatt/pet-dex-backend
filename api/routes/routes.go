package routes

import (
	petcontroller "pet-dex-backend/v2/api/controllers/pet"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	UpdatePetController *petcontroller.UpdatePetController
	ExampleController   *petcontroller.ExampleController
	FindPetController *petcontroller.FindPetController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Route("/pets", func(r chi.Router) {
			r.Get("/", controllers.ExampleController.ExampleHandler)
			r.Patch("/{id}", controllers.UpdatePetController.Update)
			r.Get("/{id}", controllers.FindPetController.FindPet)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {

		})
	})
}
