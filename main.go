package main

import (
	"log"

	"github.com/email-body-template/clients"
	"github.com/email-body-template/email"
	"github.com/email-body-template/model"
)

func main() {
	emailTemplate := model.EmailTemplate{
		Name: "ramakrishna",
		URL:  "http://geektrust.in",
	}

	r := email.NewRequest([]string{"rksmannem@gmail.com", "mannemrks@gmail.com"}, "Hello Junk!", "Hello, World!")

	err := r.ParseTemplate("template.html", emailTemplate)
	if err != nil {
		log.Fatalf("fail to parse template, err: %v", err.Error())
	}

	smtpClient := clients.NewClient("mannemrks@gmail.com", "January@123", "smtp.gmail.com:587")

	if err := r.Send(smtpClient); err != nil {
		log.Fatalf("fail to send email using template, err: %v", err.Error())
	}
}
