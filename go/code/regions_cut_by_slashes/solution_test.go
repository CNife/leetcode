package regions_cut_by_slashes

import "testing"

func TestRegionsBySlashes(t *testing.T) {
	tests := []struct {
		grid []string
		want int
	}{
		{
			grid: []string{
				" /",
				"/ ",
			},
			want: 2,
		},
		{
			grid: []string{
				" /",
				"  ",
			},
			want: 1,
		},
		{
			grid: []string{
				"\\/",
				"/\\",
			},
			want: 4,
		},
		{
			grid: []string{
				"/\\",
				"\\/",
			},
			want: 5,
		},
		{
			grid: []string{
				"//",
				"/ ",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		if got := RegionsBySlashes(tt.grid); got != tt.want {
			t.Errorf("RegionsBySlashes(%v) = %v, want %v",
				tt.grid, got, tt.want)
		}
	}
}
