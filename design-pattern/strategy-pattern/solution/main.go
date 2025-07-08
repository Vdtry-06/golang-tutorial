package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}
func (EmailNotifier) Send(message string) {
	fmt.Println("Sending email notification:", message)
}

type SMSNotifier struct{}
func (SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS notification:", message)
}

type NotificationService struct {
	notifier Notifier
}
func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

func main() {
	s := NotificationService{
		notifier: EmailNotifier{},
	}
	s.SendNotification("Hello, this is a test email!")
	s.notifier = SMSNotifier{}
	s.SendNotification("Hello, this is a test SMS!")
}