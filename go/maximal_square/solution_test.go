package maximal_square

import "testing"

func TestMaximalSquare(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{
				{'0', '0', '1', '0'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '0'},
				{'1', '1', '0', '0'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '0'},
			},
			want: 9,
		},
		{
			matrix: [][]byte{
				{'1', '0'},
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
			want: 4,
		},
		{
			matrix: [][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			want: 1,
		},
		{
			matrix: [][]byte{{'0'}},
			want:   0,
		},
		{
			matrix: [][]byte{{'1'}},
			want:   1,
		},
		{
			matrix: [][]byte{{'0', '0'}, {'0', '0'}},
			want:   0,
		},
	}
	for _, tt := range tests {
		if got := MaximalSquare(tt.matrix); got != tt.want {
			t.Errorf("MaximalSquare(%v) = %v, want %v",
				tt.matrix, got, tt.want)
		}
	}
}
