package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	countGo := 100
	countIncrement := 1000

	for i := 0; i < countGo; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < countIncrement; i++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()

	fmt.Println(counter.Value())
}
