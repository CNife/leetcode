package reorganize_string

import "testing"

func TestReorganizeString(t *testing.T) {
	tests := []struct {
		src, want string
	}{
		{"aab", "aba"},
		{"aaab", ""},
	}
	for _, tt := range tests {
		if got := ReorganizeString(tt.src); got != tt.want {
			t.Errorf("ReorganizeString(%v) = %v, want %v",
				tt.src, got, tt.want)
		}
	}
}
