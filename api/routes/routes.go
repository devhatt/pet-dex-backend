package routes

import (
	"pet-dex-backend/v2/api/controllers"
	"fmt"
	"encoding/json"
	"net/http"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
	"pet-dex-backend/v2/api/middlewares"
	"pet-dex-backend/v2/infra/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Controllers struct {
	PetController     *controllers.PetController
	UserController    *controllers.UserController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {
	tokenAuth := config.GetToken()

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
			r.Post("/token", func(w http.ResponseWriter, r *http.Request) {
				token := config.GetToken()
				_, tokenString, _ := token.Encode(map[string]interface{}{"user_id": 123})
				w.Header().Add("authorization", tokenString)
				json.NewEncoder(w).Encode(struct {
					Token string `json:"token"`
				}{
					Token: tokenString,
				})
				w.WriteHeader(201)
			})
		})
	})
}
