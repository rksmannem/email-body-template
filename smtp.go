package clients

import (
	"net/smtp"
)

type Smtp struct {
	Auth smtp.Auth
}

func NewClient(from, password, host string ) *Smtp {
	return &Smtp{
		//Auth: smtp.PlainAuth("", "mannemrks@gmail.com", "January@123", "smtp.gmail.com"),
		Auth: smtp.PlainAuth("", from, password, host),
	}
}
