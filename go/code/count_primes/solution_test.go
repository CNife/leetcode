package count_primes

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{10, 4},
	}
	for _, tt := range tests {
		if got := CountPrimes(tt.n); got != tt.want {
			t.Errorf("CountPrimes(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
