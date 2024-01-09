package mail

type MailInterface interface {
	Send(request *EmailSendRequest) error
}
