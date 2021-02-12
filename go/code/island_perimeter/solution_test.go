package island_perimeter

import "testing"

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		if got := IslandPerimeter(tt.grid); got != tt.want {
			t.Errorf("IslandPerimeter(%v) = %v, want %v", tt.grid, got, tt.want)
		}
	}
}
