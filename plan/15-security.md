# 15 — Security for Backend Engineers (Week 14)

**Why it matters:** senior engineers are trusted to ship code that doesn't leak data. Security
questions come up inside system design ("how do you secure this API?"), in code review
rounds, and in behavioral rounds ("tell me about a security issue you handled").

## Checklist

### L1 — Use
- [ ] OWASP Top 10 (current edition) and API Security Top 10: know each item and one example
- [ ] SQL injection and parameterized queries (in Node, PHP PDO, and Go `database/sql`)
- [ ] XSS and CSRF: what they are and why APIs with cookie auth still need CSRF protection
- [ ] Never log secrets, tokens, or personal data; secrets come from a secrets manager, not from code

### L2 — Internals
- [ ] Hashing vs encryption vs encoding; why passwords use slow hashes (bcrypt, scrypt, Argon2) with salts
- [ ] Symmetric (AES-GCM) vs asymmetric (RSA, ECDSA/Ed25519) cryptography; where each is used in TLS
- [ ] HMAC: signing webhooks and request bodies; constant-time comparison
- [ ] JWT attacks: `alg: none`, algorithm confusion (RS256 vs HS256), missing expiry and audience checks
- [ ] OAuth2 and OIDC security details: PKCE, state parameter, redirect URI validation, token storage
- [ ] Authorization models: RBAC vs ABAC vs relationship-based (Zanzibar-style); broken object-level authorization (IDOR) as the #1 API risk
- [ ] CORS: what it actually protects (browsers only) and common misconfigurations

### L3 — Production
- [ ] SSRF: why letting users supply URLs (webhooks, image fetchers) is dangerous, and how to defend (allowlists, blocking internal IPs and cloud metadata endpoints)
- [ ] Multi-tenant isolation: tenant ID in every query; Postgres row-level security
- [ ] Rate limiting and brute-force protection on login and OTP endpoints
- [ ] Dependency and supply chain risk: `npm audit`, `govulncheck`, `composer audit`, lockfiles, pinned versions
- [ ] Least privilege in AWS IAM; short-lived credentials; secret rotation
- [ ] Encryption at rest and in transit; KMS and envelope encryption (concept)
- [ ] Audit logging and what to keep for incident response
- [ ] Threat modeling a feature with STRIDE (lightweight)

## Lab (weekend, in `labs/15-security/`)

In Go, using the services from earlier labs:
1. Write an endpoint vulnerable to **SQL injection** using string concatenation. Exploit it with `curl`. Fix it with a parameterized query.
2. Add an **IDOR** bug: `GET /orders/{id}` without checking ownership. Exploit it as a second user. Fix it with an ownership check, and add a test that proves the fix.
3. Implement **webhook signing** with HMAC-SHA256 and verify with `hmac.Equal`. Explain why `==` is unsafe.
4. Write a URL-fetching endpoint and show how it can reach `http://169.254.169.254/` (the cloud metadata IP). Add SSRF protection.
5. Run `govulncheck ./...` on the repo and read the output.
6. Write a one-page threat model (STRIDE) for your notification system in `notes/15-security/`.

## Reading
- OWASP Top 10 and OWASP API Security Top 10 (owasp.org)
- OWASP Cheat Sheet Series: Authentication, Password Storage, SSRF Prevention, JWT
- *Serious Cryptography* (2nd ed.) — Jean-Philippe Aumasson: ch. 1–2 and the hashing and authenticated encryption chapters
- *API Security in Action* — Neil Madden (optional, deep)

## Self-test
1. How do you store passwords, and why not SHA-256?
2. A user can see another user's invoice by changing the ID in the URL. What's the bug class, and how do you prevent it systematically (not just in one handler)?
3. Your service accepts webhook URLs from customers. What's the risk and how do you mitigate it?
4. How do you verify that a webhook really came from Stripe (or from your own system)?
5. Where do your Node/Go services get database passwords from in production, and how are they rotated?
6. Walk through securing the notification system: who can send, to whom, and how you stop abuse.

## Beyond the interview
- Read one real postmortem of a security incident (e.g. the Capital One SSRF breach write-ups) and note which controls would have stopped it.
- Do a few labs from PortSwigger Web Security Academy (free): SQL injection, SSRF, access control.
