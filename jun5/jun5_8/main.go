package main

import (
	"fmt"
	"jun5_8/counter"
	"sync"
)

type Counter interface {
	Inc()
	Value() int
}

func main() {
	var c Counter = counter.NewMutexCounter()
	var rwc Counter = counter.NewRWMutexCounter()
	gw := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		gw.Add(1)
		go func() {
			defer gw.Done()
			c.Inc()
			rwc.Inc()
		}()
	}
	gw.Wait()

	fmt.Println(c.Value(), rwc.Value())
}
