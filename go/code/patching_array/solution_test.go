package patching_array

import "testing"

func TestMinPatches(t *testing.T) {
	tests := []struct {
		nums []int
		n    int
		want int
	}{
		{[]int{1, 3}, 6, 1},
		{[]int{1, 5, 10}, 20, 2},
		{[]int{1, 2, 2}, 5, 0},
	}
	for _, tt := range tests {
		if got := MinPatches(tt.nums, tt.n); got != tt.want {
			t.Errorf("MinPatches(%v, %v) = %v, want %v",
				tt.nums, tt.n, got, tt.want)
		}
	}
}
