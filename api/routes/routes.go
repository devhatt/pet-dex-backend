package routes

import (
	"pet-dex-backend/v2/api/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Controllers struct {
	PetController  *controllers.PetController
	UserController *controllers.UserController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))

		r.Route("/pets", func(r chi.Router) {
			r.Get("/{id}", controllers.PetController.FindPet)

			r.Patch("/{petID}", controllers.PetController.Update)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/token", controllers.UserController.GenerateToken)
			r.Post("/", controllers.UserController.Insert)
		})
	})
}
