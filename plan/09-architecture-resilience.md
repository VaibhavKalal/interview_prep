# 09 — APIs, Service Architecture & Resilience (Week 10)

**Why it matters:** senior engineers are expected to shape architecture, not just
implement tickets. This module covers how services are cut, how they talk, and how they
survive each other's failures.

## Checklist

### L1 — Use
- [ ] REST design: resources, nouns, status codes, error format (RFC 9457 problem details)
- [ ] Pagination, filtering, sorting; API versioning strategies
- [ ] AuthN vs AuthZ; sessions vs JWT; OAuth2 flows (authorization code + PKCE, client credentials); OpenID Connect

### L2 — Internals
- [ ] gRPC: protobuf, HTTP/2 streams, deadlines; when gRPC beats REST
- [ ] GraphQL: resolvers, the N+1 problem and DataLoader, when it's a bad fit
- [ ] JWT internals: header.payload.signature, signing algorithms, why you can't revoke easily
- [ ] API gateway vs service mesh (sidecars, mTLS, retries at the mesh level)

### L3 — Production — architecture
- [ ] Monolith vs modular monolith vs microservices: honest tradeoffs, when to split
- [ ] Domain-Driven Design basics: bounded contexts, aggregates, ubiquitous language
- [ ] Data ownership: database per service; avoiding shared databases
- [ ] CQRS and event sourcing: what they are and when they're overkill
- [ ] Strangler fig pattern for migrating legacy systems (e.g. PHP monolith → services)
- [ ] Architecture Decision Records (ADRs)

### L3 — Production — resilience
- [ ] Timeouts everywhere; timeout budgets across a call chain
- [ ] Retries: only idempotent calls, exponential backoff + jitter, retry budgets, avoiding retry storms
- [ ] Circuit breakers (closed/open/half-open), bulkheads, load shedding
- [ ] Rate limiting algorithms: token bucket, leaky bucket, fixed vs sliding window
- [ ] Graceful degradation and fallbacks

## Lab (weekend, in `labs/09-resilience/`)

In Go:
1. Implement a **token bucket rate limiter** as HTTP middleware (per API key). Test it with concurrent requests.
2. Implement a **circuit breaker** wrapper around an HTTP client from scratch (closed → open after N failures → half-open after a cooldown).
3. Build two services: A calls B. Make B slow. Show A falling over without timeouts, then stable with timeout + retry with jitter + circuit breaker.
4. Write one ADR in `notes/09-architecture/` for a real decision from your work (e.g. "why SQS for notifications").

## Reading
- *Release It!* — Michael Nygard: stability patterns and anti-patterns chapters
- *Building Microservices* (2nd ed.) — Sam Newman: ch. 1–4
- AWS Builders' Library: "Timeouts, retries, and backoff with jitter" and "Avoiding insurmountable queue backlogs"
- Stripe's blog post on idempotent requests; Stripe API docs as a model of REST design

## Self-test
1. When should a team *not* split a monolith into microservices?
2. Service A calls B calls C, each with 3 retries. C goes down. What happens to traffic on C, and how do you prevent it?
3. JWT vs server-side sessions: how do you log a user out everywhere with JWTs?
4. Explain a circuit breaker's states and what metrics drive it.
5. Design the API for a multi-step checkout so a mobile client on a flaky network can't double-charge.

## Beyond the interview

- Read Martin Fowler's articles on Microservices, the Strangler Fig application, and "MonolithFirst" (martinfowler.com).
- Read *Learning Domain-Driven Design* — Vlad Khononov (more practical than the original DDD book).
- Write an ADR for a real decision your current team is facing, and share it with the team.
