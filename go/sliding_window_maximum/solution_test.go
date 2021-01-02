package sliding_window_maximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{9, 11}, 2, []int{11}},
		{[]int{4, -2}, 2, []int{4}},
	}
	for _, tt := range tests {
		if got := MaxSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MaxSlidingWindow(%v, %v) = %v, want %v",
				tt.nums, tt.k, got, tt.want)
		}
	}
}
