package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
