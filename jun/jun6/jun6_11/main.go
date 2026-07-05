package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int) {
	for i := 1; i < 11; i++ {
		ch <- i
	}

	close(ch)
}

func consumer(jobs <-chan int, id int) {
	for j := range jobs {
		fmt.Printf("worker %d processed %d\n", id, j)
	}
}

func main() {
	jobs := make(chan int)
	wg := sync.WaitGroup{}

	go producer(jobs)

	for i := 1; i < 4; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			consumer(jobs, id)
		}(i)
	}
	wg.Wait()

	fmt.Println("end")
}
