package word_pattern

import "testing"

func TestWordPattern(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		pattern, s string
		want       bool
	}{
		{"abc", "dog cat dog", false},
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	}
	for _, tt := range tests {
		if got := WordPattern(tt.pattern, tt.s); got != tt.want {
			t.Errorf("WordPattern(%v, %v) = %v, want %v",
				tt.pattern, tt.s, got, tt.want)
		}
	}
}
