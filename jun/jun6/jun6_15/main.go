package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func producer(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
				fmt.Println("-> Отправляю в канал: ", i)
			case <-ctx.Done():
				fmt.Println("---> Завершаю отправку <---")
				return
			}
		}
	}()

	return ch
}

func consumer(ctx context.Context, ch <-chan int) {
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println("<- Получено из канал: ", v)
		case <-ctx.Done():
			fmt.Println("OFF")

			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	ch := producer(ctx)

	timer := time.NewTimer(4 * time.Microsecond)
	defer timer.Stop()
	go func() {
		<-timer.C
		cancel()
	}()

	wg.Go(func() {
		consumer(ctx, ch)
	})
	wg.Wait()
}
