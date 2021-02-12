package maximum_depth_of_binary_tree

import . "github.com/CNife/leetcode/go/types"

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		var leftDepth, rightDepth int
		if root.Left == nil {
			leftDepth = 0
		} else {
			leftDepth = MaxDepth(root.Left)
		}
		if root.Right == nil {
			rightDepth = 0
		} else {
			rightDepth = MaxDepth(root.Right)
		}
		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}
