package routes

import (
	"pet-dex-backend/v2/api/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Controllers struct {
	PetController   *controllers.PetController
	UserController  *controllers.UserController
	OngController   *controllers.OngController
	BreedController *controllers.BreedController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))

		r.Route("/pets", func(r chi.Router) {
			r.Route("/breeds", func(r chi.Router) {
				r.Get("/", controllers.BreedController.List)
				r.Get("/{breedID}", controllers.BreedController.FindBreed)
			})

			r.Get("/breeds", controllers.BreedController.List)
			r.Patch("/{petID}", controllers.PetController.Update)
			r.Post("/", controllers.PetController.CreatePet)
		})

		r.Route("/ongs", func(r chi.Router) {
			r.Post("/", controllers.OngController.Insert)
			r.Get("/", controllers.OngController.List)
			r.Get("/{ongID}", controllers.OngController.FindByID)
			r.Patch("/{ongID}", controllers.OngController.Update)
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/{id}/my-pets", controllers.PetController.ListUserPets)
			r.Post("/token", controllers.UserController.GenerateToken)
			r.Post("/", controllers.UserController.Insert)
			r.Patch("/{id}", controllers.UserController.Update)
			r.Get("/{id}", controllers.UserController.FindByID)
		})
	})
}
