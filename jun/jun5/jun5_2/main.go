package main

import (
	"fmt"
	"sync"
)

func worker(id int) {
	fmt.Printf("worker %d\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 4; i++ {
		wg.Go(func() {
			worker(i)
		})
	}
	wg.Wait()

	fmt.Println("all workers finished")
}
