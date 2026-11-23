package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"anagram", "anagram", "nagaram", true},
		{"not anagram", "rat", "car", false},
		{"different lengths", "a", "ab", false},
		{"empty strings", "", "", true},
		{"same letters different counts", "aab", "abb", false},
		{"unicode", "héllo", "lléoh", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
