package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestISValidPassword(t *testing.T) {
	cases := map[string]struct {
		pass     string
		expected bool
	}{
		"Valid Password": {
			pass:     "Patrick123!",
			expected: true,
		},
		"Short Password": {
			pass:     "Paa1!",
			expected: false,
		},
		"No Upper Letter Password": {
			pass:     "patrick123!",
			expected: false,
		},
		"No Special Character Password": {
			pass:     "patrick123",
			expected: false,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			result := IsValidPassword(test.pass)
			//ASSERT
			assert.Equal(t, test.expected, result, "expected output mismatch got %t expected %t", result, test.expected)
		})
	}
}
