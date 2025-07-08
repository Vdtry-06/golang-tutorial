package main

import (
	"fmt"
)

type NotificationService struct {
	notifierType string
}

func (s NotificationService) SendNotification(message string) {
	if s.notifierType == "email" {
		fmt.Println("Sending email notification:", message)
	} else if s.notifierType == "sms" {
		fmt.Println("Sending SMS notification:", message)
	}
}

func main() {
	s := NotificationService{notifierType: "email"}
	s.SendNotification("Hello, this is a test email!")
	s.notifierType = "sms"
	s.SendNotification("Hello, this is a test SMS!")
}
