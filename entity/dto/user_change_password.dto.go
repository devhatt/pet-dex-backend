package dto

import (
	"errors"
	"pet-dex-backend/v2/pkg/utils"
)

type UserChangePasswordDto struct {
	OldPassword      string `json:"oldPassword"`
	NewPassword      string `json:"newPassword"`
	NewPasswordAgain string `json:"newPasswordAgain"`
}

func (u *UserChangePasswordDto) Validate() error {
	if u.NewPassword == "" {
		return errors.New("password cannot be empty")
	}
	if u.OldPassword == u.NewPassword {
		return errors.New("old password cannot be the same as new password")
	}
	if u.NewPassword != u.NewPasswordAgain {
		return errors.New("new passwords do not match")
	}

	if !utils.IsValidPassword(u.NewPassword) {
		return errors.New("new password must be at least 6 characters long and contain at least one uppercase letter and one special character")
	}
	return nil
}
