package sort_items_by_groups_respecting_dependencies

import (
	"reflect"
	"testing"
)

func TestSortItems(t *testing.T) {
	tests := []struct {
		n, m        int
		group       []int
		beforeItems [][]int
		want        []int
	}{
		{
			n:           4,
			m:           1,
			group:       []int{-1, 0, 0, -1},
			beforeItems: [][]int{{}, {0}, {1, 3}, {2}},
			want:        nil,
		},
		{
			n:           8,
			m:           2,
			group:       []int{-1, -1, 1, 0, 0, 1, 0, -1},
			beforeItems: [][]int{nil, {6}, {5}, {6}, {3, 6}, nil, nil, nil},
			want:        []int{6, 3, 4, 1, 5, 2, 0, 7},
		},
		{
			n:           8,
			m:           2,
			group:       []int{-1, -1, 1, 0, 0, 1, 0, -1},
			beforeItems: [][]int{nil, {6}, {5}, {6}, {3}, nil, {4}, nil},
			want:        nil,
		},
	}
	for _, tt := range tests {
		got := SortItems(tt.n, tt.m, tt.group, tt.beforeItems)
		if !reflect.DeepEqual(got, tt.want) {
			t.Logf("SortItems(%v, %v, %v, %v) = %v, want %v",
				tt.n, tt.m, tt.group, tt.beforeItems, got, tt.want)
		}
	}
}
