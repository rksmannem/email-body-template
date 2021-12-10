package clients

import (
	"github.com/email-body-template/config"
	"net/smtp"
)

type Smtp struct {
	Auth smtp.Auth
}

func NewClient(smtpCfg config.Smtp) *Smtp {
	return &Smtp{
		//Auth: smtp.PlainAuth("", "mannemrks@gmail.com", "January@123", "smtp.gmail.com"),
		Auth: smtp.PlainAuth("", smtpCfg.From, smtpCfg.Password, smtpCfg.Host),
	}
}
