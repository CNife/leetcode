package min_cost_climbing_stairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		costs []int
		want  int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tt := range tests {
		if got := MinCostClimbingStairs(tt.costs); got != tt.want {
			t.Errorf("MinCostClimbingStairs(%v) = %v, want %v",
				tt.costs, got, tt.want)
		}
	}
}
