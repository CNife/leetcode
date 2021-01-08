package rotate_array

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		got := cloneSlice(tt.nums)
		Rotate(got, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Rotate(%v, %v) = %v, want %v",
				tt.nums, tt.k, got, tt.want)
		}
	}
}

func cloneSlice(slice []int) []int {
	result := make([]int, len(slice))
	for i, num := range slice {
		result[i] = num
	}
	return result
}
