package binary_tree_maximum_path_sum

import (
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		tree     *TreeNode
		expected int
	}{
		{
			tree: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 6,
		},
		{
			tree: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: 42,
		},
		{
			tree:     &TreeNode{Val: -3},
			expected: -3,
		},
	}

	for _, tc := range cases {
		if output := MaxPathSum(tc.tree); output != tc.expected {
			t.Errorf("%v\noutput: %v\nexpected: %v", tc.tree, output, tc.expected)
			t.Fail()
		}
	}
}
