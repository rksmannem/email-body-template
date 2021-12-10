package email

import (
	"bytes"
	"net/smtp"
	"text/template"

	"github.com/email-body-template/clients"
)

// Request ...
type Request struct {
	From    string
	To      []string
	Subject string
	Body    string
}

// NewRequest ...
func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

// Send ...
func (r *Request) Send(smtpClient *clients.Smtp, host, port string) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.Body)
	//addr := "smtp.gmail.com:587"
	addr := host + ":" + port

	err := smtp.SendMail(addr, smtpClient.Auth, "mannemrks@gmail.com", r.To, msg)
	if err != nil {
		return err
	}

	return nil
}

// ParseTemplate ...
func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.Body = buf.String()
	return nil
}
