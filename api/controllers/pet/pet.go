package petcontroller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase/pet"
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

	var petToBeUpdated entity.Pet
	json.NewDecoder(r.Body).Decode(&petToBeUpdated)
	err := cntrl.UseCase.Do(paramId, &petToBeUpdated)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
}
