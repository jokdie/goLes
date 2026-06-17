package logger

import "fmt"

type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Printf("[LOG] %s\n", message)
}
