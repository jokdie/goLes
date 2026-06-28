package main

import (
	"fmt"
	"jun5_9/config"
	"sync"
)

func main() {
	l := &config.ConfigLoader{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := l.Get()

			fmt.Println(i, c)
		}()
	}
	wg.Wait()
	fmt.Println("END")
}
