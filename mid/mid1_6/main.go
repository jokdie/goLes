package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for range 5 {
		wg.Go(func() {
			fmt.Println("пытаюсь захватить mutex")
			mu.Lock()
			fmt.Println("захватил mutex")
			time.Sleep(250 * time.Millisecond)
			fmt.Println("освобождаю mutex")
			mu.Unlock()
		})
	}
	wg.Wait()
}
