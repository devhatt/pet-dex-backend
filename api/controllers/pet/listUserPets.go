package petcontroller

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/usecase"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ListUserPetsController struct {
	UseCase *usecase.PetUseCase
}

func NewListUserPetsController(usecase *usecase.PetUseCase) *ListUserPetsController {
	return &ListUserPetsController{
		UseCase: usecase,
	}
}

func (cntrl *ListUserPetsController) ListUserPets(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userID")

	userID, erro := strconv.Atoi(idStr)
	if erro != nil {
		http.Error(w, "Error converting 'userID' to int", http.StatusBadRequest)
		return
	}
	pets, err := cntrl.UseCase.ListUserPets(userID)

	if err != nil {
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(&pets)
	w.WriteHeader(http.StatusOK)
}
