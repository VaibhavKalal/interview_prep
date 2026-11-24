# 08 — Messaging & Event-Driven Architecture (Week 9)

**Why it matters:** you built a notification system on SQS. This module turns that
experience into deep, defensible expertise. Expect your interviewer to dig in here.

## Checklist

### L1 — Use
- [ ] Queue vs pub/sub vs log (stream)
- [ ] SQS, SNS, SNS→SQS fan-out, EventBridge; RabbitMQ; Kafka; Redis Streams: where each fits
- [ ] Synchronous (request/response) vs asynchronous communication tradeoffs

### L2 — Internals
- [ ] **SQS deep-dive:** visibility timeout, receive count, long polling, batch operations, message retention, max message size, delay queues
- [ ] SQS standard (at-least-once, best-effort ordering) vs FIFO (message groups, deduplication ID, throughput limits)
- [ ] Dead-letter queues, `maxReceiveCount`, redrive
- [ ] **Kafka:** topics, partitions, offsets, consumer groups, rebalancing, retention, compaction, replication (ISR, `acks=all`)
- [ ] Ordering: why ordering only holds per partition / per message group
- [ ] RabbitMQ: exchanges (direct, topic, fanout), acks, prefetch

### L3 — Production
- [ ] Idempotent consumers: dedupe tables, natural idempotency, idempotency keys
- [ ] Retries with exponential backoff and jitter; poison messages
- [ ] The dual-write problem and the **transactional outbox** pattern (+ change data capture with Debezium)
- [ ] Sagas: choreography vs orchestration; compensating actions
- [ ] Backpressure and consumer scaling: scaling on queue depth and age of oldest message
- [ ] Priority handling: separate queues for OTPs vs marketing
- [ ] Event schema evolution and versioning
- [ ] Monitoring: queue depth, `ApproximateAgeOfOldestMessage`, DLQ size alarms, consumer lag (Kafka)

## Lab (weekend, in `labs/08-messaging/`)

Use LocalStack (`docker run -p 4566:4566 localstack/localstack`) or a real AWS free tier account.
1. Rebuild a mini version of **your notification system** in Go: producer → SQS → worker pool consumer → fake "email provider" that randomly fails 20% of the time.
2. Add retries via visibility timeout, a DLQ after 3 failures, and an idempotency table so retries never double-send.
3. **Break it:** make processing slower than the visibility timeout. Watch messages get processed twice. Fix it (extend visibility, or tune the timeout).
4. Implement the **outbox pattern**: write an order and an outbox row in one Postgres transaction; a relay goroutine publishes outbox rows to SQS.
5. Then fill in `system-design/notification-system.md` Part B with what you learned.

## Reading
- AWS SQS Developer Guide: visibility timeout, DLQs, FIFO queues, best practices
- *DDIA*: ch. 11 (stream processing)
- *Kafka: The Definitive Guide* — ch. 1–4 (or Confluent's free docs on core concepts)
- microservices.io: "Transactional outbox", "Saga" pattern pages (Chris Richardson)

## Self-test
1. SQS is at-least-once. Walk through exactly how a duplicate happens and how your consumer handles it.
2. You write to the database, then publish an event, and the publish fails. Now what? How does the outbox fix it?
3. Kafka vs SQS for an order events stream consumed by 5 different services: which and why?
4. How do you keep order for events of the same user while still processing users in parallel?
5. A marketing blast of 10M notifications delays OTP messages. Redesign it.

## Beyond the interview

- Read "The Log: What every software engineer should know about real-time data's unifying abstraction" — Jay Kreps (LinkedIn engineering blog).
- Read the AWS Builders' Library article "Avoiding insurmountable queue backlogs".
- Write a blog post or internal doc about lessons from your SQS notification system. Writing it forces clarity, and it's a great interview artifact.
