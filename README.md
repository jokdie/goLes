# Go Learning Roadmap

Current level: Middle 1
Next task: mid1_5

Weak Areas:
- Scheduler state transitions
- Asynchronous Preemption
- Mutex internals
- Memory Model

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
