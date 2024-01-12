package petcontroller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase"
	"strconv"
)

type UpdatePetController struct {
	UseCase *usecase.UpdateUseCase
}

func NewUpdatePetController(usecase *usecase.UpdateUseCase) *UpdatePetController {
	return &UpdatePetController{
		UseCase: usecase,
	}
}
func (cntrl *UpdatePetController) UpdateSize(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	convetedId, convertErr := strconv.Atoi(id)
	if convertErr != nil {
		fmt.Errorf("Invalid id!")
		w.WriteHeader(400)
		w.Write([]byte("Invalid Id"))
	}

	var petToBeUpdated entity.Pet
	json.NewDecoder(r.Body).Decode(&petToBeUpdated)
	err := cntrl.UseCase.Do(convetedId, petToBeUpdated.PetDetails.Size)

	if err != nil {
		fmt.Errorf("Invalid id!")
		w.WriteHeader(400)
		w.Write([]byte("Invalid Id"))
	}
}
