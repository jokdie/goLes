package main

import (
	"errors"
	"fmt"
	"jun2_9/mailer"
	userService "jun2_9/service"
)

func checkErr(err error) {
	var emailErr *userService.UserEmailValidationError

	switch {
	case errors.Is(err, mailer.MailerSendError):
		fmt.Println("MailerSendError:", err)
	case errors.As(err, &emailErr):
		fmt.Println("UserEmailValidationError:", err)
	default:
		fmt.Println("Unknown error:", err)
	}
}

func main() {
	us := userService.New(mailer.ConsoleMailer{})

	if err := us.Register(""); err != nil {
		checkErr(err)
	}
	if err := us.Register("testmail.ru"); err != nil {
		checkErr(err)
	}
	if err := us.Register("test@mail"); err != nil {
		checkErr(err)
	}
	if err := us.Register("test@mail.ru"); err != nil {
		checkErr(err)
	}
}
