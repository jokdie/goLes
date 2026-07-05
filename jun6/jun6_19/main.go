package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, d time.Duration, n int) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("- %d - Worker: %v\n", n, t)
		case <-ctx.Done():
			fmt.Printf("- %d - Worker:Context - OFF\n", n)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Go(func() {
		d := 200 * time.Millisecond
		worker(ctx, d, 1)
	})

	wg.Go(func() {
		d := 400 * time.Millisecond
		worker(ctx, d, 2)
	})

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	wg.Wait()
}
