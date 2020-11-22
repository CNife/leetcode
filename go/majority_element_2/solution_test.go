package majority_element_2

import (
	"reflect"
	"sort"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1}, []int{1}},
		{[]int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		got := MajorityElement(tt.nums)
		sort.Ints(got)
		sort.Ints(tt.want)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MajorityElement(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
