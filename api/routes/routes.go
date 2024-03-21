package routes

import (
	"pet-dex-backend/v2/api/controllers"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	PetController     *controllers.PetController
	FindPetController *petcontroller.FindPetController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Route("/pets", func(r chi.Router) {
			r.Get("/{id}", controllers.FindPetController.FindPet)
			r.Patch("/{petID}", controllers.PetController.Update)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/{id}/my-pets", controllers.PetController.ListUserPets)
		})
	})
}
