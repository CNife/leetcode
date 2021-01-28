package find_pivot_index

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{-1, -1, -1, 0, 1, 1}, 0},
	}
	for _, tt := range tests {
		if got := PivotIndex(tt.nums); got != tt.want {
			t.Errorf("PivotIndex(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
