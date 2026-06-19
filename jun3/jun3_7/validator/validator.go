package validator

import (
	"jun3_7/repository"
	"net/mail"
)

const (
	nameField  = "Name"
	emailField = "Email"
)

type ValidateUserError struct {
	Field string
}

func (v ValidateUserError) Error() string {
	switch v.Field {
	case nameField:
		return "User validation error. Field \"" + nameField + "\" is invalid"
	case emailField:
		return "User validation error. Field \"" + emailField + "\" is invalid"
	default:
		return "User validation error. Unknown field"
	}
}

func ValidateUser(u repository.User) error {
	if u.Name == "" {
		return ValidateUserError{Field: nameField}
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return ValidateUserError{Field: emailField}
	}

	return nil
}
