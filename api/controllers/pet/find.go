package petcontroller

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Error converting 'userID' to UUID", http.StatusBadRequest)
		return
	}
	pet, err := cntrl.UseCase.FindById(id)

	if err != nil {
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(&pet)
	w.WriteHeader(http.StatusOK)
}
