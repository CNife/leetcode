package minimum_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t, want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
	}
	for _, tt := range tests {
		if got := MinWindow(tt.s, tt.t); got != tt.want {
			t.Errorf("MinWindow(%v, %v) = %v, want %v",
				tt.s, tt.t, got, tt.want)
		}
	}
}
