package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 0
		close(ch)
	}()

	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("off")
			break
		}
		fmt.Printf("v=%d ok=%v\n", v, ok)
	}
}
