package main

import (
	"github.com/email-body-template/config"
	"log"

	"github.com/email-body-template/clients"
	"github.com/email-body-template/email"
	"github.com/email-body-template/model"
)

func main() {

	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("fail to get config data, err: %v", err.Error())
	}
	smtpClient := clients.NewClient(cfg.Smtp)

	emailTemplate := model.EmailTemplate{
		Name: "ramakrishna",
		URL:  "http://geektrust.in",
	}

	r := email.NewRequest(cfg.To, "Hello Junk!", "Hello, World!")
	if err := r.ParseTemplate("template.html", emailTemplate); err != nil {
		log.Fatalf("fail to parse template, err: %v", err.Error())
	}

	if err := r.Send(smtpClient, cfg.Smtp.Host, cfg.Smtp.Port); err != nil {
		log.Fatalf("fail to send email using template, err: %v", err.Error())
	}
}
