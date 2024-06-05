package controllers

import (
	"encoding/json"
	"net/http"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
	"pet-dex-backend/v2/usecase"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	usecase *usecase.UserUsecase
	logger  config.Logger
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{
		usecase: usecase,
		logger:  *config.GetLogger("user-controller"),
	}
}

func (uc *UserController) Insert(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserInsertDto
	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		uc.logger.Error("Error on user controller insert: ", err)
		http.Error(w, "Error decoding request ", http.StatusBadRequest)
		return
	}

	err = userDto.Validate()

	if err != nil {
		uc.logger.Error("Error on user controller insert: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.usecase.Save(userDto)

	if err != nil {
		uc.logger.Error("Error on user controller insert: ", err)
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (uc *UserController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var userLoginDto dto.UserLoginDto
	err := json.NewDecoder(r.Body).Decode(&userLoginDto)

	if err != nil {
		uc.logger.Error("error on user controller: ", err)
		http.Error(w, "Error decoding request ", http.StatusBadRequest)
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
		uc.logger.Error("[#UserController.Update] Invalid ID -> Error: ", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var userUpdateDto dto.UserUpdateDto
	err = json.NewDecoder(r.Body).Decode(&userUpdateDto)

	if err != nil {
		uc.logger.Error("[#UserController.Update] Error decoding request -> Error: ", err)
		http.Error(w, "Error decoding request ", http.StatusBadRequest)
		return
	}

	err = uc.usecase.Update(ID, userUpdateDto)

	if err != nil {
		uc.logger.Error("[#UserController.Update] Error trying to update User -> Error: ", err)
		http.Error(w, "Error trying to update User ", http.StatusBadRequest)
		return
	}

}

func (uc *UserController) FindByID(w http.ResponseWriter, r *http.Request) {
	IDStr := chi.URLParam(r, "id")

	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		uc.logger.Error("Error on user controller insert: ", err)
		http.Error(w, "Bad Request: Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := uc.usecase.FindByID(ID)

	if err != nil {
		logger.Error("error on user controller: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(&user); err != nil {
		logger.Error("error on user controller: ", err)
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userIDFromTokenStr := r.Header.Get("UserId")
	userIDFromToken, err := uniqueEntityId.ParseID(userIDFromTokenStr)
	if err != nil {
		logger.Error("[#UserController.Delete] Erro ao tentar receber o ID do token -> Erro: %v", err)
		http.Error(w, "Erro ao converter a requisição ", http.StatusBadRequest)
		return
	}

	IDStr := chi.URLParam(r, "id")
	ID, err := uniqueEntityId.ParseID(IDStr)
	if err != nil {
		logger.Error("[#UserController.Delete] Erro ao tentar converter o body da requisição -> Erro: %v", err)
		http.Error(w, "Erro ao converter a requisição ", http.StatusBadRequest)
		return
	}

	if userIDFromToken != ID {
		logger.Error("[#UserController.Delete] Erro ao tentar excluir outro usuário -> Erro: %v", err)
		http.Error(w, "Usuário não autorizado a excluir este usuário", http.StatusUnauthorized)
		return
	}

	err = uc.usecase.Delete(ID)
	if err != nil {
		logger.Error("[#UserController.Delete] Erro ao tentar deletar o usuário -> Erro: %v", err)
		http.Error(w, "Erro ao tentar atualizar o usuário ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
