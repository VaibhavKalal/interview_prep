# Mistakes Log

The most valuable file in this repo. Every time you get something wrong
(in a mock, a coding problem, or a quiz), write it down here.
Review it every week.

| Date | Topic | What I said / did | Correct answer | Why it matters |
|---|---|---|---|---|
| 2026-11-24 | Networking: OSI layers | Called an L4 load balancer a "network layer" LB | L4 is the **transport** layer (TCP/UDP, ports). L3 is the network layer (IP). | Interviewers use L3/L4/L7 precisely. Mixing them up undermines everything you say after it. |
| 2026-11-24 | Networking: request lifecycle | Put the load balancer **after** the TCP and TLS handshakes in the lifecycle | With an ALB or Nginx, the client's TCP and TLS handshakes are **with the load balancer**. The LB then opens its own (often pooled) connection to the app. | It explains why the LB's idle timeout, TLS certs, and keep-alive settings matter, and why there are two separate connections to debug. |
| 2026-11-24 | Networking: cold requests | Listed DNS + TCP + TLS as the reasons a first request is slow, but no round-trip counts and no TCP slow start | Cold HTTPS ≈ 4 RTTs (DNS 1+, TCP 1, TLS 1.3 1 / TLS 1.2 2, HTTP 1) vs 1 RTT warm. Also TCP slow start (small initial window ~10 packets) and server-side cold caches/pools. | Senior answers quantify the cost. "Round trips dominate latency" is the reason for keep-alive, pooling, and CDNs. |
| 2026-11-24 | HTTP: LB status codes | Said the ALB returns 502 when its idle timeout fires before the app responds | **504 Gateway Timeout** = backend too slow (idle timeout hit). **502 Bad Gateway** = backend closed/reset the connection or sent an invalid response (e.g. keep-alive mismatch). The idle timer resets whenever bytes flow. | The code tells you where to look during an incident: 504 → slow app/DB; 502 → connection handling or crashes. |
| 2026-11-24 | HTTP: keep-alive mismatch | Knew "Node closes at 5s, ALB reuses it" but couldn't explain why the ALB sends on a closed connection, the fix, or why it's intermittent | Race: Node sends FIN at 5s while the ALB is already writing a request on that pooled connection → Node's kernel replies RST → ALB returns 502, app logs nothing. Fix: backend keep-alive > LB idle timeout (Node `keepAliveTimeout` 65s, `headersTimeout` 66s vs ALB 60s), so the LB always closes first. Intermittent because the request must land in a ms-wide window; worst under low/bursty traffic. | Classic "random 502s, no app errors" production bug. Rule: each hop's idle timeout > the hop in front of it. Evidence is in ALB logs (`target_status_code = -`). |
