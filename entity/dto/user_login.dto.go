package dto

import "errors"

type UserLoginDto struct {
	Email    string `json:"Email"`
	Password string `json:"password"`
}

func (u *UserLoginDto) Validate() error {
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}
