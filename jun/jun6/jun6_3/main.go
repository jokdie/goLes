package main

import "fmt"

func producer(out chan<- int) {
	out <- 1
	out <- 2
	out <- 3
	close(out)
}

func consumer(in <-chan int) {
	for value := range in {
		fmt.Println(value)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
