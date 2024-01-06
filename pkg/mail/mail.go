package mail

import (
	"fmt"
	"net/smtp"
)

type EmailSendRequest struct {
	From    string
	To      []string
	Html    string
	Subject string
	Cc      []string
	Bcc     []string
	ReplyTo []string
}

type Mail struct {
	Config *Config
}

func NewMail() *Mail {
	return &Mail{
		Config: &Config{},
	}
}

func (m *Mail) SetConfig(config *Config) {
	m.Config = config
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
