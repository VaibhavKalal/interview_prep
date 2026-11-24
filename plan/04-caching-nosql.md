# 04 — Caching, Redis & NoSQL (Week 5)

**Why it matters:** caching is the first tool for scale and the source of the subtlest
bugs (stale data, stampedes). NoSQL choices come up in every system design round.

## Checklist

### L1 — Use
- [ ] Redis data types: string, hash, list, set, sorted set, stream; TTLs
- [ ] Cache-aside (lazy loading) pattern
- [ ] When to use a document store, key-value store, wide-column store

### L2 — Internals
- [ ] Why Redis is single-threaded and still fast; its event loop
- [ ] Redis persistence: RDB snapshots vs AOF; durability tradeoffs
- [ ] Eviction policies: LRU, LFU, `allkeys` vs `volatile`
- [ ] Redis replication, Sentinel, and Cluster (hash slots)
- [ ] LSM trees + SSTables + compaction (Cassandra, RocksDB, DynamoDB's style) vs B-trees: write vs read tradeoffs
- [ ] DynamoDB: partition key and sort key, GSIs, single-table design, capacity modes
- [ ] Consistent hashing and virtual nodes

### L3 — Production
- [ ] Write-through, write-behind, refresh-ahead: when to use each
- [ ] Cache invalidation strategies; why "delete on write" beats "update on write"
- [ ] Cache stampede / thundering herd: request coalescing (singleflight), locks, probabilistic early expiry, jittered TTLs
- [ ] Hot keys and big keys in Redis; how to detect and fix them
- [ ] Distributed locks with Redis: `SET NX PX`, why fencing tokens are needed, the Redlock debate
- [ ] Rate limiting with Redis (sorted sets or Lua scripts)
- [ ] Choosing SQL vs NoSQL: access patterns first, then consistency and scale needs

## Lab (weekend, in `labs/04-caching/`)

1. In Go, build an **LRU cache** from scratch with a map + doubly linked list (this is also a top interview coding question). Make it safe for concurrent use.
2. Put Redis in front of the Postgres orders table from module 03 using cache-aside. Measure latency with and without the cache.
3. **Break it:** expire a hot key while 200 goroutines request it. Count how many hit the database. Fix it with `golang.org/x/sync/singleflight` and count again.
4. Implement a sliding-window rate limiter with a Redis sorted set.

## Reading
- *DDIA*: ch. 3 (LSM vs B-tree section), ch. 6 (partitioning)
- Redis docs: "Redis persistence", "Key eviction", "Distributed locks"
- Martin Kleppmann, "How to do distributed locking" (blog post critiquing Redlock)
- AWS DynamoDB docs: "Best practices for designing and architecting with DynamoDB"

## Self-test
1. Your cache hit rate is 95%, but the database still falls over every hour on the hour. What's happening?
2. Why is deleting a cache key on write safer than updating it?
3. Why does a Redis-based lock need a fencing token to be safe?
4. LSM tree vs B-tree: which suits a write-heavy event log, and why?
5. Design the DynamoDB keys for "get all orders of a customer, newest first".
