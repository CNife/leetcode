package count_of_range_sum

import "testing"

func TestCountRangeSum(t *testing.T) {
	tests := []struct {
		nums          []int
		lower, higher int
		want          int
	}{
		{
			nums:   []int{-2, 5, -1},
			lower:  -2,
			higher: 2,
			want:   3,
		},
	}
	for _, tt := range tests {
		got := CountRangeSum(tt.nums, tt.lower, tt.higher)
		if got != tt.want {
			t.Errorf("CountRangeSum(%v, %v, %v) = %v, want %v",
				tt.nums, tt.lower, tt.higher, got, tt.want)
		}
	}
}
