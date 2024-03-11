package routes

import (
	"pet-dex-backend/v2/api/controllers"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	FindPetController *petcontroller.FindPetController
	ExampleController *petcontroller.ExampleController
	UserController    *controllers.UserController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Route("/pets", func(r chi.Router) {
			r.Get("/", controllers.ExampleController.ExampleHandler)
			r.Get("/{id}", controllers.FindPetController.FindPet)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/", controllers.UserController.Insert)
		})
	})
}
