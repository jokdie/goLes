package main

import (
	"context"
	"errors"
	"fmt"
	"mid1_2/group"
	"time"
)

func main() {
	g := group.NewGroup(context.Background())

	g.Go(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return errors.New("worker 1. Context OFF")

			default:
				time.Sleep(55 * time.Millisecond)
				fmt.Println("worker 1 start")

				return errors.New("alarm")
			}
		}
	})
	g.Go(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return errors.New("worker 2. Context OFF")

			default:
				fmt.Println("worker 2 start")

				return nil
			}
		}
	})
	g.Go(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return errors.New("worker 2. Context OFF")

			default:
				fmt.Println("worker 3 start")
				time.Sleep(100 * time.Millisecond)
				fmt.Println("worker 3 OK")

				return nil
			}
		}
	})
	g.Go(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return errors.New("worker 2. Context OFF")

			default:
				fmt.Println("worker 4 start")
				time.Sleep(150 * time.Millisecond)

				return errors.New("alarm")
			}
		}
	})

	err := g.Wait()

	fmt.Println("err: ", err)
}
