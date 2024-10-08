package routes

import (
	"pet-dex-backend/v2/api/controllers"
	"pet-dex-backend/v2/api/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
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
				r.Get("/{petID}", controllers.PetController.FindPet)
				r.Post("/", controllers.PetController.CreatePet)
			})

			private.Route("/ongs", func(r chi.Router) {
				r.Post("/", controllers.OngController.Insert)
				r.Get("/", controllers.OngController.List)
				r.Get("/{ongID}", controllers.OngController.FindByID)
				r.Patch("/{ongID}", controllers.OngController.Update)
				r.Delete("/{ongID}", controllers.OngController.Delete)
			})

			private.Route("/user", func(r chi.Router) {
				r.Get("/{userID}/my-pets", controllers.PetController.ListUserPets)
				r.Patch("/{userID}/pets/{petID}", controllers.PetController.Update)
				r.Patch("/{userID}", controllers.UserController.Update)
				r.Get("/{userID}", controllers.UserController.FindByID)
				r.Delete("/{userID}", controllers.UserController.Delete)
			})
			private.Route("/settings", func(r chi.Router) {
				r.Patch("/push-notifications", controllers.UserController.UpdatePushNotificationSettings)
			})
		})

		r.Group(func(public chi.Router) {
			public.Get("/user", controllers.UserController.RetrieveUserList)
			public.Post("/user/create-account", controllers.UserController.Insert)
			public.Post("/user/{provider}/login", controllers.UserController.ProviderLogin)
			public.Post("/user/login", controllers.UserController.Login)
			public.Get("/pets/", controllers.PetController.ListAllPets)
			public.Get("/swagger/*", httpSwagger.Handler(
				httpSwagger.URL("doc.json"), //The url endpoint to API definition
			))
		})

	})
}
