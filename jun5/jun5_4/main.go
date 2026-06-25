package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int

	countGo := 1000

	wg.Add(countGo)
	for i := 0; i < countGo; i++ {
		go func() {
			defer wg.Done()
			counter++ // race condition demo
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
