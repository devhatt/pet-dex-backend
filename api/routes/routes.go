package routes

import (
	"pet-dex-backend/v2/api/controllers"
	"fmt"
	"encoding/json"
	"net/http"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
	"pet-dex-backend/v2/api/middlewares"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt"
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
			r.Post("/token", func(w http.ResponseWriter, r *http.Request) {
				encoder := middlewares.NewEncoderAdapter(config.GetEnvConfig().JWT_SECRET)
				user := &interfaces.UserClaims{}
				json.NewDecoder(r.Body).Decode(&user)
				token, _ := encoder.NewAccessToken(interfaces.UserClaims{
					Id:    user.Id,
					Name:  user.Email,
					Email: user.Email,
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: time.Now().Add(time.Hour).Unix(),
					},
				})
				w.Header().Add("Authorization", token)
				json.NewEncoder(w).Encode(struct {
					Token string `json:"token"`
				}{
					Token: token,
				})
				w.WriteHeader(201)
			})
		})
	})
}
