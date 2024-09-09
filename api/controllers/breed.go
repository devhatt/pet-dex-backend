package controllers

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
)

var logger = config.GetLogger("breed-controller")

type BreedController struct {
	Usecase *usecase.BreedUseCase
}

func NewBreedController(usecase *usecase.BreedUseCase) *BreedController {
	return &BreedController{
		Usecase: usecase,
	}
}

// // List retrieves list of breed information for all pets.
// @Summary View list of all Breed
// @Description // List retrieves list of information of all pet breeds
// @Tags Pet
// @Produce json
// @Success 200 {object} dto.BreedList
// @Failure 400
// @Failure 500
// @Router /pets/breeds/ [get]
func (cntrl *BreedController) List(responseWriter http.ResponseWriter, request *http.Request) {
	breeds, err := cntrl.Usecase.List()
	if err != nil {
		logger.Error("error listing breeds", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	err = json.NewEncoder(responseWriter).Encode(breeds)
	if err != nil {
		logger.Error("error encoding json", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func (cntrl *BreedController) FindBreed(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "id")

	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	breed, err := cntrl.Usecase.FindByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&breed); err != nil {
		http.Error(w, "Failed to encode breed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
