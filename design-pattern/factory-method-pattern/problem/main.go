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

func main() {
	s := NotificationService{
		// I don't want my users init a new notifier like this.
		// They should call to something to produce a notifier with its specific type.
		// => CreateNotifier(type) Notifier
		notifier: EmailNotifier{},
	}
	s.SendNotification("Hello, this is a test email!")
}