package redundant_connection

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		edges [][]int
		want  []int
	}{
		{
			edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:  []int{2, 3},
		},
		{
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			want:  []int{1, 4},
		},
	}
	for _, tt := range tests {
		got := FindRedundantConnection(tt.edges)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FindRedundantConnection(%v) = %v, want %v",
				tt.edges, got, tt.want)
		}
	}
}
