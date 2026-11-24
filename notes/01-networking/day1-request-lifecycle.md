# Day 1 — What happens when you hit `https://api.example.com/users`?

A classic opener. It can go as deep as the interviewer wants, so know each step.

> **New to the basics?** Read [Day 0: Networking foundations](day0-foundations.md) first: packets, layers, ports, sockets, round trips, TLS, and blocking I/O.

## 1. DNS resolution
- Browser cache → OS cache → recursive resolver (ISP or 8.8.8.8)
- Resolver asks: root server → `.com` TLD server → authoritative server for `example.com`
- Gets an IP (A record for IPv4, AAAA for IPv6). Result is cached for its **TTL**.
- Interview angle: low TTLs help failover but increase DNS load.

## 2. TCP handshake
- SYN → SYN-ACK → ACK: **1 round trip** before any data is sent.
- This is why connection reuse (keep-alive, connection pools) matters so much.

## 3. TLS handshake
- Server sends its certificate; client verifies it against trusted CAs.
- Both sides agree on a symmetric session key (via key exchange such as ECDHE).
- TLS 1.3: 1 round trip. TLS 1.2: 2 round trips. TLS 1.3 resumption can do 0-RTT.

## 4. HTTP request
- Method, path, headers, body.
- HTTP/1.1: keep-alive reuses the connection, but one request at a time per connection.
- HTTP/2: many requests multiplexed on one connection; header compression.
- HTTP/3: runs over QUIC (UDP); avoids TCP head-of-line blocking.

## 5. Load balancer
- **L4** (transport): routes by IP and port. Fast, doesn't read the request.
- **L7** (application): reads HTTP, so it can route by path, header, or cookie, and do SSL termination, retries, and rate limiting.
- Need L7 when: path-based routing (`/api` vs `/static`), sticky sessions by cookie, canary by header.

## 6. App server (depends on the runtime)
- **Node.js:** single-threaded event loop picks up the request. I/O (DB, HTTP calls) is non-blocking. CPU-heavy work blocks *everyone*.
- **PHP-FPM:** a worker process from the pool handles the whole request, then resets. A slow request blocks only that worker, but the pool size caps concurrency.
- **Go:** `net/http` starts a goroutine per request. The Go scheduler spreads goroutines across OS threads and CPU cores.

## 7. Database / cache
- App takes a connection from the **pool** (opening one is expensive: TCP + auth).
- Query uses an index (or doesn't, and does a full scan).
- Cache (Redis) may answer first.

## 8. Response
- Status code, headers (`Cache-Control`, `ETag`, `Content-Type`), body.
- Travels back along the same connection.

---

## Check yourself (answer in your own words, then review with Claude)

1. Why is the **first** request to a new server slower than the following ones? Name at least two reasons.
2. What's the difference between an **L4 and an L7 load balancer**, and when would you need L7?
3. In Node.js, if one request runs a heavy synchronous loop for 2 seconds, what happens to the other requests? What would happen in Go or PHP-FPM?

## My answers

1. DNS resolution isn't cached for the first request so it goes through the entire flow to fetch the ip. TCP and TLS handshake add to the time for the initial request.
2. L4 is a network layer load balancer which routes based on ip and port, L7 is application layer load balancer which can parse http payload that enable advanced routing based on path. L7 is useful routing is to be performed baised on path, as it can parse the requests we can introduce security features like firewall or ratelimiting.
3. As Node.js is single threaded if execution take 2 seconds the thread will be block for that time, rendering the application un-responsive. In case of PHP-FPM as the request is served by the pool of workers even if the execution is block then consicutive requests are served by the other workers from the pool. In case of go requests are served by go routine each are executed by the os threads if any of the thread is blocked the request is served by the other available thread.

## My 3-line summary

- Any requests go through following lifecycle i.e. DNS resolution -> TCP handshake -> TLS handshake -> HTTP Request -> Load Balancer -> Application Server -> Database -> response.
- There are two types of load balancers i.e. L4 (network layer) and L7 (application layer)
- As Node.js is single threaded if execution take 2 seconds the thread will be block for that time, rendering the application un-responsive. In case of PHP-FPM as the request is served by the pool of workers even if the execution is block then consicutive requests are served by the other workers from the pool. In case of go requests are served by go routine each are executed by the os threads if any of the thread is blocked the request is served by the other available thread.
