# Interview Prep — Senior Backend Engineer

Target: Senior Software Engineer (backend / architecture / DevOps).
Stack: Node.js, PHP, Go (learning).

## How to use this repo

- **Plan:** a 15-week deep plan in [`plan/`](plan/README.md). Each module has three depth levels
  (use → internals → production), a hands-on lab, primary reading, and a self-test.
- **Daily:** ~1h on the current module + ~1h on the [DSA track](plan/14-dsa-track.md) in Go.
- **Weekend:** the module's lab (`labs/`), the next [capstone](plan/capstone.md) milestone, and the self-test, reviewed with Claude.
- **Got something wrong?** Add it to [`mistakes.md`](mistakes.md). Review it every Sunday.

## Layout

| Folder | What goes there |
|---|---|
| `plan/` | The 15 modules + capstone: checklists, labs, reading, self-tests |
| `capstone/` | Notification Platform v2, built up across all modules (create it in week 1) |
| `notes/<module>/` | Your explanations in your own words |
| `labs/<module>/` | Hands-on builds (Go), including what you broke and observed |
| `system-design/` | One file per design, using [`_template.md`](system-design/_template.md) |
| `stories/` | STAR stories for behavioral rounds |
| `coding/go/` | DSA solutions with table-driven tests |
| `mistakes.md` | Everything you got wrong, and the correct answer |

## 15-week tracker

Tick a week when its module checklist, lab, and self-test are done.

| Week | Module | DSA (data structure + pattern) | Done |
|---|---|---|---|
| 1 | [01 Networking & HTTP](plan/01-networking-http.md) | Basics: Go, Big-O, warm-ups | ⬜ |
| 2 | [02 OS, Concurrency & I/O](plan/02-os-concurrency-io.md) | Hash maps (build one) + arrays & hashing | ⬜ |
| 3 | [03 Databases](plan/03-databases.md) (part 1: SQL, indexes, storage) | Sorting (build merge/quick sort) + two pointers | ⬜ |
| 4 | [03 Databases](plan/03-databases.md) (part 2: transactions, MVCC, replication) | Sliding window | ⬜ |
| 5 | [04 Caching, Redis & NoSQL](plan/04-caching-nosql.md) | Stack, queue, ring buffer | ⬜ |
| 6 | [05 Node.js Internals + PHP-FPM](plan/05-nodejs-php.md) | Binary search | ⬜ |
| 7 | [06 Go, In Depth](plan/06-go.md) | Linked lists + build an LRU cache | ⬜ |
| 8 | [07 Distributed Systems](plan/07-distributed-systems.md) | Trees, B-trees, tries | ⬜ |
| 9 | [08 Messaging & Event-Driven](plan/08-messaging.md) | Heaps + priority queues | ⬜ |
| 10 | [09 Architecture & Resilience](plan/09-architecture-resilience.md) | Graphs, topological sort, union-find | ⬜ |
| 11 | [10 System Design Practice](plan/10-system-design.md) (problems 1–5) | Dijkstra, intervals, greedy | ⬜ |
| 12 | [10 System Design Practice](plan/10-system-design.md) (problems 6–10) | Backtracking, bits, Bloom filters | ⬜ |
| 13 | [11 DevOps & Cloud](plan/11-devops-cloud.md) + [12 Observability](plan/12-observability-reliability.md) | 1-D DP | ⬜ |
| 14 | [15 Security](plan/15-security.md) | 2-D DP | ⬜ |
| 15 | [13 Behavioral & Mock Loops](plan/13-behavioral-mocks.md) | Review + timed mocks | ⬜ |

**Start here:** [Module 01](plan/01-networking-http.md) + [`coding/go/basics`](coding/go/basics/README.md). At the weekend, the first [capstone](plan/capstone.md) milestone.

## Coding progress (Go)

Run all tests from the repo root:

```bash
go test ./...
```

**New to DSA? Start with [`coding/go/basics/`](coding/go/basics/README.md)**: Go essentials, Big-O, a 5-step problem approach, and 8 warm-up exercises. Do those before the problems below.

| # | Problem | Pattern | Status |
|---|---|---|---|
| 0 | [Basics: 8 warm-ups](coding/go/basics) | Loops, maps, two pointers, sliding window | ⬜ |
| 1 | [Two Sum](coding/go/arrays/twosum) | Hash map | ⬜ |
| 2 | [Valid Anagram](coding/go/arrays/validanagram) | Counting | ⬜ |
