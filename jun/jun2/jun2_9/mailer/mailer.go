package mailer

import (
	"errors"
	"fmt"
)

var MailerSendError = errors.New("ConsoleMailerSend error")

type Mailer interface {
	Send(to string, message string) error
}

type ConsoleMailer struct{}

func (c ConsoleMailer) Send(to string, message string) error {
	if to == "" {
		return fmt.Errorf("%w: \"to\" is empty", MailerSendError)
	}
	if message == "" {
		return fmt.Errorf("%w: \"message\" is empty", MailerSendError)
	}

	fmt.Printf("Сообщение для \"%s\": %s\n", to, message)

	return nil
}
