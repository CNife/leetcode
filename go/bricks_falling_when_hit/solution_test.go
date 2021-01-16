package bricks_falling_when_hit

import (
	"reflect"
	"testing"
)

func TestHitBricks(t *testing.T) {
	tests := []struct {
		grid [][]int
		hits [][]int
		want []int
	}{
		{
			grid: [][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 0},
			},
			hits: [][]int{
				{1, 0},
			},
			want: []int{2},
		},
		{
			grid: [][]int{
				{1, 0, 0, 0},
				{1, 1, 0, 0},
			},
			hits: [][]int{
				{1, 1},
				{1, 0},
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		got := HitBricks(tt.grid, tt.hits)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("HitBricks(%v, %v) = %v, want %v",
				tt.grid, tt.hits, got, tt.want)
		}
	}
}
