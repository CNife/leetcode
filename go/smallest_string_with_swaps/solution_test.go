package smallest_string_with_swaps

import "testing"

func TestSmallestStringWithSwaps(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s     string
		pairs [][]int
		want  string
	}{
		{"dcab", [][]int{{0, 3}, {1, 2}}, "bacd"},
		{"dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}, "abcd"},
		{"cba", [][]int{{0, 1}, {1, 2}}, "abc"},
	}
	for _, tt := range tests {
		got := SmallestStringWithSwaps(tt.s, tt.pairs)
		if got != tt.want {
			t.Errorf("SmallestStringWithSwaps(%v, %v) = %v, want %v",
				tt.s, tt.pairs, got, tt.want)
		}
	}
}
