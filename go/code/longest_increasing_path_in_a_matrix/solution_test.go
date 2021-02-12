package longest_increasing_path_in_a_matrix

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			want: 4,
		},
		{
			matrix: [][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		if got := LongestIncreasingPath(tt.matrix); got != tt.want {
			t.Error(tt, got)
		}
	}
}
