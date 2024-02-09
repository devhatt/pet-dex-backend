package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/entity"
	api "pet-dex-backend/v2/infra"
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

	if userID == "" || petID == "" {
		apiError := api.NewApiError("001", "UserId or PetId not be empty or null")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&apiError)
	}

	var petToBeUpdated entity.Pet
	json.NewDecoder(r.Body).Decode(&petToBeUpdated)

	if &petToBeUpdated == nil {
		apiError := api.NewApiError("002", "No body")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&apiError)
	}

	usecase := usecase.NewUpdateUseCase(pc.repository)
	err := usecase.Do(petID, userID, &petToBeUpdated)

	if err != nil {
		fmt.Printf("Retorno usecase %s", err.Error())
		// apiError := api.NewApiError("003", err.Error())
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error trying update pet"))
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	return
}
