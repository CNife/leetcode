package non_overlapping_intervals

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
	}
	for _, tt := range tests {
		if got := EraseOverlapIntervals(tt.intervals); got != tt.want {
			t.Errorf("EraseOverlapIntervals(%v) = %v, want %v",
				tt.intervals, got, tt.want)
		}
	}
}
