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

var logger = config.GetLogger("user-controller")

type UserController struct {
	uusecase *usecase.UserUsecase
}

func NewUserController(uusecase *usecase.UserUsecase) *UserController {
	return &UserController{
		uusecase: uusecase,
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

	err = uc.uusecase.Save(userDto)

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
	token, err := uc.uusecase.GenerateToken(&userLoginDto)
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
		logger.Errorf("[#UserController.Update] ID Inválido -> Erro: %v", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var userUpdateDto dto.UserUpdateDto
	err = json.NewDecoder(r.Body).Decode(&userUpdateDto)

	if err != nil {
		logger.Errorf("[#UserController.Update] Erro ao tentar converter o body da requisiçao -> Erro: %v", err)
		http.Error(w, "Erro ao converter a requisição ", http.StatusBadRequest)
		return
	}

	err = uc.uusecase.Update(ID, userUpdateDto)

	if err != nil {
		logger.Errorf("[#UserController.Update] Erro ao tentar atualizar o usuário -> Erro: %v", err)
		http.Error(w, "Erro ao converter a requisição ", http.StatusBadRequest)
		return
	}

}
