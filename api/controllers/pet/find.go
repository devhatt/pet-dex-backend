package petcontroller

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/usecase"

	uniqueEntity "pet-dex-backend/v2/pkg/entity"

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

	id, err := uniqueEntity.ParseID(idStr)
	if err != nil {
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
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
