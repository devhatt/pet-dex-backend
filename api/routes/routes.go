package routes

import (
	"github.com/go-chi/chi/v5"
	"pet-dex-backend/v2/api/controllers"
)

func InitRouter(r chi.Router) {
	r.Route(
		"/api/pets",
		func(r chi.Router) {
			r.Patch("/aniversario/{id}", controllers.AtualizarAniversario)
			r.Patch("/dia_doacao/{id}", controllers.AtualizarDiadeAdocao)
		},
	)
}
