# The Deep Plan (15 weeks)

This plan has two goals, and they reinforce each other:

1. **Clear senior backend interviews.** Every module ends in interview-style self-tests.
2. **Become a genuinely deeper engineer.** Understanding that survives the third "why?"
   follow-up question is what interviewers look for, and it's also what makes you better at
   your job: faster debugging, better designs, and better decisions.

So each module goes beyond interview answers: you build things, break them, read primary
sources, and tie everything back to systems you work on.

## The depth method (use it for every topic)

Reading alone gives shallow knowledge. For each module, go through all five steps:

1. **Learn:** read the primary source (book chapter, official docs), not just blog summaries.
2. **Build:** do the lab. Implementing a thing in Go teaches what reading can't.
3. **Break:** make it fail on purpose (kill a process, add latency, fill a queue) and observe.
4. **Explain:** answer the module's self-test out loud or in writing, without notes.
   If you can't explain it simply, you don't know it yet.
5. **Defend:** send your answers to Claude for an interviewer-style review with follow-ups.
   Log what you got wrong in `mistakes.md`.

## Three levels per topic

Each module's checklist is split into levels. Senior interviews probe L2 and L3.

| Level | Meaning | Example (database indexes) |
|---|---|---|
| **L1 — Use** | Know what it is and how to use it | "Add an index on `email`." |
| **L2 — Internals** | Know how it works underneath | "It's a B+tree; lookups are O(log n); a composite index follows the leftmost-prefix rule." |
| **L3 — Production** | Know tradeoffs, failure modes, and how to debug it | "Indexes slow writes and bloat; here's how I found an unused index with `pg_stat_user_indexes` and a slow query with `EXPLAIN ANALYZE`." |

## What every module contains

- **Checklist** in three levels (L1 use → L2 internals → L3 production)
- **Lab:** build it, then break it on purpose
- **Reading:** specific chapters from primary sources
- **Self-test:** interview-style questions to answer without notes
- **Beyond the interview:** papers, source code, courses, and real-world postmortems for lasting depth.
  Optional during the plan; a great backlog after it.

## Time budget

About **2 hours on weekdays** (1h module + 1h DSA) and **4–5 hours on one weekend day**
(lab + capstone milestone + self-test review). Falling behind is fine; skipping the lab and self-test is not.

## Modules

| Weeks | Phase | Module |
|---|---|---|
| 1 | Foundations | [01 — Networking & HTTP](01-networking-http.md) |
| 2 | Foundations | [02 — OS, Concurrency & I/O](02-os-concurrency-io.md) |
| 3–4 | Data | [03 — Relational Databases, In Depth](03-databases.md) |
| 5 | Data | [04 — Caching, Redis & NoSQL](04-caching-nosql.md) |
| 6 | Runtimes | [05 — Node.js Internals (+ PHP-FPM)](05-nodejs-php.md) |
| 7 | Runtimes | [06 — Go, In Depth](06-go.md) |
| 8 | Distributed | [07 — Distributed Systems Fundamentals](07-distributed-systems.md) |
| 9 | Distributed | [08 — Messaging & Event-Driven Architecture](08-messaging.md) |
| 10 | Architecture | [09 — APIs, Service Architecture & Resilience](09-architecture-resilience.md) |
| 11–12 | Architecture | [10 — System Design Practice](10-system-design.md) |
| 13 | Operations | [11 — Containers, Kubernetes, CI/CD & Cloud](11-devops-cloud.md) |
| 13 | Operations | [12 — Observability, Reliability & Performance](12-observability-reliability.md) |
| 14 | Security | [15 — Security for Backend Engineers](15-security.md) |
| 15 | Senior signals | [13 — Behavioral, Leadership & Mock Loops](13-behavioral-mocks.md) |
| 1–15 | Ongoing | [14 — DSA Track](14-dsa-track.md): data structures from the inside, plus problems (every weekday) |
| 1–15 | Ongoing | [Capstone — Notification Platform v2](capstone.md): one real system built up across all modules (weekends) |

## Core books (buy or borrow these three first)

1. **Designing Data-Intensive Applications** — Martin Kleppmann. The most important
   book for senior backend interviews. Used in modules 03, 04, 07, 08.
2. **System Design Interview, Vol. 1 & 2** — Alex Xu. Interview-shaped practice. Module 10.
3. **Concurrency in Go** — Katherine Cox-Buday. Module 06.

Free and excellent:
- *Operating Systems: Three Easy Pieces* (ostep.org) — module 02
- *High Performance Browser Networking* (hpbn.co) — module 01
- *Use The Index, Luke* (use-the-index-luke.com) — module 03
- *Google SRE Book* (sre.google/books) — module 12
- *Learn Go with Tests* (quii.gitbook.io/learn-go-with-tests) — modules 06 and 14
- NeetCode roadmap (neetcode.io/roadmap) — module 14
- OWASP Cheat Sheet Series (cheatsheetseries.owasp.org) — module 15
- CMU Database Systems lectures and MIT 6.5840 Distributed Systems labs — for going deeper after the plan
