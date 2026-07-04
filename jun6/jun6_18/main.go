package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	wg.Go(func() {
	Loop:
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)

			case <-ctx.Done():
				fmt.Println("OFF")
				break Loop
			}
		}
	})

	wg.Wait()
	ticker.Stop()
}
