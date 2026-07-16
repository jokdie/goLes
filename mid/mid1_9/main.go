package main

import (
	"fmt"
	"sync"
	"time"
)

func blockMu(mu *sync.Mutex, v int) {
	fmt.Printf("[G%d] trying to lock mutex\n", v)
	mu.Lock()
	fmt.Printf("[G%d] acquired mutex\n", v)
	mu.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	ch := make(chan int)

	wg.Go(func() {
		fmt.Println("[G1] waiting for mutex")
		blockMu(&mu, 1)
	})
	wg.Go(func() {
		fmt.Println("[G2] waiting for mutex")
		blockMu(&mu, 2)
	})
	wg.Go(func() {
		fmt.Println("[G3] waiting for channel")
		<-ch
		fmt.Println("[G3] finished")
	})
	wg.Go(func() {
		fmt.Println("[G4] sleeping")
		time.Sleep(5 * time.Second)
		fmt.Println("[G4] wake up")
	})
	wg.Go(func() {
		fmt.Println("[G5] sending to channel")
		ch <- 10
		fmt.Println("[G5] finished")
	})
	wg.Wait()

	fmt.Println("end")
}

G1
| Событие          | Состояние |
| ---------------- | --------- |
| Создана          | Runnable  |
| Scheduler выбрал | Running   |
| Дошла до мьютекса| Running (тк он свободный был в этот момент)   |
| Прошла критическую зону | Running   |
| Освободила мьютекс | Running   |
| Закончила | Dead   |

G2
| Событие          | Состояние |
| ---------------- | --------- |
| Создана          | Runnable  |
| Scheduler выбрал | Running   |
| Дошла до мьютекса| Running (spinning) (но не пошла парковаться, а начала спининг, учитывая что критическая область примитивная то она не паркуется)   |
| Mutex был освобождён. Runtime перевёл ожидающую goroutine в Runnable | Runnable   |
| Scheduler выбрал | Running   |
| Дошла до мьютекса| Running (тк он свободный был в этот момент)   |
| Прошла критическую зону | Running   |
| Освободила мьютекс | Running   |
| Закончила | Dead   |

G3 // сценарий когда писателя еще нет
| Событие          | Состояние |
| ---------------- | --------- |
| Создана          | Runnable  |
| Scheduler выбрал | Running   |
| Дошла до чтения канала в которой еще никто не написал, рантайм припарковал ее | Waiting |
| Chanal увидел что событие ожидания выполненно и разбудил | Runnable   |
| Scheduler выбрал | Running   |
| Закончила | Dead   |

G3 // сценарий когда писатель есть
| Событие          | Состояние |
| ---------------- | --------- |
| Создана          | Runnable  |
| Scheduler выбрал | Running   |
| Закончила | Dead   |

G4
| Событие          | Состояние |
| ---------------- | --------- |
| Создана          | Runnable  |
| Scheduler выбрал | Running   |
| Дошла до слипа и уснула, рантайм припарковал ее | Waiting |
| Time увидел что событие ожидания выполненно и разбудил | Runnable   |
| Scheduler выбрал | Running   |
| Закончила | Dead   |

G5 // сценарий когда читателя еще нет
| Событие          | Состояние |
| ---------------- | --------- |
| Создана          | Runnable  |
| Scheduler выбрал | Running   |
| Дошла до записи канала в который еще никто не читает, рантайм припарковал ее | Waiting |
| Chanal увидел что событие ожидания выполненно и разбудил | Runnable   |
| Scheduler выбрал | Running   |
| Закончила | Dead   |

G5 // сценарий когда читателя есть
| Событие          | Состояние |
| ---------------- | --------- |
| Создана          | Runnable  |
| Scheduler выбрал | Running   |
| Закончила | Dead   |