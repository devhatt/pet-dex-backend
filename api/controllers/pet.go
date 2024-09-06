package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/api/errors"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/usecase"
	"strconv"
	"strings"
	"time"

	"pet-dex-backend/v2/pkg/encoder"
	"pet-dex-backend/v2/pkg/uniqueEntityId"

	"github.com/go-chi/chi/v5"
)

type PetController struct {
	Usecase *usecase.PetUseCase
}

func NewPetController(usecase *usecase.PetUseCase) *PetController {
	return &PetController{
		Usecase: usecase,
	}
}

// Update Pet to the database.
// @Summary Update an Pet existing.
// @Description Update the Pet's registration data via the request body for persistence in the database.
// @Tags User
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param petID path string true "Pet ID"
// @Param petDto body dto.PetUpdateDto true "Pet object information for update of data"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /user/{userID}/pets/{petID} [patch]
func (pc *PetController) Update(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	petID := chi.URLParam(r, "petID")

	var petUpdateDto dto.PetUpdateDto
	err := json.NewDecoder(r.Body).Decode(&petUpdateDto)
	defer r.Body.Close()

	if err != nil {
		fmt.Printf("Invalid request: could not decode pet data from request body %s", err.Error())
		err := errors.ErrInvalidBody{
			Description: "The body is invalid",
		}

		w.WriteHeader(http.StatusBadRequest)
		json_err := json.NewEncoder(w).Encode(err)
		if json_err != nil {
			logger.Error("error encoding json", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	err = pc.Usecase.Update(petID, userID, petUpdateDto)

	if err != nil {
		fmt.Printf("Error in usecase: %s", err.Error())

		err := errors.ErrInvalidID{
			Description: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		json_err := json.NewEncoder(w).Encode(err)
		if json_err != nil {
			logger.Error("error encoding json", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}

// FindPet Retrieves Pet information from its provided ID.
// @Summary Find Pet by ID
// @Description Retrieves Pet details based on the pet ID provided as a parameter.
// @Tags Pet
// @Accept json
// @Produce json
// @Param petID path string true "ID of the Pet to be retrieved"
// @Success 200 {object} entity.Pet
// @Failure 400
// @Failure 500
// @Router /pets/{petID} [get]
func (cntrl *PetController) FindPet(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "petID")

	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	pet, err := cntrl.Usecase.FindByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&pet); err != nil {
		http.Error(w, "Failed to encode pet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// List all pets from a provided user id.
// @Summary List pets by user id
// @Description List all pets owned by the user corresponding to the provided user ID
// @Tags User
// @Accept json
// @Produce json
// @Param userID path string true "ID of the User"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /user/{userID}/my-pets [get]
func (cntrl *PetController) ListUserPets(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "userID")

	userID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		http.Error(w, "Bad Request: Invalid userID", http.StatusBadRequest)
		return
	}

	pets, err := cntrl.Usecase.ListUserPets(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&pets); err != nil {
		http.Error(w, "Failed to encode pets", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Add Pet to the database.
// @Summary Create Pet by petDto
// @Description Sends the Pet's registration data via the request body for persistence in the database.
// @Tags Pet
// @Accept json
// @Produce json
// @Param petDto body dto.PetInsertDto true "Pet object information for registration"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /pets/ [post]
func (cntrl *PetController) CreatePet(w http.ResponseWriter, r *http.Request) {
	var petToSave dto.PetInsertDto

	err := json.NewDecoder(r.Body).Decode(&petToSave)
	defer r.Body.Close()

	if err != nil {
		fmt.Printf("Invalid request: could not decode pet data from request body %s", err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json_err := json.NewEncoder(w).Encode(errors.ErrInvalidBody{
			Description: "The body is invalid",
		})
		if json_err != nil {
			logger.Error("error encoding json", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = petToSave.Validate()
	if err != nil {
		fmt.Printf("Invalid request: could not validate pet data from request body %s", err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json_err := json.NewEncoder(w).Encode(err)
		if json_err != nil {
			logger.Error("error encoding json", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = cntrl.Usecase.Save(petToSave)

	if err != nil {
		fmt.Printf("Error in usecase: %s", err.Error())

		err := err.Error()

		w.WriteHeader(http.StatusBadRequest)
		json_err := json.NewEncoder(w).Encode(err)
		if json_err != nil {
			logger.Error("error encoding json", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ListAllPets Retrieves the list of all pets.
// @Summary View list of all pets.
// @Description Public route for viewing all pets.
// @Tags Pet
// @Produce json
// @Success 200 {object} entity.Pet
// @Failure 400
// @Failure 500
// @Router /pets/ [get]
func (cntrl *PetController) ListAllPets(w http.ResponseWriter, r *http.Request) {
	encoderAdapter := encoder.NewEncoderAdapter(config.GetEnvConfig().JWT_SECRET)
	var pageNumber int
	var err error
	var pets []*entity.Pet
	pageStr := r.URL.Query()

	if pageStr.Get("page") == "" {
		pageNumber = 1
	} else {
		pageNumber, err = strconv.Atoi(pageStr.Get("page"))
	}

	if err != nil {
		http.Error(w, "Bad Request: Invalid page number", http.StatusBadRequest)
		return
	}

	if pageNumber < 0 {
		http.Error(w, "Bad Request: Page number cannot be negative", http.StatusBadRequest)
		return
	}

	authHeader := r.Header.Get("Authorization")
	isUnauthorized := true

	headerSplited := strings.Split(authHeader, " ")
	if len(headerSplited) == 2 {
		bearerToken := headerSplited[1]

		userclaims := encoderAdapter.ParseAccessToken(bearerToken)
		isUnauthorized = userclaims.ExpiresAt != 0 && userclaims.ExpiresAt < time.Now().Unix()
	}

	pets, err = cntrl.Usecase.ListPetsByPage(pageNumber, isUnauthorized)

	if err != nil {
		http.Error(w, "Failed to retrieve pets", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(&pets); err != nil {
		http.Error(w, "Failed to encode pets", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
