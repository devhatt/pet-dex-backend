package routes

import (
<<<<<<< HEAD
	"pet-dex-backend/v2/api/controllers"
	"fmt"
	"encoding/json"
=======
>>>>>>> 07b8768 (abstração da logica de geração de token para um usecase)
	"net/http"
	"pet-dex-backend/v2/api/controllers"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
	"pet-dex-backend/v2/api/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Controllers struct {
	PetController     *controllers.PetController
	UserController    *controllers.UserController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))

		r.Route("/pets", func(r chi.Router) {
			r.Get("/{id}", controllers.PetController.FindPet)

			r.Get("/{id}", controllers.FindPetController.FindPet)
			r.Patch("/{petID}", controllers.PetController.Update)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/token", controllers.UserController.GenerateToken)
		})
	})
}
