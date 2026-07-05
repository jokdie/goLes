package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	errorChan := make(chan error, 1)

	wg.Go(func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("worker 1 OK")
	})
	wg.Go(func() {
		time.Sleep(150 * time.Millisecond)
		errorChan <- errors.New("worker 2 down")
	})
	wg.Go(func() {
		timer := time.NewTimer(300 * time.Millisecond)

		select {
		case <-timer.C:
			fmt.Println("worker 3 OK")
		case <-ctx.Done():
			timer.Stop()
			fmt.Println("worker 3 Context OFF")
			return
		}
	})

	err := <-errorChan
	cancel()
	fmt.Println("---> ERROR: ", err)

	wg.Wait()
}
