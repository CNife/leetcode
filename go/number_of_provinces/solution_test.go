package number_of_provinces

import "testing"

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		connects [][]int
		want     int
	}{
		{},
		{
			connects: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			connects: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		if got := FindCircleNum(tt.connects); got != tt.want {
			t.Errorf("FindCircleNum(%v) = %v, want %v",
				tt.connects, got, tt.want)
		}
	}
}
