package isomorphic_strings

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}
	for _, tt := range tests {
		if got := IsIsomorphic(tt.s, tt.t); got != tt.want {
			t.Errorf("IsIsomorphic(%v, %v) = %v, want %v",
				tt.s, tt.t, got, tt.want)
		}
	}
}
