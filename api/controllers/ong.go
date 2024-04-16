package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/entity"
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
	var ong entity.Ong
	err := json.NewDecoder(r.Body).Decode(&ong)

	if err != nil {
		fmt.Println(fmt.Errorf("error on ong decode: %w", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = oc.ongUsecase.Save(&ong)
	if err != nil {
		fmt.Println(fmt.Errorf("error saving ong: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
