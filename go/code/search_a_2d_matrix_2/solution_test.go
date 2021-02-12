package search_a_2d_matrix_2

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			want:   true,
		},
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			want:   false,
		},
	}
	for _, tt := range tests {
		if got := SearchMatrix(tt.matrix, tt.target); got != tt.want {
			t.Error(tt, got)
		}
	}
}
