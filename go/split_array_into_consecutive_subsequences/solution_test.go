package split_array_into_consecutive_subsequences

import (
	"testing"
)

func TestIsPossible(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 3, 4, 5}, true},
		{[]int{1, 2, 3, 3, 4, 4, 5, 5}, true},
		{[]int{1, 2, 3, 4, 4, 5}, false},
	}
	for _, tt := range tests {
		if got := IsPossible(tt.nums); got != tt.want {
			t.Errorf("IsPossible(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
