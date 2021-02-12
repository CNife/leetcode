package longest_mountain_in_array

import "testing"

func TestLongestMountain(t *testing.T) {
	tests := []struct {
		array []int
		want  int
	}{
		{[]int{2, 1, 4, 7, 3, 2, 5}, 5},
		{[]int{2, 2, 2}, 0},
	}
	for _, tt := range tests {
		if got := LongestMountain(tt.array); got != tt.want {
			t.Errorf("LongestMountain(%v) = %v, want %v", tt.array, got, tt.want)
		}
	}
}
