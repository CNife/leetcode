package rotate_image

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		matrix, want [][]int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}
	for _, tt := range tests {
		m, n := len(tt.matrix), len(tt.matrix[0])
		clonedMatrix := make([][]int, m)
		for i := range tt.matrix {
			clonedMatrix[i] = make([]int, n)
			for j := range clonedMatrix[i] {
				clonedMatrix[i][j] = tt.matrix[i][j]
			}
		}

		Rotate(clonedMatrix)
		if !reflect.DeepEqual(clonedMatrix, tt.want) {
			t.Errorf("Rotate(%v) = %v, want %v", tt.matrix, clonedMatrix, tt.want)
		}
	}
}
