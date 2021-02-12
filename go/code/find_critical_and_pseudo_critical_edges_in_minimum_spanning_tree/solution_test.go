package find_critical_and_pseudo_critical_edges_in_minimum_spanning_tree

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindCriticalAndPseudoCriticalEdges(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  [][]int
	}{
		{
			n: 5,
			edges: [][]int{
				{0, 1, 1},
				{1, 2, 1},
				{2, 3, 2},
				{0, 3, 2},
				{0, 4, 3},
				{3, 4, 3},
				{1, 4, 6},
			},
			want: [][]int{
				{0, 1},
				{2, 3, 4, 5},
			},
		},
		{
			n: 4,
			edges: [][]int{
				{0, 1, 1},
				{1, 2, 1},
				{2, 3, 1},
				{0, 3, 1},
			},
			want: [][]int{
				nil,
				{0, 1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		got := FindCriticalAndPseudoCriticalEdges(tt.n, tt.edges)
		for i := 0; i < 2; i++ {
			sort.Ints(got[i])
			sort.Ints(tt.want[i])
		}
		for i := 0; i < 2; i++ {
			if !reflect.DeepEqual(got[i], tt.want[i]) {
				t.Errorf("FindCriticalAndPseudoCriticalEdges(%v, %v) = %v, want %v",
					tt.n, tt.edges, got, tt.want)
			}
		}
	}
}
