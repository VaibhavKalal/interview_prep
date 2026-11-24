# 07 — Distributed Systems Fundamentals (Week 8)

**Why it matters:** this is the theory behind every system design answer. It's what lets you
say *why* a design works instead of only *what* it is.

## Checklist

### L1 — Use
- [ ] Why distribute at all: scale, availability, latency
- [ ] Vertical vs horizontal scaling; stateless services
- [ ] Availability numbers: what 99.9% vs 99.99% means in downtime per month

### L2 — Internals
- [ ] The 8 fallacies of distributed computing
- [ ] Partial failures and why you can't tell a slow node from a dead one
- [ ] Clocks: why wall clocks lie; logical clocks (Lamport), vector clocks (concept)
- [ ] Replication: single-leader, multi-leader, leaderless (quorums: W + R > N)
- [ ] Consistency models: linearizability, sequential, causal, eventual, read-your-writes
- [ ] CAP theorem (and its limits) and PACELC
- [ ] Consensus: what problem it solves; Raft at the level of leader election + log replication (used by etcd, Consul, Kafka KRaft)
- [ ] Partitioning: by key range vs hash; rebalancing; consistent hashing

### L3 — Production
- [ ] Split brain and why you need a quorum
- [ ] Leader election in practice (etcd/ZooKeeper leases)
- [ ] Delivery semantics: at-most-once, at-least-once, "exactly-once" (= at-least-once + idempotency)
- [ ] Idempotency keys: design and storage
- [ ] Distributed transactions: why two-phase commit is avoided; sagas instead (see module 08)
- [ ] Unique IDs at scale: UUIDv4 vs UUIDv7 vs Snowflake-style IDs, and why sortability matters for indexes

## Lab (weekend, in `labs/07-distributed/`)

1. Play with the Raft visualization at raft.github.io. Kill the leader, partition the network, watch the election. Write a one-paragraph explanation in your notes.
2. In Go, implement a **consistent hash ring** with virtual nodes. Add and remove nodes and measure how many keys move (should be about 1/N).
3. Implement a **Snowflake-style ID generator** (timestamp + machine ID + sequence). Test it for uniqueness across goroutines. Think: what happens if the clock goes backwards?
4. Implement an idempotent HTTP endpoint using an `Idempotency-Key` header stored in Postgres. Retry the same request 10 times concurrently and prove only one side effect happens.

## Reading
- *DDIA*: ch. 5 (replication), ch. 8 (the trouble with distributed systems), ch. 9 (consistency and consensus). **Ch. 8 is essential.**
- "In Search of an Understandable Consensus Algorithm" (the Raft paper) — sections 1–5
- Jepsen.io analyses (skim one, e.g. on a database you've used)

## Self-test
1. Explain CAP with a concrete example. Why is "CA" not a real choice?
2. With N=3 replicas, what W and R give you strong reads? What's the cost?
3. Why can't you use timestamps from different servers to order events reliably?
4. What problem does Raft solve, and what happens during a network partition?
5. "We need exactly-once delivery." How do you respond?
