package dto

import "errors"

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
	return nil
}
