package target_sum

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			want:   5,
		},
	}
	for _, tt := range tests {
		if got := FindTargetSumWays(tt.nums, tt.target); got != tt.want {
			t.Error(tt, got)
		}
	}
}
