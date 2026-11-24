# 01 — Networking & HTTP (Week 1)

**Why it matters:** every backend bug eventually touches the network: timeouts,
connection pool exhaustion, TLS errors, load balancer health checks, keep-alive mismatches.

## Checklist

### L1 — Use
- [ ] Foundations: packets, layers, IP, ports, sockets, RTT, DNS, TLS, blocking I/O ([Day 0 notes](../notes/01-networking/day0-foundations.md))
- [x] The request lifecycle: DNS → TCP → TLS → HTTP → LB → app → DB ([Day 1 notes](../notes/01-networking/day1-request-lifecycle.md))
- [ ] HTTP methods, status codes (know 200/201/204/301/302/304/400/401/403/404/409/422/429/500/502/503/504)
- [ ] Headers: `Cache-Control`, `ETag`, `Authorization`, `Content-Type`, `Connection`, `X-Forwarded-For`
- [ ] Which HTTP methods are safe and which are idempotent, and why it matters for retries

### L2 — Internals
- [ ] OSI vs TCP/IP model; what L4 vs L7 really means
- [ ] TCP: 3-way handshake, sequence numbers, retransmission, flow control (window), congestion control (slow start)
- [ ] TCP connection states, especially `TIME_WAIT` and `CLOSE_WAIT`, and what causes a pile-up of each
- [ ] UDP vs TCP; why QUIC/HTTP3 is built on UDP
- [ ] TLS 1.3 handshake, certificates and chain of trust, SNI, mTLS
- [ ] HTTP/1.1 vs HTTP/2 (multiplexing, HPACK, streams) vs HTTP/3 (QUIC, no TCP head-of-line blocking)
- [ ] DNS: record types (A, AAAA, CNAME, MX, TXT, NS), TTL, recursive vs authoritative
- [ ] WebSockets, Server-Sent Events, long polling: how each works on the wire

### L3 — Production
- [ ] Timeouts: connect vs read vs idle; why every outbound call needs one
- [ ] Keep-alive mismatch: why a load balancer idle timeout shorter than the app's causes random 502s
- [ ] Connection pooling for outbound HTTP (Node `http.Agent`, Go `http.Transport`)
- [ ] Ephemeral port exhaustion: symptoms and fixes
- [ ] Load balancing algorithms: round robin, least connections, consistent hashing; health checks
- [ ] CDNs, reverse proxies (Nginx), SSL termination
- [ ] Debugging tools: `curl -v`, `dig`, `ss`/`netstat`, `tcpdump`, `openssl s_client`

## Lab (weekend, in `labs/01-http/`)

Build in Go with only the standard library:
1. A raw TCP echo server using `net.Listen`; connect to it with `nc localhost 8080`.
2. A minimal HTTP/1.1 server **on top of raw TCP**: parse the request line and headers yourself,
   return a valid response. Hit it with `curl -v`. (This makes HTTP stop being magic.)
3. **Break it:** make the handler sleep 10s, then call it with `curl --max-time 2`. Then write a Go
   client *without* a timeout and one *with* `http.Client{Timeout: ...}`. Observe the difference.
4. Run `ss -tan` during and after the requests and find connections in `TIME_WAIT`.

## Reading
- *High Performance Browser Networking* (hpbn.co): ch. 1–4 (latency, TCP, UDP, TLS), ch. 9–12 (HTTP/1.1, HTTP/2)
- MDN HTTP docs: caching, status codes, CORS

## Self-test (explain without notes)
1. Why is the first request to a new server slower than the following ones?
2. A service shows thousands of sockets in `CLOSE_WAIT`. Whose bug is it, and what's the likely cause?
3. Users get occasional 502s from the ALB, but app logs show no errors. What do you check?
4. What does HTTP/2 fix compared to HTTP/1.1, and what problem does it still have that HTTP/3 fixes?
5. A client retries a `POST /payments` after a timeout. What can go wrong and how do you prevent it?
6. L4 vs L7 load balancer: when do you need L7?

## Beyond the interview

- Read Cloudflare's blog posts on HTTP/3 and QUIC, and on how they handle DDoS at the network layer.
- Capture your own HTTPS request with Wireshark (or `tcpdump` + Wireshark) and identify the TCP handshake, TLS ClientHello/ServerHello, and HTTP/2 frames.
- Read the source of Go's `net/http` server loop (`server.go`, `func (srv *Server) Serve`) to see `accept` → goroutine per connection.
