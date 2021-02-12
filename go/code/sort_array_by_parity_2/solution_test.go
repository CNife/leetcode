package sort_array_by_parity_2

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	arrays := [][]int{
		{4, 2, 5, 7},
	}
	for _, array := range arrays {
		got := SortArrayByParity(array)
		for i, num := range got {
			if (i^num)&1 != 0 {
				t.Errorf("parity error at %v of %v", i, got)
			}
		}
		if !reflect.DeepEqual(sortedArray(array), sortedArray(got)) {
			t.Errorf("not same array, src %v, got %v", array, got)
		}
	}
}

func sortedArray(src []int) []int {
	copied := append([]int{}, src...)
	sort.Ints(copied)
	return copied
}
