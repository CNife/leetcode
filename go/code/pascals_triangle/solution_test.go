package pascals_triangle

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		n    int
		want [][]int
	}{
		{
			n: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		if got := Generate(tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Generate(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
