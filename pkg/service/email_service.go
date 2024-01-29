package service

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	gomail "gopkg.in/gomail.v2"
	"log"
)

type EmailMessage struct {
	Email   string
	Subject string
	Content string
}

func SendEmailToUser(message kafka.Message, senderMail, passwordAdmin string) {
	sender := senderMail
	password := passwordAdmin

	var email EmailMessage
	err := json.Unmarshal(message.Value, &email)
	if err != nil {
		log.Println("unable to unmarshal message")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", email.Email)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Content)

	log.Println("message sent to ", email.Email)
	d := gomail.NewDialer("smtp.gmail.com", 587, sender, password)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Could not send mail %v", err)
	} else {
		log.Printf("Email Sent Succesfully")
	}
}
