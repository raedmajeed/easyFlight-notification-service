package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	gomail "gopkg.in/gomail.v2"
	"log"
)

func SendEmailToUser(ctx context.Context, messages chan kafka.Message) error {
	msg := <-messages
	fmt.Printf("messages that came from kafka is %v", msg.Value)
	sender := "raedam786@gmail.com"
	password := "gtuemqhhzagxvjrh"
	recipient := "raedam786@gmail.com"

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "Test")
	m.SetBody("text/plain", "This is email body")

	d := gomail.NewDialer("smtp.gmail.com", 587, sender, password)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Could not send mail %v", err)
		return err
	} else {
		log.Printf("Email Sent Succesfully")
	}
	return nil
}
