package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/api/errors"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase"

	"pet-dex-backend/v2/pkg/uniqueEntityId"

	"github.com/go-chi/chi/v5"
)

type PetController struct {
	Usecase *usecase.PetUseCase
}

func NewPetController(usecase *usecase.PetUseCase) *PetController {
	return &PetController{
		Usecase: usecase,
	}
}

func (pc *PetController) Update(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	petID := chi.URLParam(r, "petIB")

	var petToBeUpdated entity.Pet
	err := json.NewDecoder(r.Body).Decode(&petToBeUpdated)
	defer r.Body.Close()

	if err != nil {
		fmt.Printf("Invalid request: could not decode pet data from request body %s", err.Error())
		err := errors.ErrInvalidBody{
			Description: "The body is invalid",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	newSize := petToBeUpdated.Size

	err = pc.Usecase.Update(petID, userID, &entity.Pet{Size: newSize})

	if err != nil {
		fmt.Printf("Error in usecase: %s", err.Error())

		err := errors.ErrInvalidID{
			Description: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (cntrl *PetController) ListUserPets(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "id")

	userID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		http.Error(w, "Bad Request: Invalid userID", http.StatusBadRequest)
		return
	}

	pets, err := cntrl.Usecase.ListUserPets(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&pets); err != nil {
		http.Error(w, "Failed to encode pets", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
