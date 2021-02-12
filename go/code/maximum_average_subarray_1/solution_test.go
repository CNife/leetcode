package maximum_average_subarray_1

import "testing"

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
	}
	for _, tt := range tests {
		got := FindMaxAverage(tt.nums, tt.k)
		if got-tt.want >= 1e-5 || tt.want-got >= 1e-5 {
			t.Errorf("FindMaxAverage(%v, %v) = %v, want %v",
				tt.nums, tt.k, got, tt.want)
		}
	}
}
