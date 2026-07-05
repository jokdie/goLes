package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, v int) {
	ch <- v
}

func coordinator(ch chan<- int) {
	wg := sync.WaitGroup{}
	sources := []int{11, 10, 1}

	for _, src := range sources {
		wg.Go(func() {
			producer(ch, src)
		})
	}
	wg.Wait()

	close(ch)
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)

	go coordinator(ch)
	consumer(ch)
}
