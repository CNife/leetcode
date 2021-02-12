package largest_perimeter_triangle

import "testing"

func TestLargestPerimeter(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 1, 2}, 5},
		{[]int{1, 2, 1}, 0},
		{[]int{3, 2, 3, 4}, 10},
		{[]int{3, 6, 2, 3}, 8},
	}
	for _, tt := range tests {
		if got := LargestPerimeter(tt.nums); got != tt.want {
			t.Errorf("LargestPerimeter(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
