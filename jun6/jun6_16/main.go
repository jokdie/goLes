package main

import (
	"errors"
	"fmt"
	"time"
)

func work() (string, error) {
	result := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)

		result <- "Работа выполнена"
	}()

	select {
	case res := <-result:
		return res, nil

	case <-time.After(1 * time.Second):
		return "", errors.New("operation timed out")
	}
}

func main() {
	res, err := work()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
