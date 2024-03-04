package routes

import (
	"pet-dex-backend/v2/api/controllers"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	PetController *controllers.PetController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}/pets", func(r chi.Router) {
				r.Patch("/{petID}", controllers.PetController.Update)
			})
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {

		})
	})
}
