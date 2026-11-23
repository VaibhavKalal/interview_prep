# Interview Prep — Senior Backend Engineer

Target: Senior Software Engineer (backend / architecture / DevOps).
Stack: Node.js, PHP, Go (learning).

## How to use this repo

- **Daily:** one topic session (notes) + one coding problem in Go (`coding/go`).
- **After every session:** write 3 lines in your own words in the week's notes folder.
- **Got something wrong?** Add it to [`mistakes.md`](mistakes.md). Review it every Sunday.

## Layout

| Folder | What goes there |
|---|---|
| `notes/weekN-*/` | Lessons and your own summaries |
| `system-design/` | One file per design, using [`_template.md`](system-design/_template.md) |
| `stories/` | STAR stories for behavioral rounds |
| `coding/go/` | Go solutions with table-driven tests |
| `mistakes.md` | Everything you got wrong, and the correct answer |

## 6-week plan

### Week 1 — Fundamentals
- [ ] Day 1: What happens when you hit a URL (DNS, TCP, TLS, HTTP, LB) — [notes](notes/week1-fundamentals/day1-request-lifecycle.md)
- [ ] Day 2: Processes, threads, I/O models (blocking, non-blocking, epoll)
- [ ] Day 3: HTTP deep-dive: methods, status codes, idempotency, caching headers, HTTP/1.1 vs 2 vs 3
- [ ] Day 4: TCP vs UDP, WebSockets, long polling, SSE
- [ ] Day 5: Load balancing algorithms, reverse proxies, CDNs
- [ ] Day 6: Auth basics: sessions vs JWT, OAuth2, hashing vs encryption
- [ ] Day 7: Review week + mistakes.md

### Week 2 — Databases
- [ ] SQL refresher: joins, group by, window functions
- [ ] Indexes: B-tree, composite indexes, leftmost prefix, covering indexes, `EXPLAIN`
- [ ] Transactions, ACID, isolation levels and anomalies, MVCC
- [ ] Normalization vs denormalization, N+1 problem, connection pooling
- [ ] Redis: data types, eviction, persistence, use cases
- [ ] SQL vs NoSQL (DynamoDB, MongoDB, Cassandra): when and why
- [ ] Replication and sharding basics

### Week 3 — Stack deep-dive
- [ ] Node.js event loop phases, microtasks vs macrotasks, libuv thread pool
- [ ] Node.js streams, backpressure, worker threads vs cluster, memory leaks
- [ ] PHP-FPM process model, OPcache, shared-nothing architecture
- [ ] Go: goroutines, GMP scheduler, channels vs mutexes
- [ ] Go: `context`, `select`, `errgroup`, race detector
- [ ] API design: REST, gRPC, GraphQL, pagination, versioning, error formats

### Week 4 — Distributed systems
- [ ] Caching strategies, invalidation, cache stampede
- [ ] Message queues: SQS vs Kafka vs RabbitMQ, delivery guarantees, DLQs
- [ ] Idempotency, retries with backoff and jitter, timeouts
- [ ] CAP / PACELC, consistency models
- [ ] Rate limiting algorithms (token bucket, sliding window)
- [ ] Outbox pattern, sagas, distributed locks

### Week 5 — System design practice
- [ ] Framework: requirements, estimates, API, data model, design, deep dives
- [ ] **Your notification system (SQS)** — [file](system-design/notification-system.md)
- [ ] URL shortener
- [ ] Rate limiter
- [ ] Chat system
- [ ] News feed
- [ ] Payment system

### Week 6 — DevOps, behavioral, mocks
- [ ] Docker: layers, multi-stage builds
- [ ] Kubernetes: pods, deployments, services, probes, HPA
- [ ] CI/CD, blue-green and canary deploys
- [ ] Observability: logs, metrics, traces, SLIs/SLOs
- [ ] 6–8 STAR stories written — [stories/](stories/)
- [ ] 2 full mock loops

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
