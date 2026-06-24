package main

import (
	"fmt"
	"sync"
)

func printMessage() {
	fmt.Println("goroutine working")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		printMessage()
	}()
	wg.Wait()

	fmt.Println("main finished")
}
