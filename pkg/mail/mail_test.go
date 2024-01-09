package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfigForMailPkg(t *testing.T) {
	cfg, err := CreateConfig("example@gmail.com", "secret", "")

	assert.Nil(t, err)
	assert.Equal(t, cfg.EmailAdress, "example@gmail.com")
	assert.Equal(t, cfg.EmailSecretPassword, "secret")
	assert.Equal(t, cfg.Provider, "smtp.gmail.com")
	assert.NotNil(t, cfg)
}

func TestCreateMail(t *testing.T) {
	cfg, err := CreateConfig("example@gmail.com", "secret", "")

	assert.Nil(t, err)
	assert.Equal(t, cfg.EmailAdress, "example@gmail.com")

	mail := NewMail(cfg)
	assert.NotNil(t, mail)
}
