package ongcontroller

import (
	"encoding/json"
	"net/http"

	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase"
)

type CreateOngController struct {
	UseCase *usecase.OngUseCase
}

func NewCreateOngcontroller(usecase *usecase.OngUseCase) *CreateOngController {
	return &CreateOngController{
		UseCase: usecase,
	}
}

func (cntrl *CreateOngController) CreateOng(w http.ResponseWriter, r *http.Request) {
	var ong entity.Ong

	err := json.NewDecoder(r.Body).Decode(&ong)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = cntrl.UseCase.Save(ong)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
}
