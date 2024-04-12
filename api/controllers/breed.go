package controllers

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/usecase"
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

func (breedControllerc *BreedController) List(responseWriter http.ResponseWriter, request *http.Request) {
	breeds, err := breedControllerc.Usecase.List()
	if err != nil {
		logger.Error("error listing breeds", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(breeds)
}
