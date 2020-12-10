package first_missing_positive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, 1, -1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}
	for _, tt := range tests {
		if got := FirstMissingPositive(tt.nums); got != tt.want {
			t.Errorf("FirstMissingPositive(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
