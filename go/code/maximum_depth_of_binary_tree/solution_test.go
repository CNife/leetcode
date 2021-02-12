package maximum_depth_of_binary_tree

import (
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTree(3, 9, 20, -1, -1, 15, 7),
			want: 3,
		},
	}
	for _, tt := range tests {
		if got := MaxDepth(tt.root); got != tt.want {
			t.Error(tt, got)
		}
	}
}
