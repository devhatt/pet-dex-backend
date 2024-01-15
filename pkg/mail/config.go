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

func (c *Config) validate() error {
	if !ValidateEmail(c.EmailAdress) {
		return fmt.Errorf("ConfigEmailAddress must be a valid email address")
	}

	if c.EmailSecretPassword == "" {
		return fmt.Errorf("Secret password cannot be empty")
	}

	if c.Provider == "" {
		return fmt.Errorf("Provider cannot be empty")
	}

	if c.HostAddress == "" {
		return fmt.Errorf("Host address cannot be empty")
	}

	if c.HostPort == "" {
		return fmt.Errorf("Host port cannot be empty")
	}

	return nil
}

func CreateConfig(emailAddress, emailSecret, provider, hostAddress, hostPort string) (*Config, error) {
	cfg := Config{
		EmailAdress:         emailAddress,
		EmailSecretPassword: emailSecret,
		Provider:            provider,
		HostAddress:         hostAddress,
		HostPort:            hostPort,
	}

	err := cfg.validate()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) setAuth() smtp.Auth {
	auth := smtp.PlainAuth("", c.EmailAdress, c.EmailSecretPassword, c.Provider)

	return auth
}
