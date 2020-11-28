package reverse_pairs

import "testing"

func TestReversePairs(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 3, 1}, 2},
		{[]int{2, 4, 3, 5, 1}, 3},
	}
	for _, tt := range tests {
		if got := ReversePairs(tt.nums); got != tt.want {
			t.Errorf("ReversePairs(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
