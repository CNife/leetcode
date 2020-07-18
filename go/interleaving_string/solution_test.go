package interleaving_string

import "testing"

func TestIsInterleave(t *testing.T) {
	//noinspection SpellCheckingInspection
	tests := []struct {
		lhs, rhs, target string
		want             bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"a", "", "a", true},
		{"a", "", "b", false},
	}
	for _, tt := range tests {
		if got := IsInterleave(tt.lhs, tt.rhs, tt.target); got != tt.want {
			t.Error(tt, got)
		}
	}
}
