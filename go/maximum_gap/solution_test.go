package maximum_gap

import "testing"

func TestMaximumGap(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nil, 0},
		{[]int{1}, 0},
		{[]int{1, 10000000}, 9999999},
		{[]int{3, 6, 9, 1}, 3},
	}
	for _, tt := range tests {
		if got := MaximumGap(tt.nums); got != tt.want {
			t.Errorf("MaximumGap(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
