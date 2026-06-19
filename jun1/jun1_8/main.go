package main

import "fmt"

type EmptyMessageError struct{}

func (EmptyMessageError) Error() string {
	return "Empty Message Error"
}

type NotificationService struct {
	notifier Notifier
}

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) error {
	// представим что тут что-то происходит
	return nil
}

type SMSNotifier struct{}

func (e SMSNotifier) Send(message string) error {
	// представим что тут что-то происходит
	return nil
}

func (s NotificationService) Notify(message string) error {
	if message == "" {
		return EmptyMessageError{}
	}

	if err := s.notifier.Send(message); err != nil {
		return err
	}

	return nil
}

func NewNotificationService(n Notifier) NotificationService {
	return NotificationService{
		notifier: n,
	}
}

func main() {
	n := NewNotificationService(EmailNotifier{})
	if err := n.Notify(""); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
	n = NewNotificationService(SMSNotifier{})
	if err := n.Notify("Hello"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}
