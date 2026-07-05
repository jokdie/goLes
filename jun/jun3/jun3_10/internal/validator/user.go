package validator

import (
	"errors"
	"jun3_10/internal/dto"
)

var (
	UserValidatorNameError = errors.New("name is required")
	UserValidatorAgeError  = errors.New("invalid age")
)

type UserValidator struct {
}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (v *UserValidator) Validate(dto *dto.CreateUserRequest) error {
	if dto.Name == "" {
		return UserValidatorNameError
	}
	if dto.Age <= 0 {
		return UserValidatorAgeError
	}

	return nil
}
