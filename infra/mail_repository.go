package db

import (
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
	"pet-dex-backend/v2/pkg/mail"
)

type MailRepository struct {
	mailPkg     mail.IMail
	mailMessage mail.Message
}

func NewMailRepository(mpkg mail.IMail, msg mail.Message) interfaces.Emailrepository {
	return &MailRepository{
		mailPkg:     mpkg,
		mailMessage: msg,
	}
}

/*
	compor msg com nome e email do user
	-  como confirmações de cadastro
	-  notificações e recuperação de senha
*/

// confirmação de cadastro
func (mr *MailRepository) SendConfirmationEmail(user *entity.User) error {

	to := []string{user.Email}

	msg := MailRepository{
		mailMessage: *mail.NewMessage(to, "olaaaa este é um email de confirmação!!!"),
	}

	err := mr.mailPkg.Send(&msg.mailMessage)

	if err != nil {
		fmt.Println("Error, login or password is wrong!") // msg teste
	}

	return nil
}

// notification e reset de senhas
func (mr *MailRepository) SendNotificationEmail(message string, recipient string) error {

	return nil
}
