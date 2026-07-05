package counter

import "sync"

type RWMutexCounter struct {
	mu    sync.RWMutex
	value int
}

func NewRWMutexCounter() *RWMutexCounter {
	return &RWMutexCounter{}
}

func (c *RWMutexCounter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *RWMutexCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.value
}
