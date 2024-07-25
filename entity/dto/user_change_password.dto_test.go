package dto

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangePasswordValidate(t *testing.T) {
	cases := map[string]struct {
		userChangePassword UserChangePasswordDto
		expected           error
	}{
		"Valid Password": {
			userChangePassword: UserChangePasswordDto{
				OldPassword:      "oldPassword123!",
				NewPassword:      "NewPassword123!",
				NewPasswordAgain: "NewPassword123!",
			},
			expected: nil,
		},
		"Empty New Password": {
			userChangePassword: UserChangePasswordDto{
				OldPassword:      "oldPassword123!",
				NewPassword:      "",
				NewPasswordAgain: "NewPassword123!",
			},
			expected: errors.New("password cannot be empty"),
		},
		"New Password is equal to Old Password": {
			userChangePassword: UserChangePasswordDto{
				OldPassword:      "oldPassword123!",
				NewPassword:      "oldPassword123!",
				NewPasswordAgain: "NewPassword123!",
			},
			expected: errors.New("old password cannot be the same as new password"),
		},
		"New Passwords does not match": {
			userChangePassword: UserChangePasswordDto{
				OldPassword:      "oldPassword123!",
				NewPassword:      "NewPassword123!",
				NewPasswordAgain: "DifferentNewPassword123!",
			},
			expected: errors.New("new passwords do not match"),
		},
		"New Password does not match the security requirements": {
			userChangePassword: UserChangePasswordDto{
				OldPassword:      "oldPassword123!",
				NewPassword:      "password123!",
				NewPasswordAgain: "password123!",
			},
			expected: errors.New("new password must be at least 6 characters long and contain at least one uppercase letter and one special character"),
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			result := test.userChangePassword.Validate()
			assert.Equal(t, test.expected, result, test.expected)
		})
	}
}
