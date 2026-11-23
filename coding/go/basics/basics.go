// Package basics contains warm-up exercises to do before LeetCode.
// Read README.md in this folder first.
//
// Each function currently returns a placeholder value, so its tests fail.
// Replace the body with your solution, and write its Big-O above it.
package basics

// 1. Sum returns the sum of all numbers in nums. Sum of an empty slice is 0.
//
// Big-O: time ?, space ?
func Sum(nums []int) int {
	return -1 // TODO
}

// 2. Max returns the largest number in nums.
// If nums is empty, return 0, false. Otherwise return the max, true.
// Careful: all numbers might be negative.
//
// Big-O: time ?, space ?
func Max(nums []int) (int, bool) {
	return 0, false // TODO
}

// 3. CountFrequency returns how many times each number appears.
// Example: [1, 2, 2, 3, 3, 3] → {1: 1, 2: 2, 3: 3}
//
// Big-O: time ?, space ?
func CountFrequency(nums []int) map[int]int {
	return nil // TODO
}

// 4. ContainsDuplicate reports whether any number appears at least twice.
// (This is LeetCode #217.) First think of the O(n²) way, then do it in O(n).
//
// Big-O: time ?, space ?
func ContainsDuplicate(nums []int) bool {
	return false // TODO
}

// 5. ReverseInPlace reverses nums without creating a new slice.
// Hint: one index at the start, one at the end, swap, move both inward.
// In Go you can swap with: a, b = b, a
//
// Big-O: time ?, space ?
func ReverseInPlace(nums []int) {
	// TODO
}

// 6. ReverseString returns s reversed. It must work for Unicode,
// e.g. "héllo" → "olléh". Hint: convert to []rune first.
//
// Big-O: time ?, space ?
func ReverseString(s string) string {
	return "" // TODO
}

// 7. IsPalindrome reports whether s reads the same forwards and backwards.
// Only lowercase letters a–z will be given. Empty string is a palindrome.
// Do it with two pointers, without building a reversed copy.
//
// Big-O: time ?, space ?
func IsPalindrome(s string) bool {
	return false // TODO
}

// 8. MaxSumOfK returns the largest sum of any k consecutive numbers.
// Example: nums = [2, 1, 5, 1, 3, 2], k = 3 → 9 (5 + 1 + 3)
// You can assume 1 <= k <= len(nums).
//
// Brute force: sum every window from scratch → O(n·k).
// Sliding window: when the window moves right by one, add the new number
// and subtract the one that left → O(n).
//
// Big-O: time ?, space ?
func MaxSumOfK(nums []int, k int) int {
	return 0 // TODO
}
