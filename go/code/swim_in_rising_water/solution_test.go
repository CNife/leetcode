package swim_in_rising_water

import "testing"

func TestSwimInWater(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{0, 2},
				{1, 3},
			},
			want: 3,
		},
		{
			grid: [][]int{
				{0, 1, 2, 3, 4},
				{24, 23, 22, 21, 5},
				{12, 13, 14, 15, 16},
				{11, 17, 18, 19, 20},
				{10, 9, 8, 7, 6},
			},
			want: 16,
		},
		{
			grid: [][]int{
				{10, 12, 4, 6},
				{9, 11, 3, 5},
				{1, 7, 13, 8},
				{2, 0, 15, 14},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		if got := SwimInWater(tt.grid); got != tt.want {
			t.Errorf("SwimInWater(%v) = %v, want %v",
				tt.grid, got, tt.want)
		}
	}
}
