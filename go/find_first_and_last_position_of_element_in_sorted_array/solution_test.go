package find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{nil, 0, []int{-1, -1}},
	}
	for _, tt := range tests {
		got := SearchRange(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SearchRange(%v, %v) = %v, want %v",
				tt.nums, tt.target, got, tt.want)
		}
	}
}
