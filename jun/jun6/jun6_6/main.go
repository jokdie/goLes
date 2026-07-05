package main

import "fmt"

func tryReceive(ch <-chan int) {
	select {
	case v := <-ch:
		fmt.Println("received:", v)
	default:
		fmt.Println("no value")
	}
}

func main() {
	ch := make(chan int, 1)

	tryReceive(ch)

	ch <- 42

	tryReceive(ch)
}
