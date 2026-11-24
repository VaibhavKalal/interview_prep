# 14 — DSA Track (every weekday, weeks 1–14, in parallel)

**Goal:** solve medium problems in about 25 minutes while explaining your thinking, in Go.
At senior level, ~100–150 well-understood problems beat 400 memorized ones.

**Daily routine (45–60 min):** 1 new problem + revisit 1 old problem from memory.
Use the 5-step method from [`coding/go/basics/README.md`](../coding/go/basics/README.md).

## Schedule by pattern

| Week | Pattern | Learn | Practice (examples from NeetCode 150) |
|---|---|---|---|
| 1 | **Basics** | Go for DSA, Big-O | [`coding/go/basics`](../coding/go/basics) 8 warm-ups, Two Sum, Valid Anagram |
| 2 | Arrays & hashing | maps, sets, counting | Contains Duplicate, Group Anagrams, Top K Frequent, Product of Array Except Self, Longest Consecutive Sequence |
| 3 | Two pointers | sorted inputs, converging pointers | Valid Palindrome, Two Sum II, 3Sum, Container With Most Water |
| 4 | Sliding window | fixed and variable windows | Best Time to Buy/Sell Stock, Longest Substring Without Repeating, Character Replacement, Permutation in String |
| 5 | Stack | monotonic stack | Valid Parentheses, Min Stack, Daily Temperatures, Evaluate RPN |
| 6 | Binary search | search on answer | Binary Search, Search in Rotated Array, Koko Eating Bananas, Find Min in Rotated Array |
| 7 | Linked lists | fast/slow pointers, dummy nodes | Reverse List, Merge Two Lists, Linked List Cycle, Reorder List, **LRU Cache** |
| 8 | Trees | DFS/BFS recursion | Invert Tree, Max Depth, Same Tree, Level Order, Validate BST, LCA |
| 9 | Heaps | `container/heap` in Go | Kth Largest, K Closest Points, Task Scheduler, Find Median from Data Stream |
| 10 | Graphs | BFS/DFS on grids and adjacency lists | Number of Islands, Clone Graph, Rotting Oranges, Course Schedule (topological sort) |
| 11 | Intervals + greedy | sorting by start/end | Merge Intervals, Insert Interval, Non-overlapping Intervals, Meeting Rooms II |
| 12 | Backtracking | choose/explore/unchoose | Subsets, Permutations, Combination Sum, Word Search |
| 13 | 1-D dynamic programming | recurrence → memo → table | Climbing Stairs, House Robber, Coin Change, Longest Increasing Subsequence |
| 14 | Review | redo missed problems | Everything in `mistakes.md`; timed mock problems |

## Backend-flavored coding questions (common at senior level, do in weeks 7–13)
- [ ] LRU cache (with concurrency safety)
- [ ] Rate limiter (token bucket)
- [ ] Concurrent web crawler with a worker pool
- [ ] In-memory key-value store with TTL
- [ ] Parse and aggregate a log file (top N endpoints by p99 latency)
- [ ] Merge K sorted streams

## Resources
- NeetCode roadmap (neetcode.io/roadmap): video explanations by pattern
- *Learn Go with Tests* for idiomatic Go
- `container/heap`, `sort`, `slices`, `maps` packages in the Go standard library

## Tracking
Add each solved problem to `coding/go/<pattern>/<problem>/` with tests, and to the coding
table in the root README. Log each problem you couldn't solve in `mistakes.md`, then redo it 3 days later.
