package main

import (
	"errors"
	"fmt"
	"time"
)

func work() (string, error) {
	res := make(chan string, 1)
	timer := time.NewTimer(500 * time.Microsecond)

	go func() {
		time.Sleep(1000 * time.Second)
		res <- "text"
		close(res)
	}()

	select {
	case <-timer.C:
		return "", errors.New("timeout")

	case s := <-res:
		timer.Stop()

		return s, nil
	}
}

func main() {
	res, err := work()

	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println("res", res)
}
