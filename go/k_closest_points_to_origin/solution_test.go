package k_closest_points_to_origin

import (
	"reflect"
	"sort"
	"testing"
)

func TestKClosest(t *testing.T) {
	tests := []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		{
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
	}
	for _, tt := range tests {
		got := sortPoints(KClosest(tt.points, tt.k))
		want := sortPoints(tt.want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("KClosest(%v, %v) = %v, want %v", tt.points, tt.k, got, want)
		}
	}
}

func sortPoints(s [][]int) [][]int {
	sort.Slice(s, func(i, j int) bool {
		if s[i][0] == s[j][0] {
			return s[i][1] < s[j][1]
		}
		return s[i][0] < s[j][0]
	})
	return s
}
