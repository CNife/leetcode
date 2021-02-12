package maximal_rectangle

import "testing"

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			want: 1,
		},
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 6,
		},
		{nil, 0},
		{[][]byte{{'0'}}, 0},
		{[][]byte{{'1'}}, 1},
		{[][]byte{{'0', '0'}}, 0},
	}
	for _, tt := range tests {
		if got := MaximalRectangle(tt.matrix); got != tt.want {
			t.Errorf("MaximalRectangle(%v) = %v, want %v",
				tt.matrix, got, tt.want)
		}
	}
}
