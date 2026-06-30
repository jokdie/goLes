package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, i int) {
	ch <- i
}

func coordinator(channels ...chan<- int) {
	wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Go(func() {
			for i := 1; i < 5; i++ {
				producer(ch, i)
			}
		})
	}

	wg.Wait()
	for _, ch := range channels {
		close(ch)
	}
}

func fanIn(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Go(func() {
			for v := range ch {
				merged <- v
			}
		})
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go coordinator(ch1, ch2)
	merged := fanIn(ch1, ch2)

	fmt.Println("Результат")
	consumer(merged)
}
