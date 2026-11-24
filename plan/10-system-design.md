# 10 — System Design Practice (Weeks 11–12)

**Why it matters:** this round usually decides the senior level. Modules 01–09 gave you the
building blocks; now you practice assembling them under time pressure, out loud.

## The 45-minute framework

| Minutes | Step | What to produce |
|---|---|---|
| 0–5 | Requirements | Functional + non-functional; agree on scope. Ask questions. |
| 5–10 | Estimates | QPS (average and peak), storage, bandwidth; read/write ratio |
| 10–15 | API + data model | Endpoints; tables/keys; access patterns |
| 15–25 | High-level design | Boxes and arrows; walk one request through |
| 25–40 | Deep dives | 2–3 hard parts: scaling the bottleneck, consistency, failure handling |
| 40–45 | Wrap-up | Bottlenecks, monitoring, what you'd do with more time |

**Numbers to memorize:** latency numbers every programmer should know (L1 cache ~1ns,
RAM ~100ns, SSD read ~100µs, same-datacenter round trip ~0.5ms, cross-continent ~100ms);
1 day ≈ 86,400 s ≈ 10⁵ s; 1M requests/day ≈ 12 QPS.

## Problems (do each with the [template](../system-design/_template.md), one file per problem in `system-design/`)

- [ ] **Your notification system** (redo it as a clean design at 100x scale)
- [ ] URL shortener (IDs, caching, redirects, analytics)
- [ ] Rate limiter as a distributed service
- [ ] Chat / messaging (WebSockets, presence, message ordering, offline delivery)
- [ ] News feed (fan-out on write vs read, celebrity problem)
- [ ] Payment system (idempotency, ledgers, double-entry, reconciliation)
- [ ] Distributed job scheduler (cron at scale, exactly-once-ish execution)
- [ ] Search autocomplete (tries, top-k, caching)
- [ ] File storage / Dropbox (chunking, dedupe, sync)
- [ ] Ride matching / proximity (geohash, quadtrees)

## How to practice each problem
1. Set a 45-minute timer. Talk out loud. Draw on paper or excalidraw.com.
2. Write the design up in the template.
3. Compare with Alex Xu's chapter (or other references). List what you missed in `mistakes.md`.
4. Ask Claude to run it as a mock: Claude plays interviewer and pushes on your weak spots.

## Reading
- *System Design Interview* Vol. 1 & 2 — Alex Xu
- *DDIA* (re-read chapters relevant to each problem)
- Engineering blogs for real architectures: Discord (message storage), Uber, Stripe, Slack, Netflix

## Self-test
- Can you do any problem from the list in 45 minutes without notes, and handle "what if traffic grows 100x?" and "what if this component dies?"
