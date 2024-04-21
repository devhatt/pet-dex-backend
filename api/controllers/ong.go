package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/usecase"
)

type OngController struct {
	ongUsecase *usecase.OngUsecase
}

func NewOngcontroller(ongUsecase *usecase.OngUsecase) *OngController {
	return &OngController{
		ongUsecase: ongUsecase,
	}
}

func (oc *OngController) Insert(w http.ResponseWriter, r *http.Request) {
	var ongDto dto.OngInsertDto
	err := json.NewDecoder(r.Body).Decode(&ongDto)

	if err != nil {
		fmt.Println(fmt.Errorf("error on ong decode: %w", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = oc.ongUsecase.Save(&ongDto)

	if err != nil {
		fmt.Println(fmt.Errorf("error saving ong: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
