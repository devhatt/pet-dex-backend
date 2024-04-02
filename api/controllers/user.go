package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/entity/dto"
	"pet-dex-backend/v2/usecase"
)

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
