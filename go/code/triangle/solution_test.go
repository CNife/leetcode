package triangle

import "testing"

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		triangle [][]int
		want     int
	}{
		{
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		if got := MinimumTotal(tt.triangle); got != tt.want {
			t.Errorf("MinimumTotal(%v)=>%v, want %v", tt.triangle, got, tt.want)
		}
	}
}
