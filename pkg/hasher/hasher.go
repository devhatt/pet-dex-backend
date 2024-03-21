package hasher

import (
	"fmt"
	"pet-dex-backend/v2/interfaces"

	"golang.org/x/crypto/bcrypt"
)

const saltRound = 8

type Hasher struct {
}

func NewHasher() interfaces.Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(key string) (string, error) {
	var err error
	var bytes []byte
	if (len(key) != 0) {
		bytes, err = bcrypt.GenerateFromPassword([]byte(key), saltRound)
	} else {
		err = fmt.Errorf("empty string given")
	}
	
	if err != nil {
		fmt.Println("#Hasher.Hash error: %w", err)
		err = fmt.Errorf("error on hashing")
		return "", err
	}
	return string(bytes), err
}

func (h *Hasher) Compare(key, toCompare string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(toCompare), []byte(key))
	return err == nil
}
