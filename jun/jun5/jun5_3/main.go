package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	names := []string{
		"Alex",
		"Bob",
		"Charlie",
	}

	wg.Add(len(names))
	for _, u := range names {
		go func() {
			defer wg.Done()
			fmt.Println(u)
		}()
	}
	wg.Wait()

	fmt.Println("all workers finished")
}
