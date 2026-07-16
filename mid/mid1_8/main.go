package main

import (
	"fmt"
	"mid1_8/internal/cache"
	"sync"
	"time"
)

func main() {
	c := &cache.Config{
		HTTPAdrr: ":1234",
		Timeout:  1,
		Version:  "v1",
	}
	s := cache.NewConfigStore(c)
	wg := sync.WaitGroup{}

	for range 3 {
		wg.Go(func() {
			for {
				time.Sleep(500 * time.Millisecond)
				actualCfg := s.Get()

				fmt.Println(actualCfg)
				if actualCfg.Version == "v4" {
					break
				}
			}
		})
	}

	wg.Go(func() {
		time.Sleep(1000 * time.Millisecond)
		c = &cache.Config{
			HTTPAdrr: ":8080",
			Timeout:  2,
			Version:  "v2",
		}
		s.Update(c)

		time.Sleep(1000 * time.Millisecond)
		c = &cache.Config{
			HTTPAdrr: ":9999",
			Timeout:  5,
			Version:  "v3",
		}
		s.Update(c)

		time.Sleep(1000 * time.Millisecond)
		c = &cache.Config{
			HTTPAdrr: ":4444",
			Timeout:  1,
			Version:  "v4",
		}
		s.Update(c)
	})
	wg.Wait()
}
