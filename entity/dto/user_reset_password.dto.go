package dto

import "errors"

type UserChangePasswordDto struct {
	OTP          string `json:"OTP"`
	NewPass      string `json:"newPass"`
	NewPassAgain string `json:"newPassAgain"`
}

type UserOTPDto struct {
	Email string `json:"email"`
}

func (u *UserChangePasswordDto) Validate() error {
	if u.NewPass == "" {
		return errors.New("password cannot be empty")
	}
	if u.NewPass != u.NewPassAgain {
		return errors.New("new passwords do not match")
	}
	return nil
}
