package score_after_flipping_matrix

import "testing"

func TestMatrixScore(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}, 39},
	}
	for _, tt := range tests {
		if got := MatrixScore(tt.matrix); got != tt.want {
			t.Errorf("MatrixScore(%v) = %v, want %v", tt.matrix, got, tt.want)
		}
	}
}
