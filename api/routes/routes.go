package routes

import (
	ongcontroller "pet-dex-backend/v2/api/controllers/ong"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	FindPetController   *petcontroller.FindPetController
	ExampleController   *petcontroller.ExampleController
	CreateOngController *ongcontroller.CreateOngController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Route("/pets", func(r chi.Router) {
			r.Get("/", controllers.ExampleController.ExampleHandler)
			r.Get("/{id}", controllers.FindPetController.FindPet)
		})

		r.Route("/ong", func(r chi.Router) {
			r.Post("/", controllers.CreateOngController.CreateOng)
		})

		r.Route("/user", func(r chi.Router) {

		})
	})
}
