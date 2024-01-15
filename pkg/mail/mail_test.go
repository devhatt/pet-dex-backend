package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfigForMailPkg(t *testing.T) {
	cfg, err := CreateConfig("example@gmail.com", "secret", "stmp.gmail.com", "stmp.gmail.com", "587")

	assert.Nil(t, err)
	assert.Equal(t, cfg.EmailAdress, "example@gmail.com")
	assert.Equal(t, cfg.EmailSecretPassword, "secret")
	assert.Equal(t, cfg.Provider, "smtp.gmail.com")
	assert.Equal(t, cfg.HostAddress, "stmp.gmail.com")
	assert.Equal(t, cfg.HostPort, "587")
	assert.NotNil(t, cfg)
}

func TestCreateMail(t *testing.T) {
	cfg, err := CreateConfig("example@gmail.com", "secret", "stmp.gmail.com", "stmp.gmail.com", "587")
	assert.Nil(t, err)
	assert.Equal(t, cfg.EmailAdress, "example@gmail.com")

	mail := NewMail(cfg)
	assert.NotNil(t, mail)
}

func TestErrorOnCreateConfigMissing(t *testing.T) {
	configsParamethers := []struct {
		EmailAdress         string
		EmailSecretPassword string
		Provider            string
		HostAddress         string
		HostPort            string
		ErrorMessage        string
	}{
		{EmailAdress: "example@.com", EmailSecretPassword: "secret", Provider: "smtp.gmail.com", HostAddress: "stmp.gmail.com", HostPort: "587", ErrorMessage: "ConfigEmailAddress must be a valid email address"},
		{EmailAdress: "example@gmail.com", EmailSecretPassword: "", Provider: "smtp.gmail.com", HostAddress: "stmp.gmail.com", HostPort: "587", ErrorMessage: "Secret password cannot be empty"},
		{EmailAdress: "example@gmail.com", EmailSecretPassword: "secret", Provider: "", HostAddress: "stmp.gmail.com", HostPort: "587", ErrorMessage: "Provider cannot be empty"},
		{EmailAdress: "example@gmail.com", EmailSecretPassword: "secret", Provider: "smtp.gmail.com", HostAddress: "", HostPort: "587", ErrorMessage: "Host address cannot be empty"},
		{EmailAdress: "example@gmail.com", EmailSecretPassword: "secret", Provider: "smtp.gmail.com", HostAddress: "stmp.gmail.com", HostPort: "", ErrorMessage: "Host port cannot be empty"},
	}

	for _, config := range configsParamethers {
		cfg, err := CreateConfig(config.EmailAdress, config.EmailSecretPassword, config.Provider, config.HostAddress, config.HostPort)

		assert.Nil(t, cfg)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Error(t, err, config.ErrorMessage)
	}
}
