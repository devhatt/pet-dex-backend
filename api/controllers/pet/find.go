package petcontroller

import (
	"net/http"
	"pet-dex-backend/v2/usecase"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	idStr := chi.URLParam(r, "id")

	id, erro := strconv.Atoi(idStr)
	if erro != nil {
		http.Error(w, "Erro ao converter 'id' para int", http.StatusBadRequest)
		return
	}
	_, err := cntrl.UseCase.FindById(id)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
}
