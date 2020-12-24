package find_the_difference

import "testing"

func TestFindTheDifference(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t string
		want byte
	}{
		{"abcd", "abcde", 'e'},
		{"", "y", 'y'},
		{"a", "aa", 'a'},
		{"ae", "aea", 'a'},
	}
	for _, tt := range tests {
		if got := FindTheDifference(tt.s, tt.t); got != tt.want {
			t.Errorf("FindTheDifference(%v, %v) = %c, want %c",
				tt.s, tt.t, got, tt.want)
		}
	}
}
