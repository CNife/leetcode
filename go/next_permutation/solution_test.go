package next_permutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
	}
	for _, tt := range tests {
		got := append([]int{}, tt.nums...)
		NextPermutation(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NextPermutation(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
