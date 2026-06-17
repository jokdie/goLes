package service

import (
	"jun2_9/mailer"
	"regexp"
)

type UserEmailValidationError struct{}

func (UserEmailValidationError) Error() string {
	return "UserEmailValidation error"
}

type UserService struct {
	mailer mailer.Mailer
}

func New(mailer mailer.Mailer) UserService {
	return UserService{mailer: mailer}
}

var emailRegexp = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

func (s UserService) validate(email string) error {
	if !emailRegexp.MatchString(email) {
		return UserEmailValidationError{}
	}

	return nil
}

func (s UserService) Register(email string) error {
	if err := s.validate(email); err != nil {
		return err
	}

	msg := "register new user"

	if err := s.mailer.Send(email, msg); err != nil {
		return err
	}

	return nil
}
