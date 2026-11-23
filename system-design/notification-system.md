# Notification System (my real project — SQS)

This is your strongest interview asset: a real system you built.
Fill it in from memory first, then we'll stress-test it like an interviewer would.

## Part A — What I actually built

- **What it did** (channels: email / SMS / push / in-app?):
- **Scale** (notifications per day, peak per second):
- **Architecture** (producers → SQS → consumers → providers?):
- **Tech** (language, DB, providers like SES / SNS / FCM / Twilio):
- **My role** (designed it? built part? owned it in production?):
- **Hardest problem I solved:**
- **What I'd do differently now:**

## Part B — Questions an interviewer will ask (prepare answers)

1. SQS standard vs FIFO: which did you use and why?
2. SQS is at-least-once. How did you avoid sending the same notification twice?
3. What happened when a provider (e.g. email service) was down or slow?
4. How did retries work? Backoff? Dead-letter queue? What did you do with DLQ messages?
5. Visibility timeout: what value, and what happens if processing takes longer?
6. How did you scale consumers? What was the bottleneck?
7. How did you handle user preferences, quiet hours, unsubscribes?
8. Priorities: how would an OTP avoid waiting behind a marketing blast?
9. Rate limits from providers: how did you respect them?
10. How did you monitor it? What alerts existed? Queue depth? Age of oldest message?
11. How would you redesign it for 100x scale?

## Part C — Clean design (fill in using the template)

See [_template.md](_template.md).
