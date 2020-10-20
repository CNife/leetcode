package valid_number

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"0", true},
		{" 0.1 ", true},
		{"abc", false},
		{"1 a", false},
		{"2e10", true},
		{" -90e3 ", true},
		{" 1e", false},
		{"e3", false},
		{" 6e-1", true},
		{" 99e2.5", false},
		{"53.5e93", true},
		{" --6 ", false},
		{"-+3", false},
		{"95a54e53", false},
	}
	for _, tt := range tests {
		if got := IsNumber(tt.input); got != tt.want {
			t.Errorf("IsNumber(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
