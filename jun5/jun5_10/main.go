package main

import (
	"fmt"
	"jun5_10/notifier"
	"time"
)

func main() {
	n := notifier.NewNotifier()

	for i := 1; i < 3; i++ {
		go func(id int) {
			fmt.Printf("[Worker %d] Выполняю\n", id)
			n.Wait(id)
			time.Sleep(1 * time.Second)
		}(i)
	}

	time.Sleep(1 * time.Second)
	n.Notice()
	time.Sleep(1 * time.Second)
	n.Notice()
	time.Sleep(2 * time.Second)
}
