# 03 — Relational Databases, In Depth (Weeks 3–4)

**Why it matters:** the database is usually the bottleneck and the source of the hardest
production bugs. Senior interviews go deep here. This module gets two weeks.

## Checklist

### L1 — Use
- [ ] Joins (inner, left, anti-join), `GROUP BY`/`HAVING`, subqueries, CTEs
- [ ] Window functions: `ROW_NUMBER`, `RANK`, `LAG`, running totals
- [ ] Normalization (1NF–3NF) and when to denormalize
- [ ] Transactions and ACID, in plain words

### L2 — Internals
- [ ] Storage: pages, heap files, rows on disk
- [ ] B+tree indexes: structure, why O(log n), why range scans are cheap
- [ ] Composite indexes and the leftmost-prefix rule; covering indexes / index-only scans
- [ ] Clustered (MySQL InnoDB primary key) vs heap + secondary indexes (Postgres)
- [ ] Write-ahead log (WAL / redo log): how durability and crash recovery work
- [ ] MVCC: how readers don't block writers; Postgres tuples + VACUUM vs InnoDB undo log
- [ ] Isolation levels and their anomalies: dirty read, non-repeatable read, phantom, lost update, write skew
- [ ] Locking: row locks, `SELECT ... FOR UPDATE`, `SKIP LOCKED`, gap locks (InnoDB), deadlock detection
- [ ] Query planner: how it uses statistics; seq scan vs index scan vs bitmap scan; nested loop vs hash vs merge join

### L3 — Production
- [ ] Reading `EXPLAIN ANALYZE` output and fixing a slow query
- [ ] N+1 queries in ORMs (Sequelize/Prisma, Eloquent, GORM) and how to fix them
- [ ] Connection pooling: why connections are expensive; PgBouncer; pool sizing
- [ ] Pagination: offset vs keyset (cursor), and why offset gets slow
- [ ] Zero-downtime migrations: adding columns, backfills, adding indexes concurrently, expand/contract
- [ ] Replication: sync vs async, replication lag, read-your-writes problems with read replicas
- [ ] Partitioning (by range/hash) and sharding: shard keys, hot shards, cross-shard queries
- [ ] Optimistic locking (version column) vs pessimistic locking (`FOR UPDATE`)

## Lab (both weekends, in `labs/03-databases/`)

Run Postgres in Docker (`docker run -e POSTGRES_PASSWORD=pw -p 5432:5432 postgres:17`).
1. Generate 5 million rows of orders with `generate_series`. Time a query by `customer_id` with no index. Read `EXPLAIN ANALYZE`. Add an index; read it again.
2. Build a composite index `(customer_id, created_at)`. Prove the leftmost-prefix rule with queries that do and don't use it.
3. Compare `OFFSET 1000000` pagination with keyset pagination. Measure both.
4. **Anomalies:** open two `psql` sessions and reproduce a *lost update* under `READ COMMITTED`. Then fix it three ways: `SELECT ... FOR UPDATE`, an atomic `UPDATE ... SET x = x - 1`, and `REPEATABLE READ` with a retry.
5. Create a deadlock between two sessions on purpose and read the Postgres error.
6. Build a job queue table in Go using `SELECT ... FOR UPDATE SKIP LOCKED` with several workers. (This is a real pattern used in production.)

## Reading
- *Designing Data-Intensive Applications*: ch. 2 (data models), ch. 3 (storage and retrieval), ch. 5 (replication), ch. 6 (partitioning), ch. 7 (transactions). **Ch. 7 is essential.**
- *Use The Index, Luke* (use-the-index-luke.com): the whole site
- PostgreSQL docs: "Concurrency Control" chapter and "Using EXPLAIN"
- Optional deeper: *Database Internals* — Alex Petrov, part I

## Self-test
1. Why does adding an index speed up reads but slow down writes? What else does it cost?
2. An index on `(a, b)` exists. Which of these use it: `WHERE a=1`, `WHERE b=1`, `WHERE a=1 AND b>5`, `ORDER BY a, b`?
3. Explain MVCC to a junior engineer. Why does Postgres need VACUUM?
4. Two users book the last seat at the same time. Walk through how you prevent a double booking, with at least two approaches and their tradeoffs.
5. What is write skew, and why doesn't `REPEATABLE READ` prevent it in Postgres?
6. After a user updates their profile, they reload and see old data. Why, and how do you fix it?
7. How do you add a `NOT NULL` column to a 500M-row table with zero downtime?
8. How would you choose a shard key for a multi-tenant SaaS app?

## Beyond the interview

- Read Postgres's own explanation of its MVCC and VACUUM internals in *The Internals of PostgreSQL* (interdb.jp, free).
- Watch the CMU Database Systems lectures by Andy Pavlo (free on YouTube): the storage, indexing, and concurrency control lectures.
- Find the 3 slowest queries in a real system you work on (`pg_stat_statements` or MySQL slow query log) and write up how you'd fix each.
