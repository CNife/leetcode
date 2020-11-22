package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "cat", false},
	}
	for _, tt := range tests {
		if got := IsAnagram(tt.s, tt.t); got != tt.want {
			t.Errorf("IsAnagram(%v, %v) = %v, want %v",
				tt.s, tt.t, got, tt.want)
		}
	}
}
