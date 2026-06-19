package main

import "fmt"

type ValidationError struct {
	Field string
}

func (ve ValidationError) Error() string {
	switch ve.Field {
	case "password":
		return "field password is invalid"
	case "email":
		return "field email is invalid"
	default:
		return "another field is invalid"
	}
}

func ValidatePassword(password string) error {
	if len([]rune(password)) < 8 {
		ve := ValidationError{Field: "password"}

		return ve
	}

	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		ve := ValidationError{Field: "email"}

		return ve
	}

	return nil
}

func main() {
	if err := ValidateEmail(""); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

	if err := ValidateEmail("test@test.com"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

	if err := ValidatePassword(""); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

	if err := ValidatePassword("qweqweqwe"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

	if err := ValidatePassword("qwe"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
