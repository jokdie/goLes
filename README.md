# Go Learning Roadmap

Current level: Middle 1
Next task: mid1_16 — Worker Pool

Weak Areas:
  - Go Memory Model (атомарные операции и memory ordering)

## Completed

- [x] mid1_1 Manual Errgroup Fundamentals
  - fail-fast coordination
  - error propagation
  - context cancellation
  - WaitGroup synchronization
  - orchestrator pattern
  - cooperative goroutine shutdown

- [x] mid1_2 sync.Once Motivation
  - one-time execution problem
  - data race on shared state
  - why bool is insufficient
  - motivation for synchronization primitives

- [x] mid1_3 Manual Errgroup
  - custom Group abstraction
  - sync.Once
  - first error wins
  - fail-fast cancellation
  - cooperative cancellation
  - responsibility separation

- [x] mid1_4 Go Runtime Overview
  - Scheduler (M:P:G)
  - Goroutine lifecycle
  - Parking / Unparking
  - Work Stealing
  - GOMAXPROCS
  - WaitGroup internals
  - sync.Once internals
  - Mutex internals (overview)
  - Channel internals (overview)
  - select internals (overview)
  - Runtime mental model

- [x] mid1_5 sync/atomic Fundamentals
  - atomic.Int32
  - Add / Load
  - Atomic vs Mutex
  - Lock-Free counters
  - CPU atomic instructions (conceptually)
  - Lost Update problem

- [x] mid1_6 Mutex Internals
  - Fast Path
  - Slow Path
  - CAS
  - Spinning
  - Parking / Unparking
  - Mutex lifecycle
  - Scheduler interaction

- [x] mid1_7 Go Memory Model
  - Happens-Before
  - Memory Visibility
  - Publication of writes
  - Synchronization primitives
  - WaitGroup memory guarantees
  - Mutex memory guarantees
  - Channel memory guarantees
  - Why Sleep is not synchronization

- [x] mid1_8 Atomic Publication Patterns
  - atomic.Pointer[T]
  - Safe publication
  - Immutable snapshot
  - Lock-free reads
  - Publication pattern
  - Atomic pointer replacement

- [x] mid1_9 Scheduler State Transitions
  - Runnable
  - Running
  - Waiting
  - Dead
  - Park / Unpark (концептуально)
  - Scheduler state transitions
  - Mutex waiting lifecycle
  - Channel waiting lifecycle
  - Timer waiting lifecycle

- [x] mid1_10 Asynchronous Preemption
  - cooperative vs asynchronous scheduling
  - CPU-bound goroutines
  - safe points
  - scheduler fairness
  - GC interaction
  - goroutine preemption
  - difference from OS thread preemption

- [x] mid1_11 Channels Fundamentals
  - unbuffered channels
  - send/receive synchronization
  - channel blocking semantics
  - happens-before via channels
  - channel as communication primitive
  - park/unpark on channel operations

- [x] mid1_12 Channel Ownership
  - channel ownership
  - single sender ownership
  - channel lifecycle
  - close semantics
  - range over channel
  - producer-consumer pattern

- [x] mid1_13 Fan-In
  - multiple producers
  - coordinator pattern
  - channel ownership in fan-in
  - completion signaling
  - chan struct{}
  - separation of data and synchronization channels

- [x] mid1_14 Fan-Out
  - multiple workers
  - single jobs channel
  - work distribution
  - one job → one worker
  - worker lifecycle
  - coordinator responsibility

- [x] mid1_15 Pipeline
  - pipeline fundamentals
  - stage responsibility
  - stage composition
  - orchestrator pattern
  - channel ownership across stages
  - sequential data processing

## Topics learned

- Fail-fast pattern
- Manual error propagation
- Orchestrator responsibility
- Group lifecycle coordination
- WaitGroup + Context composition
- Why errgroup exists
- Motivation behind sync.Once
- Why shared bool is unsafe
- One-time execution problem
- Relationship between sync.Once and errgroup
- sync.Once
- First-error semantics
- Cooperative cancellation
- Group orchestration
- Responsibility separation
- Encapsulation of concurrency primitives
- Scheduler (M:P:G)
- Goroutine lifecycle
- G / M / P responsibilities
- GOMAXPROCS
- Local vs Global Run Queue
- Work Stealing
- Parking / Unparking goroutines
- Asynchronous Preemption
- WaitGroup internals
- sync.Once internals
- Mutex internals (overview)
- Channel internals (overview)
- select internals (overview)
- Memory Model (overview)
- Escape Analysis (overview)
- Garbage Collector (overview)
- sync/atomic
- atomic.Int32
- Atomic operations
- Lost Update problem
- Lock-Free counters
- Atomic vs Mutex
- CPU atomic instructions (conceptually)
- CMutex Fast Path
- CMutex Slow Path
- CCAS (Compare-And-Swap)
- CActive Spinning
- CMutex lifecycle
- CScheduler interaction with Mutex
- Go Memory Model
- Happens-Before
- Memory Visibility
- Publication of Writes
- Synchronization Guarantees
- Why time.Sleep is not synchronization
- Atomic Publication Pattern
- atomic.Pointer[T]
- Safe Publication
- Immutable Snapshot
- Lock-Free Read Path
- Atomic Pointer Replacement
- Goroutine State Machine
- Runnable
- Running
- Waiting
- Dead
- Park / Unpark
- Scheduler State Transitions
- Waiting mechanisms
- Ready queue
- Runtime wake-up model
- Asynchronous Preemption
- Safe Points
- Goroutine Preemption
- Scheduler Fairness
- GC and Preemption Interaction
- Difference Between Goroutine and Thread Preemption
- Channels Fundamentals
- Unbuffered Channels
- Send/Receive Synchronization
- Happens-Before via Channels
- Channel Blocking Semantics
- Channel Communication Model
- Channel Ownership
- Ownership Rule
- Channel Lifecycle
- Producer–Consumer
- close() semantics
- range over channels
- Fan-In
- Coordinator Pattern
- Completion Signaling
- chan struct{}
- Разделение каналов данных и каналов синхронизации
- Владение жизненным циклом общего канала
- Fan-Out
- Multiple Workers
- Work Distribution
- Single Shared Jobs Channel
- Worker Lifecycle
- Runtime Distribution of Channel Receivers
- Responsibility Separation (Producer / Worker / Coordinator)
- Pipeline
- Stage
- Sequential Processing
- Stage Responsibility
- Pipeline Composition
- Orchestrator Pattern
- Runtime behavior in Pipeline
- Channel ownership across multiple stages
