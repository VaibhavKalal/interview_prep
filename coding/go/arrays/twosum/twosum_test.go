package twosum

import (
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"basic", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"not at start", []int{3, 2, 4}, 6, []int{1, 2}},
		{"same value twice", []int{3, 3}, 6, []int{0, 1}},
		{"negatives", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if len(got) != 2 {
				t.Fatalf("TwoSum(%v, %d) = %v, want 2 indices", tt.nums, tt.target, got)
			}
			sort.Ints(got)
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
