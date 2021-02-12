package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n, want int
	}{
		{23, 12, 193536720},
		{3, 7, 28},
		{3, 2, 3},
		{7, 3, 28},
		{3, 3, 6},
	}
	for _, tt := range tests {
		if got := UniquePaths(tt.m, tt.n); got != tt.want {
			t.Errorf("UniquePaths(%v, %v) = %v, want %v",
				tt.m, tt.n, got, tt.want)
		}
	}
}
