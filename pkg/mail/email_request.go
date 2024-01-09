package mail

type EmailSendRequest struct {
	From    string
	To      []string
	Html    string
	Subject string
	Cc      []string
	Bcc     []string
	ReplyTo []string
}
