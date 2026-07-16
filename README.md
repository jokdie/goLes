# Go Learning Roadmap

Current level: Middle 1
Next task: mid1_9

Weak Areas:
- Scheduler state transitions
- Asynchronous Preemption
- Atomic publication patterns

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
