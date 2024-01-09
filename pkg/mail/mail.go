package mail

import (
	"fmt"
	"net/smtp"
)

type Mail struct {
	Config *Config
}

func NewMail(config *Config) *Mail {
	return &Mail{
		Config: config,
	}
}

func (m *Mail) Send(request *EmailSendRequest) error {
	emailValid := ValidateEmail(request.From)
	if !emailValid {
		return fmt.Errorf("invalid email address")
	}

	auth := m.Config.setAuth()
	hostAddres := m.Config.HostAddress + ":" + m.Config.HostPort
	err := smtp.SendMail(hostAddres, auth, request.From, request.To, []byte(request.Html))
	if err != nil {
		return err
	}

	return nil
}
