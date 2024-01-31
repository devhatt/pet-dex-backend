package petcontroller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pet-dex-backend/v2/usecase"
)

type FindPetController struct {
	UseCase *usecase.PetUseCase
}

func NewFindPetController(usecase *usecase.PetUseCase) *FindPetController {
	return &FindPetController{
		UseCase: usecase,
	}
}

func (cntrl *FindPetController) FindPet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	pet, err := cntrl.UseCase.FindById(id)

	if err != nil {
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(&pet)
	w.WriteHeader(http.StatusOK)
}
