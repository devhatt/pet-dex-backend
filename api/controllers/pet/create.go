package petcontroller

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase"
)

type CreatePetController struct {
	UseCase *usecase.PetUseCase
}

func NewCreatePetController(usecase *usecase.PetUseCase) *CreatePetController {
	return &CreatePetController{
		UseCase: usecase,
	}
}

func (cntrl *CreatePetController) CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet entity.Pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	err = cntrl.UseCase.Save(pet)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
