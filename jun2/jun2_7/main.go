package main

import (
	"errors"
	"fmt"
	"jun2_7/validator"
)

func main() {
	var ageErr *validator.ValidationError

	if err := validator.ValidateAge(-15); errors.As(err, &ageErr) {
		fmt.Println(ageErr.Field)
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}
