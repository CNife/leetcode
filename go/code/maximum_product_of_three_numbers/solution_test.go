package maximum_product_of_three_numbers

import "testing"

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
	}
	for _, tt := range tests {
		if got := MaximumProduct(tt.nums); got != tt.want {
			t.Errorf("MaximumProduct(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
