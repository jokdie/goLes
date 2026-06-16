package payment

import "fmt"

type PaymentProcessor interface {
	Process(amount int) error
}

type ValidationError struct {
	Amount int
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("[CardProcessor] validation %d invalid", v.Amount)
}

type CardProcessor struct{}

func (c CardProcessor) Process(amount int) error {
	if amount <= 0 {
		return &ValidationError{Amount: amount}
	}

	return nil
}

type PaymentError struct {
	Amount int
}

func (p *PaymentError) Error() string {
	return fmt.Sprintf("[CryptoProcessor] validation %d invalid", p.Amount)
}

type CryptoProcessor struct{}

func (p CryptoProcessor) Process(amount int) error {
	if amount <= 0 {
		return &PaymentError{Amount: amount}
	}

	return nil
}
