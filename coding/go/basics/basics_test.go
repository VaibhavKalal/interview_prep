package basics

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{}, 0},
		{[]int{-5, 5, 10}, 10},
	}
	for _, tt := range tests {
		if got := Sum(tt.nums); got != tt.want {
			t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		nums   []int
		want   int
		wantOK bool
	}{
		{[]int{3, 9, 2}, 9, true},
		{[]int{-7, -3, -10}, -3, true},
		{[]int{42}, 42, true},
		{[]int{}, 0, false},
	}
	for _, tt := range tests {
		got, ok := Max(tt.nums)
		if got != tt.want || ok != tt.wantOK {
			t.Errorf("Max(%v) = %d, %v; want %d, %v", tt.nums, got, ok, tt.want, tt.wantOK)
		}
	}
}

func TestCountFrequency(t *testing.T) {
	tests := []struct {
		nums []int
		want map[int]int
	}{
		{[]int{1, 2, 2, 3, 3, 3}, map[int]int{1: 1, 2: 2, 3: 3}},
		{[]int{7}, map[int]int{7: 1}},
		{[]int{}, map[int]int{}},
	}
	for _, tt := range tests {
		got := CountFrequency(tt.nums)
		if got == nil || !reflect.DeepEqual(got, tt.want) {
			t.Errorf("CountFrequency(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{}, false},
		{[]int{5, 5}, true},
	}
	for _, tt := range tests {
		if got := ContainsDuplicate(tt.nums); got != tt.want {
			t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		in := append([]int{}, tt.nums...)
		ReverseInPlace(in)
		if !reflect.DeepEqual(in, tt.want) {
			t.Errorf("ReverseInPlace(%v) gave %v, want %v", tt.nums, in, tt.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct{ s, want string }{
		{"hello", "olleh"},
		{"héllo", "olléh"},
		{"a", "a"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := ReverseString(tt.s); got != tt.want {
			t.Errorf("ReverseString(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"racecar", true},
		{"abba", true},
		{"abc", false},
		{"a", true},
		{"", true},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.s); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}

func TestMaxSumOfK(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 1, 5, 1, 3, 2}, 3, 9},
		{[]int{1, 2, 3}, 3, 6},
		{[]int{4}, 1, 4},
		{[]int{-1, -2, -3, -4}, 2, -3},
	}
	for _, tt := range tests {
		if got := MaxSumOfK(tt.nums, tt.k); got != tt.want {
			t.Errorf("MaxSumOfK(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
