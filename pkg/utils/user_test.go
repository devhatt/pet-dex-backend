package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestISValidPassword(t *testing.T) {
	cases := map[string]struct {
		pass         string
		errorMessage string
		expected     bool
	}{
		"Valid Password": {
			pass:         "Patrick123!",
			errorMessage: "Valid Password",
			expected:     true,
		},
		"Short Password": {
			pass:         "Paa1!",
			errorMessage: "Password must be at least 6 characters",
			expected:     false,
		},
		"No Upper Letter Password": {
			pass:         "patrick123!",
			errorMessage: "Password must have at least one uppercase letter",
			expected:     false,
		},
		"No Special Character Password": {
			pass:         "Patrick123",
			errorMessage: "Password must have at least one special character",
			expected:     false,
		},
		"No lower Letter Character Password": {
			pass:         "PATRICK123!",
			errorMessage: "Password must have at least one lowercase letter",
			expected:     false,
		},
		"No digit Password": {
			pass:         "Patrick!",
			errorMessage: "Password must have at least one number",
			expected:     false,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			result := IsValidPassword(test.pass)
			assert.Equal(t, test.expected, result, test.errorMessage)
		})
	}
}
