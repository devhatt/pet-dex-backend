package controllers

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"pet-dex-backend/v2/usecase"
	"strconv"
	"strings"

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

// Add Ong to the database.
// @Summary Create Ong
// @Description Sends the Ong registration data via the request body for persistence in the database.
// @Tags Ong
// @Accept json
// @Param ongDto body dto.OngInsertDto true "Ong object information for registration"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /ongs/ [post]
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

// List List of Ong information retrieval from query parameters.
// @Summary View list of Ong.
// @Description This endpoint allows you to retrieve a list Ong organized according to query parameters..
// @Description.markdown details
// @Tags Ong
// @Produce json
// @Param limit query string false "Query limits the return of 10 data." example(10)
// @Param sortBy query string false "Property used to sort and organize displayed data" example(name)
// @Param order query string false "Data can be returned in ascending (asc) or descending (desc) order" example des" example(desc)
// @Param offset query string false "Initial position of the offset that marks the beginning of the display of the next elements" default(0)
// @Success 200 {object} dto.OngListMapper
// @Failure 400
// @Failure 500
// @Router /ongs/ [get]
func (oc *OngController) List(w http.ResponseWriter, r *http.Request) {
	// pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sortBy")
	order := r.URL.Query().Get("order")
	offsetStr := r.URL.Query().Get("offset")

	validSortBy := map[string]bool{"name": true, "address": true}
	if !validSortBy[sortBy] {
		sortBy = "name"
	}

	if strings.ToLower(order) != "asc" && strings.ToLower(order) != "desc" {
		order = "asc"
	}

	// page, err := strconv.Atoi(pageStr)
	// if err != nil || page < 1 {
	// 	page = 1
	// }

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset, _ := strconv.Atoi(offsetStr)

	ongs, err := oc.usecase.List(limit, offset, sortBy, order)

	if err != nil {
		logger.Error("error listing ongs", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ongs)
	if err != nil {
		logger.Error("error encoding json", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// FindByID Retrieves ONG information from its provided ID.
// @Summary Find ONG by ID
// @Description Retrieves ONG details based on the ONG ID provided as a parameter.
// @Tags Ong
// @Accept json
// @Produce json
// @Param ongID path string true "ID of the ONG to be retrieved"
// @Success 200 {object} dto.OngListMapper
// @Failure 400
// @Failure 500
// @Router /ongs/{ongID} [get]
func (oc *OngController) FindByID(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "ongID")

	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	ong, err := oc.usecase.FindByID(ID)
	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(&ong); err != nil {
		oc.logger.Error("error on ong controller: ", err)
		http.Error(w, "Failed to encode ong", http.StatusInternalServerError)
		return
	}
}

// Update Ong to database.
// @Summary Update Ong By ID
// @Description Updates the details of an existing Ong based on the provided Ong ID.
// @Tags Ong
// @Accept json
// @Param ongID path string true "Ong id to be updated"
// @Param ongDto body dto.OngUpdateDto true "Data to update of the Ong"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /ongs/{ongID} [patch]
func (oc *OngController) Update(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "ongID")
	ID, err := uniqueEntityId.ParseID(IDStr)

	if err != nil {
		oc.logger.Error("error on ong controller:", err)
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

// Deletes the ONG entity by ID.
// @Summary Delete ONG by ID.
// @Description Deletes the ONG corresponding to the provided ID in the request parameter.
// @Tags Ong
// @Accept json
// @Produce json
// @Param ongID path string true "ID of the ONG to be deleted"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /ongs/{ongID} [delete]
func (oc *OngController) Delete(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "ongID")
	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = oc.usecase.Delete(ID)
	if err != nil {
		oc.logger.Error("error on ong controller: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
