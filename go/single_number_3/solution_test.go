package single_number_3

import (
	"reflect"
	"sort"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
	}
	for _, tt := range tests {
		got := SingleNumber(tt.nums)
		sort.Ints(got)
		sort.Ints(tt.want)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SingleNumber(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
