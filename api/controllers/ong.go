package controllers

import (
	"encoding/json"
	"net/http"

	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase"
)

var httphandler = struct{ UseCase usecase.OngUseCase }{}

func NewHttpHandler(onguse usecase.OngUseCase) {
	httphandler.UseCase = onguse
}

func CreateOng(w http.ResponseWriter, r *http.Request) {
	var ong entity.Ong
	err := json.NewDecoder(r.Body).Decode(&ong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	println(ong.CNPJ)

	w.WriteHeader(http.StatusCreated)
}
