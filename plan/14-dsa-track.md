# 14 — DSA Track (every weekday, weeks 1–15, in parallel)

**Two goals:**
1. **Interviews:** solve medium problems in about 25 minutes in Go while explaining your thinking.
2. **Real knowledge:** understand how each data structure works inside and where it shows up in
   the systems you use every day. A B-tree is a database index; a heap is a timer queue; a trie
   is an HTTP router. Once you see that, DSA stops being interview trivia.

At senior level, ~120–150 well-understood problems beat 400 memorized ones.

## Weekly rhythm

| Day | What | Time |
|---|---|---|
| Mon | **Concept:** learn how this week's data structure works inside. Write notes in `notes/dsa/`. | 45 min |
| Tue | **Build it:** implement the data structure from scratch in Go, with tests, in `coding/go/ds/` | 60 min |
| Wed–Thu | **Problems:** 2 problems per day from the list, using the 5-step method | 60 min |
| Fri | **Review:** redo one problem from a previous week from memory, timed | 45 min |

The 5-step method is in [`coding/go/basics/README.md`](../coding/go/basics/README.md). If you're stuck
after 20–25 minutes, read the solution, close it, write it from memory, and log it in `mistakes.md`.
Redo it 3 days later.

---

## Week 1 — Basics: Go for DSA, Big-O
- **Concept:** Go slices, maps, strings vs runes; Big-O; the 5-step method
- **Build:** the 8 warm-ups in [`coding/go/basics`](../coding/go/basics)
- **Problems:** Two Sum, Valid Anagram (already in `coding/go/arrays/`)
- **In real systems:** Big-O is how you predict whether a query or loop survives production data sizes

## Week 2 — Arrays & hashing
- **Concept:** how a hash map works: hash function, buckets, collisions (chaining vs open addressing), load factor, resizing. How Go's map works (buckets of 8, incremental growth; Go 1.24 switched to Swiss tables). Why map iteration order is random.
- **Build:** a hash map with chaining and resizing (`Get`, `Put`, `Delete`)
- **Problems:** Contains Duplicate, Group Anagrams, Top K Frequent Elements, Product of Array Except Self, Longest Consecutive Sequence, Encode and Decode Strings
- **In real systems:** hash indexes in Postgres, Redis's main dictionary (with incremental rehashing), hash-based sharding, HashDoS attacks

## Week 3 — Two pointers + sorting
- **Concept:** merge sort vs quicksort vs heapsort; stability; why O(n log n) is the limit for comparison sorts; what Go's `slices.Sort` uses (pdqsort)
- **Build:** merge sort and quicksort, with tests and a benchmark against `slices.Sort`
- **Problems:** Valid Palindrome, Two Sum II, 3Sum, Container With Most Water, Sort Colors, Trapping Rain Water (hard; attempt it)
- **In real systems:** external merge sort when data doesn't fit in memory (databases' `ORDER BY` spilling to disk); merge joins

## Week 4 — Sliding window
- **Concept:** fixed vs variable windows; why it turns O(n·k) into O(n)
- **Build:** a sliding-window counter (count events in the last N seconds)
- **Problems:** Best Time to Buy and Sell Stock, Longest Substring Without Repeating Characters, Longest Repeating Character Replacement, Permutation in String, Minimum Window Substring (hard)
- **In real systems:** sliding-window rate limiters, TCP's flow-control window, rolling metrics (requests in the last 5 minutes)

## Week 5 — Stack & queue
- **Concept:** stack vs queue vs deque; array-backed vs linked; ring buffers; monotonic stacks
- **Build:** a fixed-size ring buffer queue (and make it safe for concurrent use)
- **Problems:** Valid Parentheses, Min Stack, Evaluate Reverse Polish Notation, Daily Temperatures, Car Fleet, Largest Rectangle in Histogram (hard)
- **In real systems:** the call stack (and stack overflows from recursion), bounded queues for backpressure, ring buffers in logging and the Linux kernel, Go's buffered channels

## Week 6 — Binary search
- **Concept:** the invariant-based template (avoid off-by-one bugs); binary search on the answer
- **Build:** `LowerBound` and `UpperBound` functions with thorough edge-case tests
- **Problems:** Binary Search, Search a 2D Matrix, Koko Eating Bananas, Find Minimum in Rotated Sorted Array, Search in Rotated Sorted Array, Time Based Key-Value Store
- **In real systems:** lookups inside B-tree pages and SSTables, `git bisect` for finding the commit that broke something, finding the max load a service can take

## Week 7 — Linked lists + LRU cache
- **Concept:** singly vs doubly linked; dummy head nodes; fast/slow pointers; skip lists
- **Build:** an **LRU cache** with a map + doubly linked list, safe for concurrent use (also a top interview question)
- **Problems:** Reverse Linked List, Merge Two Sorted Lists, Linked List Cycle, Reorder List, Remove Nth Node From End, LRU Cache
- **In real systems:** LRU eviction in caches (Redis uses an approximate LRU), skip lists behind Redis sorted sets, free lists in memory allocators

## Week 8 — Trees, BSTs & tries
- **Concept:** traversals (pre/in/post-order, level-order); BST properties; balanced trees (AVL/red-black at concept level); **B-trees and B+trees** (why databases use them instead of BSTs); tries and radix trees
- **Build:** a trie with `Insert`, `Search`, `StartsWith`
- **Problems:** Invert Binary Tree, Maximum Depth, Same Tree, Binary Tree Level Order Traversal, Validate BST, Lowest Common Ancestor of a BST, Kth Smallest in a BST, Implement Trie
- **In real systems:** database indexes (B+trees), HTTP routers (radix trees in Go's `httprouter`/`chi`), IP routing tables, autocomplete, Merkle trees (git, DynamoDB anti-entropy)

## Week 9 — Heaps & priority queues
- **Concept:** binary heap in an array; sift up/down; O(log n) push/pop; O(n) heapify; top-k pattern
- **Build:** a min-heap from scratch, then the same with Go's `container/heap`
- **Problems:** Kth Largest Element in a Stream, Last Stone Weight, K Closest Points to Origin, Task Scheduler, Merge K Sorted Lists, Find Median from Data Stream (hard)
- **In real systems:** timer queues (Go runtime timers, Node's timers), job schedulers, top-N dashboards, merging sorted files in LSM compaction, Dijkstra

## Week 10 — Graphs: BFS, DFS, topological sort, union-find
- **Concept:** adjacency list vs matrix; BFS for shortest path in unweighted graphs; DFS; cycle detection; topological sort (Kahn's algorithm); union-find with path compression
- **Build:** a dependency resolver: given tasks and dependencies, return a valid order or report a cycle
- **Problems:** Number of Islands, Clone Graph, Rotting Oranges, Pacific Atlantic Water Flow, Course Schedule, Course Schedule II, Number of Connected Components (union-find), Redundant Connection
- **In real systems:** package managers and build tools (npm, Go modules, Terraform's dependency graph), deadlock detection in databases (cycle in the wait-for graph), microservice call graphs, social graphs

## Week 11 — Weighted graphs, intervals & greedy
- **Concept:** Dijkstra with a heap; why it fails with negative edges; interval sorting patterns; when greedy is provably correct
- **Build:** Dijkstra on a small city map
- **Problems:** Network Delay Time (Dijkstra), Cheapest Flights Within K Stops, Merge Intervals, Insert Interval, Non-overlapping Intervals, Meeting Rooms II, Jump Game, Gas Station
- **In real systems:** network routing (OSPF), maps and ride-hailing ETAs, calendar and booking systems, resource allocation (how many workers or rooms do we need at peak?)

## Week 12 — Backtracking, bits & probabilistic structures
- **Concept:** choose → explore → un-choose; pruning; bit manipulation basics (AND, OR, XOR, shifts, masks); **Bloom filters** and **HyperLogLog** (concept and tradeoffs)
- **Build:** a Bloom filter with configurable size and number of hashes; measure its false-positive rate
- **Problems:** Subsets, Permutations, Combination Sum, Word Search, N-Queens (hard), Single Number, Number of 1 Bits, Counting Bits
- **In real systems:** Bloom filters in Cassandra/RocksDB to skip SSTables, cache-miss avoidance; HyperLogLog for unique counts in Redis (`PFCOUNT`); bitmaps for feature flags, permissions, and daily-active-user tracking

## Week 13 — Dynamic programming (1-D)
- **Concept:** recurrence → memoization (top-down) → tabulation (bottom-up) → space optimization
- **Build:** solve Climbing Stairs all four ways and benchmark them
- **Problems:** Climbing Stairs, House Robber, House Robber II, Coin Change, Longest Increasing Subsequence, Word Break, Decode Ways
- **In real systems:** memoization is the same idea as caching; query optimizers use DP to choose join orders

## Week 14 — Dynamic programming (2-D)
- **Concept:** grid DP; DP over two strings
- **Build:** edit distance (Levenshtein) with a function that returns the actual edit steps
- **Problems:** Unique Paths, Longest Common Subsequence, Edit Distance, Coin Change II, Best Time to Buy and Sell Stock with Cooldown
- **In real systems:** `diff` and git's diffs (LCS-based), fuzzy search and "did you mean?" (edit distance), DNA sequence alignment

## Week 15 — Review & timed mocks
- Redo every problem in `mistakes.md`
- 3 timed mock sessions with Claude: 2 problems in 60 minutes, talking out loud
- Re-implement LRU cache and a rate limiter from memory in under 30 minutes each

---

## Backend-flavored coding questions (common at senior level)
These mix DSA with concurrency and system thinking. Do them in weeks 7–15.
- [ ] LRU cache, safe for concurrent use (week 7)
- [ ] Token bucket rate limiter (module 09 lab)
- [ ] In-memory key-value store with TTL and a background cleanup goroutine
- [ ] Concurrent web crawler with a worker pool, depth limit, and dedupe
- [ ] Parse a large access log and print the top 10 endpoints by p99 latency (streaming, bounded memory)
- [ ] Merge K sorted log files into one sorted stream (heap)
- [ ] Simple job scheduler: run jobs at given times (heap + goroutines + `time.Timer`)

## Resources
- **NeetCode roadmap** (neetcode.io/roadmap): video explanations by pattern; the problem lists above follow it
- ***Grokking Algorithms*** — Aditya Bhargava: a gentle, illustrated introduction (good for weeks 1–6)
- ***The Algorithm Design Manual*** — Steven Skiena: deeper, practical, with "war stories" from real systems (dip in for weeks 8–14)
- **VisuAlgo** (visualgo.net): animations of every data structure in this list
- Go standard library: `slices`, `maps`, `sort`, `container/heap`, `container/list`, `strings.Builder`

## Tracking
- Data structures you build: `coding/go/ds/<name>/` with tests
- Problems: `coding/go/<pattern>/<problem>/` with tests, plus a row in the root README's coding table
- Concept notes: `notes/dsa/<topic>.md`
