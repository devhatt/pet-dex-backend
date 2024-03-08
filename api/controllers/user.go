package controllers

import (
	"net/http"
	"net/mail"
	"pet-dex-backend/v2/pkg/uniqueEntityId"
)


type UserController struct {
	//ADD USECASE
}

func NewUserController() *UserController {
	return &UserController{
		
	}
}

func (uc *UserController) Find(w http.ResponseWriter, r *http.Request) {
	//_, errId := uniqueEntityId.ParseID(string(*p))
	//_, errEmail := mail.ParseAddress(string(*p))
	//isValid := !(errId != nil && errEmail != nil)
}