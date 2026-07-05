package main

import (
	"fmt"
	"sync"
)

func producer() <-chan int {
	jobs := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	return jobs
}

func coordinator(jobs <-chan int, workerCount int) <-chan int {
	results := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Go(func() {
			worker(jobs, results)
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func worker(ch <-chan int, results chan<- int) {
	for v := range ch {
		results <- v * v
	}
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	workerCount := 3

	consumer(coordinator(producer(), workerCount))
}
