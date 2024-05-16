package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
)

var loggerUserController = config.GetLogger("user-controller")

type UserController struct {
	usecase *usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{
		usecase: usecase,
	}
}

func (uc *UserController) Insert(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserInsertDto
	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		fmt.Println(fmt.Errorf("#UserController.Insert error: %w", err))
		http.Error(w, "Erro ao converter requisição ", http.StatusBadRequest)
		return
	}

	err = userDto.Validate()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.usecase.Save(userDto)

	if err != nil {
		fmt.Println(fmt.Errorf("#UserController.Save error: %w", err))
		http.Error(w, "Erro ao salvar usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (uc *UserController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var userLoginDto dto.UserLoginDto
	err := json.NewDecoder(r.Body).Decode(&userLoginDto)

	if err != nil {
		fmt.Println(fmt.Errorf("#UserController.GenerateToken error: %w", err))
		http.Error(w, "Erro ao converter requisição ", http.StatusBadRequest)
		return
	}
	err = userLoginDto.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := uc.usecase.GenerateToken(&userLoginDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Add("Authorization", token)
	json.NewEncoder(w).Encode(struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
	w.WriteHeader(201)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "id")
	ID, err := uniqueEntityId.ParseID(IDStr)

	if err != nil {
		loggerUserController.Errorf("[#UserController.Update] ID Inválido -> Erro: %v", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var userUpdateDto dto.UserUpdateDto
	err = json.NewDecoder(r.Body).Decode(&userUpdateDto)

	if err != nil {
		loggerUserController.Errorf("[#UserController.Update] Erro ao tentar converter o body da requisiçao -> Erro: %v", err)
		http.Error(w, "Erro ao converter a requisição ", http.StatusBadRequest)
		return
	}

	err = uc.usecase.Update(ID, userUpdateDto)

	if err != nil {
		loggerUserController.Errorf("[#UserController.Update] Erro ao tentar atualizar o usuário -> Erro: %v", err)
		http.Error(w, "Erro ao tentar atualizar o usuário ", http.StatusBadRequest)
		return
	}

}

func (uc *UserController) FindByID(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "id")

	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := uc.usecase.FindByID(ID)

	if err != nil {
		logger.Error("error on user controller: ", err)
		if err = json.NewEncoder(w).Encode(&user); err != nil {
			http.Error(w, "Failed to encode user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
