package validator

import "fmt"

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s", e.Field)
}

func ValidateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age"}
	}

	return nil
}
