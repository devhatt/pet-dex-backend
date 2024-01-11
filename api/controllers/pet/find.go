package petcontroller

import (
	"net/http"
	"pet-dex-backend/v2/usecase"
)

type FindPetController struct {
	UseCase *usecase.PetUseCase
}

func NewFindPetController(usecase *usecase.PetUseCase) *FindPetController {
	return &FindPetController{
		UseCase: usecase,
	}
}

func (cntrl *FindPetController) FindPet(w http.ResponseWriter, r *http.Request) {
	_, err := cntrl.UseCase.Find(1)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
}
