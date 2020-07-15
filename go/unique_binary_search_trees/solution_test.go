package unique_binary_search_trees

import "testing"

func TestNumTrees(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{3, 5},
	}
	for _, tt := range tests {
		if got := NumTrees(tt.n); got != tt.want {
			t.Errorf("NumTrees(%v)=>%v, want %v", tt.n, got, tt.want)
		}
	}
}
