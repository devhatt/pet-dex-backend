package controllers

import (
	"encoding/json"
	"net/http"

	"pet-dex-backend/v2/entity"
)

func HandleCreateOng(w http.ResponseWriter, r *http.Request) {
	var ong entity.Ong
	err := json.NewDecoder(r.Body).Decode(&ong)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	println(ong.CNPJ)

	w.WriteHeader(http.StatusCreated)
}
