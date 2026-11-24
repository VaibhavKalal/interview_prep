# 12 — Observability, Reliability & Performance (Week 13, second half)

**Why it matters:** "how would you debug this in production?" is one of the most common
senior questions. Good answers come from a structured method, not guesswork.

## Checklist

### L1 — Use
- [ ] The three pillars: logs, metrics, traces, and what each answers
- [ ] Structured logging (JSON), log levels, correlation/request IDs

### L2 — Internals
- [ ] Metric types: counter, gauge, histogram; why averages lie and percentiles (p50/p95/p99) matter
- [ ] Prometheus pull model, labels, and the cardinality problem
- [ ] Distributed tracing: trace and span IDs, context propagation (W3C `traceparent`), OpenTelemetry
- [ ] RED method (Rate, Errors, Duration) for services; USE method (Utilization, Saturation, Errors) for resources

### L3 — Production
- [ ] SLIs, SLOs, SLAs, error budgets; alerting on symptoms, not causes; burn-rate alerts
- [ ] Incident response: detect, mitigate first (rollback, scale, shed load), then find root cause
- [ ] Blameless postmortems
- [ ] Performance debugging method: find the bottleneck (CPU, memory, I/O, locks, downstream) before optimizing
- [ ] Load testing (k6, autocannon, vegeta): ramping, soak tests, finding the breaking point
- [ ] Capacity planning
- [ ] Chaos engineering basics

## Lab (weekend, in `labs/12-observability/`)

1. Add OpenTelemetry tracing and Prometheus metrics (RED metrics as histograms) to the module 09 services A → B.
2. Run Prometheus + Grafana + Jaeger with docker compose. Build a dashboard showing p50/p95/p99 latency and error rate.
3. **Break it:** inject latency into B. Find it using only the trace view. Then write a 1-page postmortem in `notes/12-observability/`.
4. Load test with k6 until something breaks. Write down what broke first and why.

## Reading
- *Google SRE Book* (free at sre.google/books): ch. 4 (SLOs), ch. 6 (monitoring), ch. 15 (postmortems)
- Brendan Gregg: "The USE Method" (brendangregg.com)
- OpenTelemetry docs: concepts

## Self-test
1. Why do you alert on p99 latency rather than average latency?
2. Define an SLO for your notification system. What would you alert on?
3. It's 2am: error rate jumped from 0.1% to 5% after no deploy. Walk through your first 15 minutes.
4. What's metric cardinality, and how can a single label take down Prometheus?
5. How does a trace ID get from a Node service, through SQS, into a Go consumer?
