# Capstone — Notification Platform v2 (built across all modules)

The labs teach one idea each. The capstone joins them into **one real, production-style
system**, so you learn how the pieces interact, which is where most real engineering
knowledge lives. It is also something concrete to talk about in interviews:
*"I rebuilt our notification system in Go to go deeper; here's what I learned."*

It lives in `capstone/` in this repo. Add to it a little each week: about 1–2 hours on the weekend,
after the module lab. Many labs can be built directly inside the capstone instead of separately.

## What it is

A multi-tenant notification service: clients send a request ("send an OTP to this user by
SMS"), and the platform delivers it reliably through email, SMS, and push providers (faked
locally), with retries, priorities, rate limits, and full observability.

```
client → API (Go) → Postgres (+ outbox) → queue (SQS on LocalStack) → workers (Go) → fake providers
                         ↑                                                    ↓
                       Redis (rate limits, idempotency, cache)        status updates + metrics + traces
```

## Milestones (tick as you go)

| After module | Add to the capstone |
|---|---|
| 01 Networking | [ ] Go HTTP API skeleton: `POST /notifications`, `GET /notifications/{id}`, proper timeouts, JSON errors |
| 02 OS & concurrency | [ ] Worker pool that processes notifications concurrently; race-free (`go test -race`) |
| 03 Databases | [ ] Postgres schema (tenants, users, notifications, delivery attempts); migrations; keyset pagination for listing; indexes justified with `EXPLAIN ANALYZE` |
| 04 Caching | [ ] Redis: cache user preferences; idempotency keys for `POST /notifications` |
| 05 Node.js | [ ] A small Node.js client SDK or admin service that calls the API (shows both stacks working together) |
| 06 Go | [ ] Context cancellation end to end; graceful shutdown; `pprof` endpoint |
| 07 Distributed | [ ] Snowflake-style or UUIDv7 IDs; idempotent processing so duplicates never double-send |
| 08 Messaging | [ ] SQS (LocalStack) between API and workers; outbox pattern; DLQ; separate high-priority queue for OTPs |
| 09 Resilience | [ ] Per-tenant token bucket rate limiting; circuit breaker + retries with jitter around providers |
| 10 System design | [ ] Write `capstone/DESIGN.md`: requirements, estimates, design, tradeoffs, what changes at 100x |
| 11 DevOps | [ ] Dockerfiles (multi-stage), docker compose for local dev, Kubernetes manifests with probes, GitHub Actions CI |
| 12 Observability | [ ] OpenTelemetry traces across API → queue → worker; RED metrics; Grafana dashboard; an SLO |
| 15 Security | [ ] API keys per tenant, tenant isolation in every query, HMAC-signed delivery webhooks, `govulncheck` in CI |
| 13 Behavioral | [ ] Write the 5-minute walk-through you'd give in an interview |

## Rules
- Write tests as you go (unit tests + a few integration tests against Postgres in Docker).
- Every non-obvious decision gets a short ADR in `capstone/adr/` (what you chose, the alternatives, why).
- Break it on purpose at each milestone and write down what happened in `capstone/LEARNINGS.md`.
