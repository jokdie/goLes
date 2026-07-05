package notifier

import (
	"fmt"
	"sync"
)

type Notifier struct {
	count int
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewNotifier() *Notifier {
	n := &Notifier{}
	n.cond = sync.NewCond(&n.mu)

	return n
}

func (n *Notifier) Wait(id int) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for n.count == 0 {
		fmt.Printf("[Worker %d] Сообщения не поступило\n", id)
		n.cond.Wait()
	}

	fmt.Printf("[Worker %d] Сообщения поступило\n", id)
	n.count--
}

func (n *Notifier) Notice() {
	n.mu.Lock()
	defer n.mu.Unlock()

	fmt.Println("Запускаю службы")

	n.count++
	n.cond.Signal()
}
