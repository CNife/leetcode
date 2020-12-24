package first_unique_character_in_a_string

import "testing"

func TestFirstUnique(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
	}
	for _, tt := range tests {
		if got := FirstUnique(tt.s); got != tt.want {
			t.Errorf("FirstUnique(%v) = %v, want %v",
				tt.s, got, tt.want)
		}
	}
}
