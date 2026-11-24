# 02 — OS, Concurrency & I/O (Week 2)

**Why it matters:** the Node.js event loop, PHP-FPM workers, and Go goroutines are three
different answers to one question: *how do we handle many requests at once?* You can't
compare them deeply without this module.

## Checklist

### L1 — Use
- [ ] Process vs thread; what each shares and doesn't share
- [ ] CPU-bound vs I/O-bound work
- [ ] Concurrency vs parallelism

### L2 — Internals
- [ ] Virtual memory, pages, stack vs heap
- [ ] Context switches: what they cost, process vs thread vs goroutine switches
- [ ] System calls and user mode vs kernel mode
- [ ] File descriptors; "everything is a file"; `ulimit -n`
- [ ] I/O models: blocking, non-blocking, I/O multiplexing (`select`/`poll`/`epoll`/`kqueue`), async I/O (`io_uring`)
- [ ] Why `epoll` lets one thread handle 10,000 connections (the C10K problem)
- [ ] Race conditions, critical sections, mutexes, semaphores, condition variables
- [ ] Deadlock: the 4 conditions, and how to prevent it (lock ordering)
- [ ] Atomic operations and compare-and-swap

### L3 — Production
- [ ] Thread-per-request (classic Java, PHP-FPM process-per-request) vs event loop (Node, Nginx) vs M:N green threads (Go): tradeoffs in memory, latency, and failure modes
- [ ] Out-of-memory kills (OOM killer), memory limits in containers
- [ ] Reading `top`/`htop`: load average, CPU steal, iowait
- [ ] Tools: `strace`, `lsof`, `vmstat`, `pidstat`

## Lab (weekend, in `labs/02-concurrency/`)

In Go:
1. Write a counter incremented by 1,000 goroutines **without** a lock. Run with `go run -race`. See the race.
2. Fix it three ways: `sync.Mutex`, `sync/atomic`, and a single goroutine owning the counter via a channel. Benchmark all three with `go test -bench`.
3. Create a deadlock on purpose with two mutexes locked in opposite order. Read Go's deadlock/hang output, then fix it with lock ordering.
4. Run `strace -f -e trace=network` on your HTTP server from module 01 and find the `accept`, `read`, `write`, and `epoll_wait` calls.

## Reading
- *Operating Systems: Three Easy Pieces* (ostep.org): Processes, Address Spaces, Concurrency intro, Locks, Condition Variables, Common Concurrency Problems
- "The C10K problem" — Dan Kegel (classic web article)

## Self-test
1. Why can Node.js handle 10,000 idle WebSocket connections on one thread, but PHP-FPM can't?
2. What actually happens when a Node.js process calls `fs.readFile`? Which thread does the work?
3. What's the difference between a mutex and a semaphore? Give a backend use for each.
4. Your container is killed with exit code 137. What happened and how do you investigate?
5. Explain a deadlock you could create in a database, not just in code.

## Beyond the interview

- Read a short `epoll` echo server in C (many exist online) to see what Node's libuv and Go's netpoller wrap.
- Read Julia Evans' zines on Linux internals and `strace` (jvns.ca); use `strace` on a real service at work (in a safe environment).
- Learn `perf` basics and generate a flame graph of a CPU-heavy program (Brendan Gregg's flame graph page).
