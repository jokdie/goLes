package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context, out chan<- string) {
	select {
	case <-ctx.Done():
		return
	case out <- "done":
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)

	go worker(ctx, ch)

	cancel()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-ctx.Done():
		fmt.Println("ctx done")
	}
}
