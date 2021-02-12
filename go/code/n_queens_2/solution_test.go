package n_queens_2

import "testing"

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{4, 2},
	}
	for _, tt := range tests {
		if got := TotalNQueens(tt.n); got != tt.want {
			t.Errorf("TotalNQueen(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
