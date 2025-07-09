package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Println("Sending email notification:", message)
}

func (SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS notification:", message)
}

type NotificationService struct {
	notifier Notifier
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message)
}

// Factory function to create a notifier based on type
func CreateNotifier(notifierType string) Notifier {
	switch notifierType {
	case "sms":
		return SMSNotifier{}
	case "email":
		return EmailNotifier{}
	default:
		return nil
	}
}

func main() {
	s := NotificationService{
		notifier: CreateNotifier("email"),
	}
	s.SendNotification("Hello, this is a test email!")
}
