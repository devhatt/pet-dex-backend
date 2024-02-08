package petcontroller

type FindPetController struct {
	//UseCase *usecase.PetUseCase
}

//func NewFindPetController(usecase *usecase.PetUseCase) *FindPetController {
//	return &FindPetController{
//		UseCase: usecase,
//	}
//}

//func (cntrl *FindPetController) FindPet(w http.ResponseWriter, r *http.Request) {
//	idStr := chi.URLParam(r, "id")
//
//	id, erro := strconv.Atoi(idStr)
//	if erro != nil {
//		http.Error(w, "Erro ao converter 'id' para int", http.StatusBadRequest)
//		return
//	}
//	pet, err := cntrl.UseCase.FindById(id)
//
//	if err != nil {
//		w.WriteHeader(400)
//		return
//	}
//	json.NewEncoder(w).Encode(&pet)
//	w.WriteHeader(http.StatusOK)
//}
