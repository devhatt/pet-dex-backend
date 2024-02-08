package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	if userID == "" || petID == "" {
		err := fmt.Errorf("Invalid request: userID and petID must be provided in the URL parameters")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var petToBeUpdated entity.Pet
	err := json.NewDecoder(r.Body).Decode(&petToBeUpdated)

	if err != nil {
		http.Error(w, "Invalid request: could not decode pet data from request body", http.StatusBadRequest)
		return
	}

	// Specifically extract the size to be updated
	newSize := petToBeUpdated.Size

	usecase := usecase.NewUpdateUseCase(pc.repository)
	err = usecase.Do(petID, userID, &entity.Pet{Size: newSize}) // Pass only the size for update

	if err != nil {
		fmt.Printf("Error in usecase: %s", err.Error())
		http.Error(w, "Failed to update pet size", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
