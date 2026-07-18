package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func process(ctx context.Context, jobs <-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				goto waitWorkers

			case job, ok := <-jobs:
				if !ok {
					goto waitWorkers
				}

				wg.Go(func() {
					worker(ctx, job, out)
				})
			}
		}

	waitWorkers:
		wg.Wait()
	}()

	return out
}

func worker(ctx context.Context, job int, out chan<- int) {
	ctx, cancel := context.WithTimeout(ctx, 250*time.Millisecond)
	defer cancel()

	var duration time.Duration

	if job%2 == 0 {
		duration = 100 * time.Millisecond
	} else {
		duration = 400 * time.Millisecond
	}

	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		logStop(job)
		return
	case <-timer.C:
		select {
		case <-ctx.Done():
			logStop(job)
			return
		case out <- job * 2:
		}
	}
}

func logStop(job int) {
	fmt.Printf("job %d timed out\n", job)
}

func producer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func main() {
	ctx := context.Background()
	jobs := producer()

	result := process(ctx, jobs)

	for v := range result {
		fmt.Println(v)
	}

	fmt.Println("Done")
}
