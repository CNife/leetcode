package dungeon_game

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct {
		dungeon [][]int
		want    int
	}{
		{
			dungeon: [][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		if got := CalculateMinimumHP(tt.dungeon); got != tt.want {
			t.Errorf("CalculateMinimumHP(%v) = %v, want %v", tt.dungeon, got, tt.want)
		}
	}
}
