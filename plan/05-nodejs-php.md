# 05 — Node.js Internals (+ PHP-FPM) (Week 6)

**Why it matters:** this is your home turf. Interviewers will expect real depth in the
languages on your CV. "I've used it for 4 years" must come with "and here's how it works inside."

## Checklist

### L1 — Use
- [ ] Callbacks → Promises → async/await; error handling in each
- [ ] Modules: CommonJS vs ES modules
- [ ] Streams API basics: readable, writable, transform, `pipeline`

### L2 — Internals
- [ ] Architecture: V8 + libuv + Node bindings
- [ ] Event loop phases in order: timers → pending callbacks → poll → check → close callbacks
- [ ] Microtasks (`process.nextTick`, Promises) vs macrotasks; why `nextTick` runs before Promise callbacks
- [ ] libuv thread pool (default size 4): which operations use it (fs, dns.lookup, crypto, zlib) and which don't (network sockets)
- [ ] V8 memory: heap spaces, generational garbage collection, `--max-old-space-size`
- [ ] Streams and backpressure: `highWaterMark`, what `write()` returning `false` means
- [ ] `cluster` vs `worker_threads` vs child processes

### L3 — Production
- [ ] Detecting event loop blocking: `perf_hooks.monitorEventLoopDelay`, symptoms in latency graphs
- [ ] Memory leaks: common causes (global caches, closures, listeners); heap snapshots in Chrome DevTools
- [ ] CPU profiling with `--cpu-prof` and flame graphs (clinic.js / 0x)
- [ ] Graceful shutdown: handling `SIGTERM`, draining connections, closing DB pools (critical in Kubernetes)
- [ ] Unhandled promise rejections and process crash behavior
- [ ] Outbound HTTP connection reuse (`keepAlive` agent; undici in modern Node)

### PHP-FPM (compare, don't go as deep)
- [ ] PHP-FPM process model: master + worker pool; `pm = static | dynamic | ondemand`; `pm.max_children` sizing
- [ ] Shared-nothing: every request starts fresh. Pros (isolation, no leaks across requests) and cons (bootstrap cost every request)
- [ ] OPcache and preloading; JIT in PHP 8
- [ ] Long-running PHP: Laravel Octane / Swoole / RoadRunner and the new bugs they introduce (state leaking between requests)

## Lab (weekend, in `labs/05-node/`)

1. Write a Node script that logs from `setTimeout`, `setImmediate`, `process.nextTick`, and `Promise.resolve().then`, inside and outside an I/O callback. **Predict the order first**, then run it.
2. Start an HTTP server with an endpoint that runs a 2-second CPU loop. Load test another endpoint with `autocannon` while calling the slow one. Watch latency explode. Fix it with `worker_threads`.
3. Run 8 parallel `crypto.pbkdf2` calls and time them. Then set `UV_THREADPOOL_SIZE=8` and time again. Explain the result.
4. Create a memory leak on purpose (an ever-growing array in a request handler). Find it with a heap snapshot.
5. Pipe a 1 GB file through a transform stream with and without `pipeline`; watch memory with and without backpressure handling.

## Reading
- Node.js docs: "The Node.js Event Loop", "Don't Block the Event Loop (or the Worker Pool)" — both essential
- *Node.js Design Patterns* — Mario Casciaro & Luciano Mammino: ch. 1–6
- PHP manual: FPM configuration

## Self-test
1. Walk through the event loop phases. Where do `setImmediate` and `setTimeout(fn, 0)` run, and why is their order sometimes random?
2. Why can `dns.lookup` slow down unrelated file reads?
3. Your Node API's p99 latency jumped from 120ms to 2s after a deploy, but CPU looks normal. How do you investigate?
4. How does your Node service shut down gracefully during a Kubernetes rolling deploy?
5. Compare how Node, PHP-FPM, and Go each handle 1,000 concurrent slow requests (each waiting 1 second on a database).

## Beyond the interview

- Read libuv's design overview (docs.libuv.org, "Design overview") and match each event loop phase to the code.
- Profile a real service at work with `--cpu-prof` and find one hot path worth optimizing.
- Watch "What the heck is the event loop anyway?" (Philip Roberts, JSConf) and Bert Belder's talk on the Node.js event loop for the libuv side.
