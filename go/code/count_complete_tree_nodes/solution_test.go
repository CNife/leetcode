package count_complete_tree_nodes

import (
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewTree(1, 2, 3, 4, 5, 6), 6},
		{nil, 0},
		{NewTree(1, 2, 3), 3},
	}
	for _, tt := range tests {
		if got := CountNodes(tt.root); got != tt.want {
			t.Errorf("CountNodes(%v) = %v, want %v",
				tt.root, got, tt.want)
		}
	}
}
