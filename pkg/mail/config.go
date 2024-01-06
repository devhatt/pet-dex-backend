package mail

import (
	"fmt"
	"net/smtp"
)

type Config struct {
	EmailAdress         string
	EmailSecretPassword string
	Provider            string
	HostAddress         string
	HostPort            string
}

func CreateConfig(emailAddress, emailSecret, provider string) (*Config, error) {
	if provider == "" {
		provider = "smtp.gmail.com"
	}

	validEmail := ValidateEmail(emailAddress)
	if !validEmail {
		return nil, fmt.Errorf("ConfigEmailAddress must be a valid email address")
	}

	return &Config{
		EmailAdress:         emailAddress,
		EmailSecretPassword: emailSecret,
		Provider:            provider,
	}, nil
}

func (c *Config) setAuth() smtp.Auth {
	auth := smtp.PlainAuth("", c.EmailAdress, c.EmailSecretPassword, c.Provider)

	return auth
}
