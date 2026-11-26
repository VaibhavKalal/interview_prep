# Day 0: Networking foundations (read before Day 1)

The Day 1 lifecycle (DNS → TCP → TLS → HTTP → LB → app) assumes you already know what a packet,
an IP address, a port, a socket, and a round trip are. This file explains those basics, so each
Day 1 step has something to rest on.

Each section ends with a **Try it** command. Run it, because seeing real output makes the ideas stick.

---

## 1. The core problem: computers can only send small chunks

A network can't send "a file" or "an HTTP request". It sends **packets**: chunks of bytes,
usually at most ~1500 bytes each (the **MTU**, maximum transmission unit, of Ethernet).

Everything else is software built on top of that, and it has to deal with these facts:
- Packets can be **lost** (a congested router drops them).
- Packets can arrive **out of order** (they may take different paths).
- Packets can be **duplicated**.
- The sender can be faster than the receiver or the network can handle.

TCP exists to hide those facts. Keep this list in mind; it explains almost every TCP feature.

---

## 2. Layers: each one solves one problem

Networking is built in layers. Each layer uses the one below it and doesn't care how that one works.

| Layer | Name | Solves | Address used | Examples |
|---|---|---|---|---|
| L2 | Link | Get a frame to the **next device** on the same local network | MAC address (`aa:bb:cc:…`) | Ethernet, Wi-Fi |
| L3 | Network | Get a packet across many networks to the **right machine** | IP address (`93.184.216.34`) | IP, ICMP (`ping`) |
| L4 | Transport | Get data to the **right program** on that machine, optionally reliably | Port (`443`) | TCP, UDP |
| L7 | Application | What the bytes **mean** | URL, headers, … | HTTP, DNS, SMTP, gRPC |

(The OSI model also has L5 session and L6 presentation. In practice they're folded into L7, so
people say "L4" and "L7" and skip the rest. Nobody load-balances at L5.)

**Encapsulation:** each layer wraps the data from the layer above with its own header:

```
[ Ethernet hdr | IP hdr | TCP hdr | HTTP request bytes ]
     L2           L3       L4          L7 (payload)
```

A router only reads up to the **IP header** (L3) to decide where to forward the packet.
It doesn't know or care that HTTP is inside.

**This is what "L4 vs L7 load balancer" means:**
- An **L4 LB** reads only IP + TCP headers (source/destination IP and port). It forwards TCP
  connections without understanding their contents. It can't see the URL, because to an L4 LB
  the HTTP bytes are an opaque payload, and with HTTPS they're encrypted anyway.
- An **L7 LB** reads the full HTTP request. To do that it must **end the client's TCP
  connection** (and decrypt TLS), read the request, then open its **own** connection to a backend.
  So there are two separate TCP connections: client↔LB and LB↔app.

> Common mistake: L4 is the **transport** layer, not the network layer. The network layer is L3 (IP).

**Try it:** `traceroute google.com` (or `tracepath google.com`) shows each L3 router your packets pass through.

---

## 3. IP addresses and routing (L3)

- An **IP address** identifies a network interface. IPv4 is 32 bits (`10.0.1.5`); IPv6 is 128 bits.
- **Private ranges** (`10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) aren't reachable from the
  internet. Your AWS VPC uses these.
- **NAT** (network address translation) lets many private machines share one public IP. An AWS
  NAT Gateway does this for private subnets that need outbound internet access.
- **CIDR** notation `10.0.0.0/16` means "the first 16 bits are fixed", so 65,536 addresses.
- IP is **best effort**: no delivery guarantee, no ordering. That's TCP's job.

**Try it:** `ip addr` (your own IPs), `ip route` (where your machine sends packets by default).

---

## 4. Ports and sockets (L4): how data reaches the right program

One machine runs many programs, so the IP address alone isn't enough. A **port** (0–65535)
picks the program: 443 for HTTPS, 5432 for Postgres, 6379 for Redis.

- A **server** asks the OS to **listen** on a fixed port: in Go `net.Listen("tcp", ":8080")`,
  in Node `server.listen(8080)`.
- A **client** gets a random **ephemeral port** from the OS (Linux: 32768–60999) for each
  outgoing connection.

A **TCP connection** is identified by a **4-tuple**:

```
(client IP, client port, server IP, server port)
(10.0.1.5,  51234,       10.0.2.9,  5432)
```

So one server port (5432) can hold thousands of connections, because each one has a different
client IP/port pair. It also means **one client talking to one server:port can have at most
~28k connections**, one per free ephemeral port. That's the root of "ephemeral port
exhaustion" (Module 01 L3).

A **socket** is the OS object your program uses for one end of a connection. To your code it's a
file descriptor (an integer) that you `read()` and `write()`. Every open connection uses one
file descriptor, which is why `ulimit -n` (max open files) matters for servers.

**Try it:** `ss -tlnp` lists listening sockets; `ss -tn` lists established connections with their 4-tuples.

---

## 5. TCP vs UDP

**UDP** = "here's a packet, good luck". No connection, no handshake, no retransmission, no ordering.
Cheap and fast. Used by DNS, video calls, games, and QUIC (which builds its own reliability on top).

**TCP** = a **reliable, ordered byte stream**. It solves each problem from section 1:

| Problem | TCP's answer |
|---|---|
| Lost packets | Receiver sends **ACKs**; sender **retransmits** anything not ACKed in time |
| Out of order / duplicates | Every byte has a **sequence number**; receiver reorders and drops duplicates |
| Receiver too slow | **Flow control**: receiver advertises a window ("I have room for N bytes") |
| Network too slow | **Congestion control**: sender starts slowly and speeds up (**slow start**) |

**Why a handshake?** Both sides must agree on starting sequence numbers before sending data:

```
Client                        Server
  | ---- SYN (my seq = x) ---->  |
  | <--- SYN-ACK (seq = y,      |
  |       ack = x+1) ---------  |
  | ---- ACK (ack = y+1) ----->  |   ← client can send data with this packet
```

That costs **1 round trip** before the first byte of your request.

**Slow start** matters for "why is the first request slow": a new connection may only send about
**10 packets (~14 KB)** before waiting for ACKs, and the limit roughly doubles each round trip.
A warm, reused connection has already grown its window. A new one starts small again.

A connection ends with FIN/ACK exchanges. Whoever closes first sits in **`TIME_WAIT`** for a
while (60s on Linux). If a side gets the peer's FIN but its app never calls `close()`, the socket
sits in **`CLOSE_WAIT`** forever. That one is always a bug in the app that didn't close.

**Try it:** `sudo tcpdump -i any -n port 80` in one terminal, `curl http://example.com` in another.
Find the `[S]`, `[S.]`, `[.]` flags: that's SYN, SYN-ACK, ACK.

---

## 6. Latency, round trips, and bandwidth

- **Latency:** time for one packet to get there. Limited by the **speed of light** in fiber,
  about 200,000 km/s, so 5 ms per 1,000 km one way.
- **RTT (round-trip time):** there and back. Mumbai → us-east-1 (Virginia) is ~13,000 km, so the
  **RTT is ~180–200 ms**. Same AWS region, between AZs: ~1 ms. Same AZ: <0.5 ms.
- **Bandwidth:** how many bytes per second fit through the pipe.

For typical API calls (small requests and responses), **round trips dominate, not bandwidth.**
Doubling bandwidth doesn't help a 2 KB JSON response; removing a round trip does.

So count round trips for a cold HTTPS request:

| Step | Round trips | At 200 ms RTT |
|---|---|---|
| DNS lookup (cache miss) | 1 to the resolver (and the resolver may need more upstream) | 200+ ms |
| TCP handshake | 1 | 200 ms |
| TLS 1.3 handshake (TLS 1.2: 2) | 1 | 200 ms |
| HTTP request → response | 1 | 200 ms |
| **Total** | **~4** | **~800 ms** |

A reused (keep-alive) connection with a cached DNS entry costs **1 round trip**: 200 ms.
That 4× difference is why connection pooling, keep-alive, and CDNs (which move the server closer,
so RTT drops) matter so much.

**Try it** (shows the time of each step, cumulative from the start):
```bash
curl -o /dev/null -s -w 'dns:  %{time_namelookup}\ntcp:  %{time_connect}\ntls:  %{time_appconnect}\nttfb: %{time_starttransfer}\ntotal: %{time_total}\n' https://example.com
```
Run it twice. Then run it against a server far away and one close by, and compare.

---

## 7. DNS: names to IP addresses

Packets need IP addresses; humans use names. DNS is a distributed, cached lookup system,
and it usually runs over **UDP port 53** (one packet each way, no handshake).

```
Your app → OS resolver (stub) → recursive resolver (e.g. 8.8.8.8, or AWS's VPC resolver)
                                    │ cache miss? walk the hierarchy:
                                    ├→ root server:  "who handles .com?"
                                    ├→ .com server:  "who handles example.com?"
                                    └→ example.com's authoritative server: "93.184.216.34, TTL 300"
```

- The **recursive resolver** does the walking and caches answers for everyone using it.
- Every answer has a **TTL** (seconds it may be cached). Low TTL → fast failover, more lookups.
- Caching happens at many levels: browser, OS, runtime (the JVM caches DNS; Node doesn't by
  default), and the recursive resolver.

**Try it:** `dig example.com` (look at the ANSWER section and TTL), `dig +trace example.com` (the walk from root).

---

## 8. TLS: what it gives you and why it costs round trips

TLS adds three guarantees on top of TCP:
1. **Confidentiality:** nobody in the middle can read the bytes (encryption).
2. **Integrity:** nobody can change the bytes without being detected.
3. **Authentication:** you're really talking to `api.example.com` (the certificate).

Two kinds of cryptography, used together:
- **Asymmetric** (public/private key pair): slow, but lets two strangers agree on a secret over an
  open network, and lets a server prove its identity. Used **only during the handshake**.
- **Symmetric** (one shared key, e.g. AES): fast. Used for **all the actual data** after the handshake.

So the handshake's job is: prove the server's identity, and agree on a shared symmetric key.

**Certificates:** the server sends a certificate saying "this public key belongs to
`api.example.com`", signed by a **Certificate Authority (CA)**. Your OS/browser ships with a list
of trusted root CAs. The client checks the chain: server cert ← signed by intermediate CA ← signed by trusted root.

**SNI:** one IP (one load balancer) can host many domains. The client puts the hostname in its
first TLS message (ClientHello) so the server knows which certificate to send.

**How the shared key is agreed: (EC)DHE, and why the secret never travels**

In TLS 1.3 the key isn't *sent* (not even encrypted). Both sides *compute* it with Diffie-Hellman:

```
Public, known to everyone: a generator g and a big prime p

Client picks secret a (never sent)         Server picks secret b (never sent)
Client sends  A = g^a mod p   ─────────►
                              ◄─────────   Server sends  B = g^b mod p
Client computes B^a = g^(ab)               Server computes A^b = g^(ab)   ← same number
```

An eavesdropper sees `g`, `p`, `A`, and `B`, but getting `g^(ab)` from them requires solving the
discrete logarithm problem, which is infeasible for big numbers. **ECDHE** is the same idea on
elliptic curves (smaller keys, faster). The final **E = ephemeral**: `a` and `b` are fresh for
every connection and thrown away afterwards. The result goes through a key derivation function
to produce the symmetric keys (AES-GCM or ChaCha20).

**So what is the certificate's private key for?** Only for **signing** the handshake, to prove
"I am `api.example.com`" (authentication). It never encrypts the session key.

**Forward secrecy:** if an attacker records traffic today and steals the server's private key
next year, they **still can't decrypt** the recordings, because each session's key came from
ephemeral secrets that no longer exist. With the stolen key they *can* impersonate the server
from then on, so you revoke and rotate the certificate.

Contrast with **TLS 1.2's RSA key exchange** (removed in 1.3): the client encrypted the secret
with the server's public key, so a later key theft decrypted *every* recorded session. That's why
TLS 1.3 made ephemeral key exchange mandatory.

**Try it:** `openssl s_client -connect example.com:443 -servername example.com </dev/null` shows the cert chain and TLS version.

---

## 9. HTTP is just text over a TCP connection

HTTP/1.1 is human-readable text sent over the TCP byte stream:

```
GET /users HTTP/1.1\r\n
Host: api.example.com\r\n
Accept: application/json\r\n
\r\n                                ← empty line = end of headers
```

```
HTTP/1.1 200 OK\r\n
Content-Type: application/json\r\n
Content-Length: 27\r\n
\r\n
[{"id":1,"name":"Vaibhav"}]
```

- `Content-Length` tells the reader where the body ends, because TCP is a stream with no message boundaries.
- **Keep-alive** (the default in HTTP/1.1) means the TCP connection stays open after the
  response, so the next request skips DNS, TCP, and TLS.
- HTTP/1.1 allows **one request at a time** per connection. Browsers open ~6 connections per host
  to work around it. HTTP/2 fixes this with many streams on one connection.

**Try it:** `printf 'GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n' | nc example.com 80`.
You just spoke HTTP by hand. The weekend lab has you write the server side of this.

---

## 10. Who actually handles the connection: processes, threads, blocking

This is the base for "what happens to other requests when one is slow" (Day 1, Q3).

- A **process** is a running program with its own memory. A **thread** is a line of execution
  inside a process; threads share the process's memory.
- The OS **scheduler** decides which thread runs on each CPU core and switches between them.
- **Blocking I/O:** the thread calls `read()` on a socket and sleeps until data arrives. While it
  sleeps, it can't do anything else, so one thread per connection means 10k connections → 10k threads.
- **Non-blocking I/O + event loop:** one thread asks the OS "tell me which of these 10k sockets
  have data" (`epoll` on Linux) and handles whichever are ready. Cheap per connection, **but**
  if the code handling one event uses the CPU for 2 seconds, nothing else runs for 2 seconds.

How the three runtimes map onto this:

| Runtime | Model | One request burns CPU for 2s |
|---|---|---|
| Node.js | 1 JS thread + event loop (`epoll`, via libuv) | **Every** request on that process waits 2s; timers and health checks are late too |
| PHP-FPM | N worker **processes**, each handles one request at a time, with blocking I/O | Only that worker is busy. Trouble starts when **all** N workers are busy: new requests queue, then Nginx returns 502/504 |
| Go | Many **goroutines** multiplexed onto a few OS threads (the M:N scheduler); I/O uses `epoll` underneath but looks blocking in your code | The busy goroutine uses one core. The scheduler **preempts** it (since Go 1.14), so other goroutines keep running on the remaining cores, and even on the same one |

Module 02 goes deep on this (the OS scheduler, `epoll`, goroutine scheduling). For now, the key
idea is: **who is waiting, and on what?**

---

## Glossary (quick lookup)

| Term | Meaning |
|---|---|
| Packet | A chunk of bytes (≤ ~1500) sent across the network |
| MTU | Max packet size on a link (1500 for Ethernet) |
| RTT | Round-trip time: there and back |
| Port | 16-bit number that picks the program on a machine |
| Ephemeral port | Random client-side port the OS assigns per outgoing connection |
| Socket | OS handle (a file descriptor) for one end of a connection |
| 4-tuple | (src IP, src port, dst IP, dst port): uniquely identifies a TCP connection |
| ACK | "I received up to byte N" |
| Keep-alive | Reusing one TCP connection for many HTTP requests |
| TTL (DNS) | How long a DNS answer may be cached |
| SNI | Hostname sent in the TLS ClientHello so the server picks the right cert |
| TLS termination | Where TLS is decrypted (often at the load balancer) |
| `epoll` | Linux API that tells one thread which of many sockets are ready |

---

## Check yourself

1. A router forwards your HTTPS packet. Which headers can it read, and which bytes can't it read? - It can read upto L3 i.e. IP and cant read the process port
2. Your Go service connects to one Postgres server at `10.0.2.9:5432`. Why can't it have more than
   ~28k connections open to it, and what in the 4-tuple would you change to get more? - For each outgoing connection a ephemeral port is assigned, the number of connection are directly proportional number of ports vacaint. we can increase the port size to get more 4-tuple. 
3. Your users are in India and your API is in `us-east-1`. Estimate the time for a cold HTTPS request,
   then for a request on a reused connection. Which single change helps cold requests most? - Lets assume RTT for India and us-east-1 is 200ms, then on a cold start; DNS 200ms, TCP 200ms, TLS 200ms, Application/Http 200ms in total 800ms. But incase of reused connection DNS, TCP, TLS are not required only Application/Http 200ms is needed so 200ms. Reducing RTT helps cold requests most
4. Why does TLS use asymmetric crypto for the handshake but symmetric crypto for the data? - asymmetric crypto is slow but allows stranges to agree on a secrete on a open network, hence it is used to share a shared secrete. this shared secrete uses symmetric crypto which is quick so that messages can be quickly encrypted and decrypted.
