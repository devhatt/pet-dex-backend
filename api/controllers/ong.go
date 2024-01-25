package ongcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase"
)

type OngController struct {
	UseCase *usecase.OngUseCase
}

func NewOngcontroller(usecase *usecase.OngUseCase) *OngController {
	return &OngController{
		UseCase: usecase,
	}
}

func (cntrl *OngController) CreateOng(w http.ResponseWriter, r *http.Request) {
	var ong entity.Ong

	err := json.NewDecoder(r.Body).Decode(&ong)
	if err != nil {
		fmt.Errorf("error on ong decode: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = cntrl.UseCase.Save(ong)
	if err != nil {
		fmt.Errorf("error saving ong: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
}
