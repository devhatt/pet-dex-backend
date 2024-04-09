package controllers

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/usecase"
)

type BreedController struct {
	Usecase *usecase.BreedUseCase
}

func NewBreedController(usecase *usecase.BreedUseCase) *BreedController {
	return &BreedController{
		Usecase: usecase,
	}
}

func (breedControllerc *BreedController) List(responseWriter http.ResponseWriter, request *http.Request) {
	breeds, err := breedControllerc.Usecase.List()

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(breeds)
}

