package fibonacci_number

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for _, tt := range tests {
		if got := Fib(tt.n); got != tt.want {
			t.Errorf("Fib(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
