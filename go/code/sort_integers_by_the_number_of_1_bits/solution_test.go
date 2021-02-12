package sort_integers_by_the_number_of_1_bits

import (
	"math/bits"
	"reflect"
	"testing"
)

func TestSortByBits(t *testing.T) {
	tests := []struct {
		array, want []int
	}{
		{
			array: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			want:  []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			array: []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			want:  []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
		{
			array: []int{10000, 10000},
			want:  []int{10000, 10000},
		},
		{
			array: []int{2, 3, 5, 7, 11, 13, 17, 19},
			want:  []int{2, 3, 5, 17, 7, 11, 13, 19},
		},
		{
			array: []int{10, 100, 1000, 10000},
			want:  []int{10, 100, 10000, 1000},
		},
	}
	for _, tt := range tests {
		if got := SortByBits(tt.array); !reflect.DeepEqual(intoBits(got), intoBits(tt.want)) {
			t.Errorf("SortByBits(%v) = %v, want %v", tt.array, got, tt.want)
		}
	}
}

func intoBits(array []int) []int {
	result := make([]int, len(array))
	for i, num := range array {
		result[i] = bits.OnesCount(uint(num))
	}
	return result
}
