package is_graph_bipartite

import "testing"

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		graph [][]int
		want  bool
	}{
		{
			graph: [][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			},
			want: true,
		},
		{
			graph: [][]int{
				{1, 2, 3},
				{0, 2},
				{0, 1, 3},
				{0, 2},
			},
			want: false,
		},
		{
			graph: [][]int{
				{},
				{2, 4, 6},
				{1, 4, 8, 9},
				{7, 8},
				{1, 2, 8, 9},
				{6, 9},
				{1, 5, 7, 8, 9},
				{3, 6, 9},
				{2, 3, 4, 6, 9},
				{2, 4, 5, 6, 7, 8},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := IsBipartite(tt.graph); got != tt.want {
			t.Errorf("graph %v, got %v, want %v", tt.graph, got, tt.want)
		}
	}
}
