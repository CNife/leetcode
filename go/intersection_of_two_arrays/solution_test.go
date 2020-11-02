package intersection_of_two_arrays

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		arr1, arr2, want []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}
	for _, tt := range tests {
		got := Intersection(tt.arr1, tt.arr2)
		sort.Ints(got)
		sort.Ints(tt.want)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Intersection(%v, %v) = %v, want %v", tt.arr1, tt.arr2, got, tt.want)
		}
	}
}
