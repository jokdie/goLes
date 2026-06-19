package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Inc() {
	c.Value++
}

func (c *Counter) Add(n int) {
	c.Value += n
}

func (c Counter) String() string {
	return fmt.Sprintf("Counter: %d", c.Value)
}

func main() {
	c := Counter{}

	c.Inc()
	c.Inc()
	c.Add(10)

	fmt.Println(c.String()) // 12
}
