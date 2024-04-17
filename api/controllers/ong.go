package controllers

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
)

type OngController struct {
	Usecase *usecase.OngUseCase
}

func NewOngController(usecase *usecase.OngUseCase) *OngController {
	return &OngController{
		Usecase: usecase,
	}
}

func (cntrl *OngController) FindOng(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "id")

	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	ong, err := cntrl.Usecase.FindByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&ong); err != nil {
		http.Error(w, "Failed to encode Ong", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
