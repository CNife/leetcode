package search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums         []int
		target, want int
	}{
		{nil, 0, 0},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, tt := range tests {
		if got := SearchInsert(tt.nums, tt.target); got != tt.want {
			t.Error(tt, got)
		}
	}
}
