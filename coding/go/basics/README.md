# Coding Basics — Start Here (before LeetCode)

LeetCode feels impossible at first because it tests three things at once:
**the language**, **complexity thinking**, and **patterns**. Learn them one at a time.

Work through this file top to bottom, then do the exercises in `basics.go`.

---

## Part 1 — Go essentials for coding problems (30 min)

You only need a small part of Go for DSA. Type these out in a scratch file and run them.

### Slices (dynamic arrays)
```go
nums := []int{3, 1, 4}         // create
nums = append(nums, 5)         // add to end
n := len(nums)                 // length
first, last := nums[0], nums[n-1]
part := nums[1:3]              // [1, 4]  (start inclusive, end exclusive)
grid := make([][]int, 3)       // 2D: 3 rows, fill each with make([]int, cols)
buf := make([]int, 0, 10)      // empty, capacity 10
```

### Loops (Go has only `for`)
```go
for i := 0; i < len(nums); i++ { }   // classic
for i, v := range nums { }           // index + value
for _, v := range nums { }           // value only
for left < right { }                 // while-loop
```

### Maps (hash maps): the most important tool in interviews
```go
count := map[int]int{}         // create
count[5]++                     // missing keys read as 0, so this just works
v, ok := count[7]              // ok == false if 7 is not in the map
delete(count, 5)
for k, v := range count { }    // order is random!

seen := map[int]bool{}         // a "set"
seen[3] = true
if seen[3] { }
```

### Strings: bytes vs runes
```go
s := "héllo"
len(s)                 // 6 — counts BYTES, not characters ('é' is 2 bytes)
s[0]                   // a byte (uint8)
for i, r := range s {} // r is a rune (a full Unicode character)
r := []rune(s)         // convert to runes to index characters safely
string(r)              // back to string

var sb strings.Builder // build strings efficiently in a loop
sb.WriteString("hi")
sb.WriteByte('!')
result := sb.String()
```

### Sorting
```go
sort.Ints(nums)                          // in place
sort.Slice(people, func(i, j int) bool { // custom order
    return people[i].Age < people[j].Age
})
```

### Multiple return values (very common in Go)
```go
func minMax(nums []int) (int, int) { ... }
lo, hi := minMax(nums)
```

---

## Part 2 — Big-O in 10 minutes

Big-O answers: **how does the work grow when the input grows?**

| Big-O | Name | Example | n = 1,000,000 |
|---|---|---|---|
| O(1) | constant | map lookup, `nums[i]` | instant |
| O(log n) | logarithmic | binary search | ~20 steps |
| O(n) | linear | one loop over the input | fine |
| O(n log n) | linearithmic | sorting | fine |
| O(n²) | quadratic | loop inside a loop | too slow (10¹² steps) |
| O(2ⁿ) | exponential | all subsets | impossible |

**Rules of thumb:**
- One loop over `n` → O(n). A loop inside a loop → O(n²).
- Map lookup/insert → O(1) on average. That's why maps turn O(n²) into O(n).
- Sorting → O(n log n).
- **Space** matters too: a map that holds every element is O(n) extra space.

**The most common interview move:** "The brute force is O(n²) with two loops.
I can make it O(n) by storing what I've seen in a map."

---

## Part 3 — How to approach any problem (use this every time)

1. **Understand:** restate the problem. Ask about edge cases: empty input? duplicates? negatives? huge input?
2. **Examples:** work through 1–2 small examples by hand, on paper.
3. **Brute force first:** say it out loud with its Big-O, even if it's slow. It proves you understand the problem.
4. **Optimize:** what is repeated work? Can a map, sorting, or two pointers remove it?
5. **Code**, then **test** by walking through your example line by line.

If you're stuck after 20–25 minutes, read the solution, understand it, then
**close it and write it yourself from scratch**. Come back to it 3 days later.
This is how everyone learns patterns; it is not cheating.

---

## Part 4 — Exercises (in `basics.go`)

Easier than LeetCode. Each one teaches a building block you'll reuse.

| # | Function | Teaches | Run |
|---|---|---|---|
| 1 | `Sum` | loops | `go test ./coding/go/basics/ -run TestSum` |
| 2 | `Max` | edge cases, multiple returns | `-run TestMax` |
| 3 | `CountFrequency` | maps (counting) | `-run TestCountFrequency` |
| 4 | `ContainsDuplicate` | map as a set → O(n) | `-run TestContainsDuplicate` |
| 5 | `ReverseInPlace` | two pointers | `-run TestReverseInPlace` |
| 6 | `ReverseString` | runes vs bytes | `-run TestReverseString` |
| 7 | `IsPalindrome` | two pointers on strings | `-run TestIsPalindrome` |
| 8 | `MaxSumOfK` | sliding window | `-run TestMaxSumOfK` |

For each exercise: write the Big-O of your solution as a comment above it.

**After all 8 pass**, you already know the ideas behind Two Sum (#4's map trick)
and Valid Anagram (#3's counting). Go do them in `../arrays/`.
