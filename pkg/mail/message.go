package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Message struct {
	From        string
	To          []string
	Html        string
	Subject     string
	Cc          []string
	Bcc         []string
	ReplyTo     []string
	Attachments Attachment
}

type Attachment = map[string][]byte

func NewMessage(to []string, msg string) *Message {
	return &Message{
		To:   to,
		Html: msg,
	}
}

func (m *Message) AttachFile(path string) error {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, filename := filepath.Split(path)
	m.Attachments[filename] = fileBytes

	return nil
}

func (m *Message) ToBytes() []byte {
	buffer := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0

	buffer.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buffer.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.Cc) > 0 {
		buffer.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.Cc, ",")))
	}

	if len(m.Bcc) > 0 {
		buffer.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.Bcc, ",")))
	}

	buffer.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buffer)
	boundary := writer.Boundary()
	if withAttachments {
		buffer.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buffer.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buffer.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}

	buffer.WriteString(m.Html)
	if withAttachments {
		for k, v := range m.Attachments {
			buffer.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buffer.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buffer.WriteString("Content-Transfer-Encoding: base64\n")
			buffer.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buffer.Write(b)
			buffer.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buffer.WriteString("--")
	}

	return buffer.Bytes()

}
