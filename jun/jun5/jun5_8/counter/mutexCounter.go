package counter

import "sync"

type MutexCounter struct {
	mu    sync.Mutex
	value int
}

func NewMutexCounter() *MutexCounter {
	return &MutexCounter{}
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}
