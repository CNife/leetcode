package matrix_cells_in_distance_order

import (
	"reflect"
	"sort"
	"testing"
)

func TestAllCellsDistanceOrder(t *testing.T) {
	tests := []struct {
		r, c, r0, c0 int
		want         [][]int
	}{
		{1, 2, 0, 0, [][]int{{0, 0}, {0, 1}}},
		{2, 2, 0, 1, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
		{2, 3, 1, 2, [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}},
	}
	mapFunc := func(points [][]int, x0, y0 int) {
		sort.Slice(points, func(i, j int) bool {
			di := distance(points[i][0], points[i][1], x0, y0)
			dj := distance(points[j][0], points[j][1], x0, y0)
			if di == dj {
				if points[i][0] == points[j][0] {
					return points[i][1] < points[j][1]
				}
				return points[i][0] < points[j][0]
			}
			return di < dj
		})
	}
	for _, tt := range tests {
		got := AllCellsDistanceOrder(tt.r, tt.c, tt.r0, tt.c0)
		mapFunc(got, tt.r0, tt.c0)
		mapFunc(tt.want, tt.r0, tt.c0)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("AllCellsDistanceOrder(%v, %v, %v, %v) = %v, want %v",
				tt.r, tt.c, tt.r0, tt.c0, got, tt.want)
		}
	}
}
