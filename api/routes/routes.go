package routes

import (
	"pet-dex-backend/v2/api/controllers"
	"pet-dex-backend/v2/api/middlewares"

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
		r.Use(middlewares.CorsMiddleware())
		r.Use(middleware.AllowContentType("application/json"))

		r.Group(func(private chi.Router) {
			private.Use(middlewares.AuthMiddleware)

			private.Route("/pets", func(r chi.Router) {
				r.Route("/breeds", func(r chi.Router) {
					r.Get("/", controllers.BreedController.List)
				})

				r.Patch("/{petID}", controllers.PetController.Update)
				r.Post("/", controllers.PetController.CreatePet)
			})

			private.Route("/ongs", func(r chi.Router) {
				r.Post("/", controllers.OngController.Insert)
				r.Get("/", controllers.OngController.List)
				r.Get("/{ongID}", controllers.OngController.FindByID)
				r.Patch("/{ongID}", controllers.OngController.Update)
			})

			private.Route("/user", func(r chi.Router) {
				r.Get("/{id}/my-pets", controllers.PetController.ListUserPets)
				r.Patch("/{id}", controllers.UserController.Update)
				r.Get("/{id}", controllers.UserController.FindByID)
				r.Delete("/{id}", controllers.UserController.Delete)
			})
			private.Route("/settings", func(r chi.Router) {
				r.Patch("/push-notifications", controllers.UserController.UpdatePushNotificationSettings)
			})
		})

		r.Group(func(public chi.Router) {
			public.Post("/user", controllers.UserController.Insert)
			public.Post("/user/google-login", controllers.UserController.GoogleLogin)
			public.Post("/user/token", controllers.UserController.GenerateToken)
			public.Get("/pets/", controllers.PetController.ListAllPets)
		})

	})
}
