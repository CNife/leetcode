package min_cost_to_connect_all_points

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			want:   20,
		},
		{
			points: [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			want:   18,
		},
		{
			points: [][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}},
			want:   4,
		},
		{
			points: [][]int{{-1000000, -1000000}, {1000000, 1000000}},
			want:   4000000,
		},
		{
			points: [][]int{{0, 0}},
			want:   0,
		},
	}
	for _, tt := range tests {
		got := MinCostConnectPoints(tt.points)
		if got != tt.want {
			t.Errorf("MinCostConnectPoints(%v) = %v, want %v",
				tt.points, got, tt.want)
		}
	}
}
