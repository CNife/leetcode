package valid_mountain_array

import "testing"

func TestValidMountainArray(t *testing.T) {
	tests := []struct {
		array []int
		want  bool
	}{
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{0, 3, 2, 1}, true},
		{[]int{0, 1, 2, 3, 2, 1, 2, 1}, false},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, false},
		{[]int{0, 1, 2, 1, 2}, false},
	}
	for _, tt := range tests {
		if got := ValidMountainArray(tt.array); got != tt.want {
			t.Errorf("ValidMountainArray(%v) = %v, want %v", tt.array, got, tt.want)
		}
	}
}
