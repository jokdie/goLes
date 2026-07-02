package main

import (
	"fmt"
)

func Producer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i < 11; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func Square(inCh <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for v := range inCh {
			ch <- v * v
		}
		close(ch)
	}()

	return ch
}

func Printer(ch <-chan int) {
	for v := range ch {
		fmt.Println("Value: ", v)
	}
}

func main() {
	Printer(Square(Producer()))
}
