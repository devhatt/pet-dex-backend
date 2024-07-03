package db

import (
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/mail"
)

type MailRepository struct {
	mailPkg     mail.IMail
	mailMessage mail.Message
	logger      config.Logger
}

func NewMailRepository(mpkg mail.IMail, msg mail.Message) interfaces.Emailrepository {
	return &MailRepository{
		mailPkg:     mpkg,
		mailMessage: msg,
		logger:      *config.GetLogger("mail-repository"),
	}
}

func (mr *MailRepository) SendConfirmationEmail(user *entity.User) error {

	to := []string{user.Email}

	message := mail.NewMessage(to, "olaaaa este é um email de confirmação!!!")

	err := mr.mailPkg.Send(message)

	if err != nil {
		mr.logger.Error("error on mail repository: ", err)
		return err
	}

	return nil
}

func (mr *MailRepository) SendNotificationEmail(message string, recipient string) error {

	to := []string{recipient}

	msg := mail.NewMessage(to, message)

	err := mr.mailPkg.Send(msg)

	if err != nil {
		mr.logger.Error("error on mail repository: ", err)
		return err
	}

	return nil
}

//TODO:
//ler testes
