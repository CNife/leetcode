package min_in_spiral_array

import "testing"

func TestMinArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
	}
	for _, tt := range tests {
		if got := MinArray(tt.nums); got != tt.want {
			t.Error(tt, got)
		}
	}
}
