package fair_candy_swap

import (
	"reflect"
	"testing"
)

func TestFairCandySwap(t *testing.T) {
	tests := []struct {
		a, b, want []int
	}{
		{
			a:    []int{1, 1},
			b:    []int{2, 2},
			want: []int{1, 2},
		},
		{
			a:    []int{1, 2},
			b:    []int{2, 3},
			want: []int{1, 2},
		},
		{
			a:    []int{2},
			b:    []int{1, 3},
			want: []int{2, 3},
		},
		{
			a:    []int{1, 2, 5},
			b:    []int{2, 4},
			want: []int{5, 4},
		},
	}
	for _, tt := range tests {
		got := FairCandySwap(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FairCandySwap(%v, %v) = %v, want %v",
				tt.a, tt.b, got, tt.want)
		}
	}
}
