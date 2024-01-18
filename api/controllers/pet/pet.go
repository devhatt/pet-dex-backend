package petcontroller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase/pet"
	"strconv"
)

type UpdatePetController struct {
	UseCase *pet.UpdateUseCase
}

func NewUpdatePetController(usecase *pet.UpdateUseCase) *UpdatePetController {
	return &UpdatePetController{
		UseCase: usecase,
	}
}
func (cntrl *UpdatePetController) Update(w http.ResponseWriter, r *http.Request) {

	paramId := chi.URLParam(r, "id")
	convetedId, convertErr := strconv.Atoi(paramId)
	if convertErr != nil {
		fmt.Errorf("Invalid id!")
		w.WriteHeader(400)
		return
	}

	var petToBeUpdated entity.Pet
	json.NewDecoder(r.Body).Decode(&petToBeUpdated)
	err := cntrl.UseCase.Do(convetedId, &petToBeUpdated)

	if err != nil {
		w.WriteHeader(400)
		return
	}
}
