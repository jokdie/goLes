package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	conter := 0

	for {
		if conter == 5 {
			return
		}
		select {
		case t := <-ticker.C:
			fmt.Println(t)
			time.Sleep(350 * time.Millisecond)
			conter++
		}
	}
}
