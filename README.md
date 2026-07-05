# Go Learning Roadmap

Current level: Middle 1
Next task: mid1_4

Weak Areas:
-

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
