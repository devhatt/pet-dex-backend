package controllers

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
)

type OngController struct {
	usecase *usecase.OngUsecase
	logger  config.Logger
}

func NewOngcontroller(usecase *usecase.OngUsecase) *OngController {
	return &OngController{
		usecase: usecase,
		logger:  *config.GetLogger("ong-controller"),
	}
}

func (oc *OngController) Insert(w http.ResponseWriter, r *http.Request) {
	var ongDto dto.OngInsertDto
	err := json.NewDecoder(r.Body).Decode(&ongDto)

	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = oc.usecase.Save(&ongDto)

	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (oc *OngController) FindByID(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "ongID")

	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	ong, err := oc.usecase.FindByID(ID)

	if err != nil {
		logger.Error("error on ong controller: ", err)
		if err = json.NewEncoder(w).Encode(&ong); err != nil {
			http.Error(w, "Failed to encode ong", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (oc *OngController) Update(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "id")
	ID, err := uniqueEntityId.ParseID(IDStr)

	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var ongDto dto.OngUpdateDto
	err = json.NewDecoder(r.Body).Decode(&ongDto)

	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = oc.usecase.Update(ID, &ongDto)

	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
