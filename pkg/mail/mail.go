package mail

import (
	"fmt"
	"net/smtp"
)

type Mail struct {
	Config *Config
}

type IMail interface {
	Send(message *Message) error
}

func NewMail(config *Config) *Mail {
	return &Mail{
		Config: config,
	}
}

func (m *Mail) Send(message *Message) error {
	emailValid := ValidateEmail(message.From)
	if !emailValid {
		return fmt.Errorf("invalid email address")
	}

	auth := m.Config.setAuth()
	hostAddres := fmt.Sprintf("%s:%s", m.Config.HostAddress, ":"+m.Config.HostPort)
	err := smtp.SendMail(hostAddres, auth, message.From, message.To, message.ToBytes())
	if err != nil {
		return err
	}

	return nil
}
