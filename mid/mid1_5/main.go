package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int32
	wg := sync.WaitGroup{}
	workers := 1000
	iterations := 100000

	for range workers {
		wg.Go(func() {
			for range iterations {
				counter.Add(1)
			}
		})
	}

	wg.Wait()

	actual := counter.Load()
	expected := workers * iterations

	fmt.Println(actual)
	fmt.Println(expected)
	fmt.Println(int(actual) == expected)
}
