package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/api/errors"
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
		fmt.Printf("Invalid request: could not decode pet data from request body %s", err.Error())
		err := errors.InvalidBody{
			Description: "The body is invalid",
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	newSize := petToBeUpdated.Size

	usecase := usecase.NewUpdateUseCase(pc.repository)
	err = usecase.Do(petID, userID, &entity.Pet{Size: newSize})

	if err != nil {
		fmt.Printf("Error in usecase: %s", err.Error())

		err := errors.ErrInvalidID{
			Description: err.Error(),
		}

		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
