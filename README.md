# Go Learning Roadmap

Current level: Junior 6
Next task: jun6_20

Weak Areas:
-

## Completed

- [x] jun6_1 First Channel
  - unbuffered channel
  - chan
  - send operation
  - receive operation
  - goroutine synchronization
  - blocking semantics

- [x] jun6_2 Buffered Channels
  - buffered channel
  - channel capacity
  - send without immediate receiver
  - blocking on full buffer
  - producer-consumer basics

- [x] jun6_3 Channel Direction
  - send-only channel
  - receive-only channel
  - channel direction
  - range over channel
  - channel ownership
  - producer closes channel

- [x] jun6_4 Select Basics
  - select
  - waiting on multiple channels
  - scheduler interaction
  - non-deterministic channel selection
  - single receive with select

- [x] jun6_5 Goroutine Leak
  - goroutine leak
  - select with send operation
  - context cancellation
  - avoiding blocked goroutines
  - worker lifecycle

- [x] jun6_6 Select with Default
  - non-blocking select
  - default case
  - try receive
  - buffered channel
  - busy waiting basics

- [x] jun6_7 Closing Channels
  - close(channel)
  - channel ownership
  - range over channel
  - graceful stream completion
  - reading until channel close

- [x] jun6_8 Comma-OK Receive
  - v, ok := <-ch
  - zero value vs closed channel
  - channel receive semantics
  - manual receive loop
  - graceful stream termination

- [x] jun6_9 Channel Ownership
  - producer / consumer separation
  - coordinator pattern
  - channel ownership
  - sync.WaitGroup.Go
  - safe channel closing
  - multiple producers

- [x] jun6_10 Fan-In Basics
  - fan-in pattern
  - merge multiple channels
  - variadic channels
  - WaitGroup.Go
  - channel ownership
  - safe merged channel closing
  - Go 1.22+ range loop variable semantics

- [x] jun6_11 Fan-Out Basics
  - fan-out pattern
  - multiple workers
  - shared jobs channel
  - producer ownership
  - worker lifecycle
  - runtime work distribution

- [x] jun6_12 Pipeline Basics
  - pipeline pattern
  - stage composition
  - channel chaining
  - stage ownership
  - independent processing stages

- [x] jun6_13 Worker Pool Basics
  - worker pool
  - fixed worker count
  - shared jobs channel
  - results channel
  - coordinator pattern
  - WaitGroup synchronization
  - safe results channel closing

- [x] jun6_14 Semaphore Basics
  - semaphore pattern
  - buffered channel as semaphore
  - limiting concurrent operations
  - permit acquisition and release
  - concurrency control
  - goroutine lifecycle
  - defer for resource release

- [x] jun6_15 Select Cancellation Pattern
  - context cancellation
  - select with send
  - producer lifecycle
  - consumer lifecycle
  - goroutine leak prevention
  - channel ownership

- [x] jun6_16 Timeout Pattern
  - time.After
  - timeout pattern
  - concurrent result waiting
  - goroutine lifecycle
  - timeout vs work cancellation
  - buffered result channel
  - goroutine leak after timeout

- [x] jun6_17 Timer Pattern
  - time.NewTimer
  - timer lifecycle
  - timer.Stop
  - managed timeout
  - difference from time.After
  - buffered result channel
  - goroutine leak prevention

- [x] jun6_18 Ticker Pattern
  - time.NewTicker
  - periodic events
  - ticker lifecycle
  - ticker.Stop
  - context-aware ticker loop
  - graceful goroutine shutdown
  - difference between Timer and Ticker

- [x] jun6_19 Multiple Tickers Coordination
  - multiple background workers
  - shared context cancellation
  - independent ticker ownership
  - worker lifecycle
  - coordinated shutdown
  - WaitGroup.Go synchronization

## Topics learned

- Отличие "состояния" от "события"
- Unbuffered Channels
- Channel Send (<-)
- Channel Receive (<-)
- Blocking Semantics
- Goroutine Synchronization Through Channels
- Buffered Channels
- Channel Capacity
- Buffered Send Semantics
- Producer / Consumer Basics
- Buffer Full Blocking
- chan<- T
- <-chan T
- Направленные каналы
- Владение каналом (Channel Ownership)
- range по каналу
- Завершение чтения после close
- select
- Ожидание нескольких каналов
- Недетерминированный выбор готового case
- Взаимодействие select с планировщиком Go Runtime
- Goroutine leak
- Отправка в select
- Использование context.Context для завершения goroutine
- Жизненный цикл worker'а
- Неблокирующий select (default) и его последствия
- Неблокирующий select
- default в select
- Busy waiting
- Когда default уместен, а когда вреден
- Отличие блокирующего и неблокирующего ожидания
- close(channel)
- Channel Ownership
- Завершение потока данных
- range по каналу
- Чтение до закрытия канала
- Поведение чтения после close
- Двухзначное получение из канала
- ok как индикатор состояния канала
- Zero value vs закрытый канал
- Ручной цикл чтения канала
- Семантика чтения после close
- Channel Ownership
- Coordinator Pattern
- Multiple Producers
- Safe Channel Closing
- Responsibility Separation
- WaitGroup.Go coordination
- Fan-In
- Merge Pattern
- Variadic Channels
- Go 1.22+ range semantics
- Concurrent Channel Merging
- Fan-Out pattern
- Multiple workers consuming from one channel
- Runtime work distribution
- Worker lifecycle
- Producer ownership of a channel
- Pipeline pattern
- Stage composition
- Independent processing stages
- Channel chaining
- Stage ownership
- Sequential concurrent processing
- Worker Pool
- Fixed Worker Count
- Shared Jobs Queue
- Coordinator Pattern
- Worker Lifecycle
- Automatic Work Distribution by Go Runtime
- Semaphore Pattern
- Buffered Channel as Semaphore
- Concurrency Limiting
- Permit Acquisition / Release
- Resource Ownership
- Limiting concurrent operations
- defer for resource release
- Отличие Worker Pool и Semaphore
- Semaphore как ограничение доступа к ресурсу
- Select Cancellation Pattern
- Context-aware producer
- Context-aware consumer
- Blocking send cancellation
- Cooperative goroutine shutdown
- Goroutine leak prevention
- Lifecycle coordination
- Timeout Pattern
- time.After
- Waiting for result or timeout
- Timeout does not stop goroutines
- Buffered result channel
- Goroutine leak after timeout
- Difference between timeout and cancellation
- time.NewTimer
- Управляемый жизненный цикл таймера
- timer.Stop
- Отличие time.After и time.NewTimer
- Буферизированный канал результата
- Связь буферизации с предотвращением утечек goroutine
- Ticker Pattern
- time.NewTicker
- Periodic event generation
- Ticker lifecycle
- ticker.Stop
- Difference between Timer and Ticker
- Runtime timer reuse
- Context-aware periodic workers
- Graceful shutdown of ticker-based goroutines
- Multiple periodic workers
- Shared context cancellation
- Independent ticker ownership
- Worker lifecycle management
- Coordinated goroutine shutdown
- Resource ownership