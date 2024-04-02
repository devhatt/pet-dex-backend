package hasher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHasher(t *testing.T) {
	var expectedType = &Hasher{}
	hasher := NewHasher()
	assert.IsTypef(t, expectedType, hasher, "error: New Hasher not returns a *Hasher{} struct", nil)
}

func TestHash(t *testing.T) {
	//ARRANGE
	cases := map[string]struct {
		input         string
		expectedError error
	}{
		"success": {
			input:         "my-pass",
			expectedError: nil,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			//ACT
			hasher := NewHasher()
			hash, err := hasher.Hash(test.input)

			//ASSERT
			assert.NotEqual(t, test.input, hash, "expected output mismatch got %s", hash, test.input)
			assert.ErrorIs(t, test.expectedError, err, "expected error mismatch")
		})
	}
}

func TestHashFail(t *testing.T) {
	//ARRANGE
	cases := map[string]struct {
		input         string
		expectedError error
	}{
		"failForEmpty": {
			input:         "",
			expectedError: fmt.Errorf("error on hashing"),
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			//ACT
			hasher := NewHasher()
			_, err := hasher.Hash(test.input)

			//ASSERT
			assert.Equal(t, test.expectedError, err, "expected error mismatch")
		})
	}
}

func TestCompare(t *testing.T) {
	//ARRANGE
	cases := map[string]struct {
		pass     string
		hash     string
		expected bool
	}{
		"success": {
			pass:     "my-pass",
			expected: true,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			//ACT
			hasher := NewHasher()
			hash, _ := hasher.Hash(test.pass)
			result := hasher.Compare(test.pass, hash)

			//ASSERT
			assert.Equal(t, test.expected, result, "expected output mismatch got %t expected %t", result, test.expected)
		})
	}
}
