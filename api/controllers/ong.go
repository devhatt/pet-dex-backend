package controllers

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/usecase"
	"strconv"
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
		logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = oc.usecase.Save(&ongDto)

	if err != nil {
		logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (oc *OngController) List(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")

	validSortBy := map[string]bool{"name": true, "city": true}
	if !validSortBy[sortBy] {
		sortBy = "name"
	}

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	ongs, err := oc.usecase.List(limit, offset, sortBy, order)

	if err != nil {
		logger.Error("error listing breeds", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ongs)
}
