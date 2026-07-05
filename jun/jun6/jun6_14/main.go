package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	sem := make(chan struct{}, 3)
	wg := sync.WaitGroup{}

	for _, v := range slice {
		sem <- struct{}{}

		wg.Go(func() {
			fmt.Println(v * v)
			<-sem
		})
	}

	wg.Wait()
	close(sem)
}
