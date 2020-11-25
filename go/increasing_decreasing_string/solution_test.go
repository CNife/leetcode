package increasing_decreasing_string

import "testing"

func TestSortString(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, want string
	}{
		{"aaaabbbbcccc", "abccbaabccba"},
		{"rat", "art"},
		{"leetcode", "cdelotee"},
		{"ggggggg", "ggggggg"},
		{"spo", "ops"},
	}
	for _, tt := range tests {
		if got := SortString(tt.s); got != tt.want {
			t.Errorf("SortString(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
