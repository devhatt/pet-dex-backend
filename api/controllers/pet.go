package controllers

import (
	"encoding/json"
	"net/http"
	api "pet-dex-backend/v2/api/erros"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
)

type PetController struct {
	repository interfaces.PetRepository
}

func NewPetController(db interfaces.PetRepository) *PetController {
	return &PetController{
		repository: db,
	}
}

func (pc *PetController) Update(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	petID := chi.URLParam(r, "petID")

	var petToBeUpdated entity.Pet
	err := json.NewDecoder(r.Body).Decode(&petToBeUpdated)

	if err != nil {
		apiError := api.NewApiError("002", "No body")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	usecase := usecase.NewUpdateUseCase(pc.repository)
	err = usecase.Do(petID, userID, &petToBeUpdated)

	if err != nil {
		apiError := api.NewApiError("003", err.Error())
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(apiError)
		return
	}
	w.WriteHeader(200)
	return
}
