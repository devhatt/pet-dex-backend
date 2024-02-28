package routes

import (
<<<<<<< HEAD
	"pet-dex-backend/v2/api/controllers"
=======
	"fmt"
	"net/http"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
	"pet-dex-backend/v2/infra/config"
>>>>>>> e51ee2a (feat: add auth middleware)

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Controllers struct {
	PetController     *controllers.PetController
	UserController    *controllers.UserController
}

func InitRoutes(controllers Controllers, c *chi.Mux) {

	c.Route("/api", func(r chi.Router) {
		r.Route("/pets", func(r chi.Router) {
			r.Get("/{id}", controllers.PetController.FindPet)
			r.Patch("/{petID}", controllers.PetController.Update)
		})

		r.Route("/ong", func(r chi.Router) {

		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/token", func(w http.ResponseWriter, r *http.Request) {
				token :=config.GetToken()
				_, tokenString, _ := token.Encode(map[string]interface{}{"user_id": 123})
				w.Header().Add("authorization", tokenString)
				w.Write([]byte(tokenString))
				
			})
		})
	})
}
