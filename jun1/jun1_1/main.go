package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func main() {
	res, err := Divide(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	res, err = Divide(10, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
