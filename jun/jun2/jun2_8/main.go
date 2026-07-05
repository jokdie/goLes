package main

import (
	"errors"
	"fmt"
	"jun2_8/payment"
)

func main() {
	s := []payment.PaymentProcessor{
		payment.CryptoProcessor{},
		payment.CardProcessor{},
	}

	for _, v := range s {
		var vErr *payment.ValidationError
		var pErr *payment.PaymentError

		err := v.Process(-100)

		switch {
		case errors.As(err, &vErr):
			fmt.Printf("[ValidationError] amount is %d\n", vErr.Amount)
		case errors.As(err, &pErr):
			fmt.Printf("[PaymentError] amount is %d\n", pErr.Amount)
		default:
			fmt.Println("Another error")
		}
	}
}
