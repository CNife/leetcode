package binary_tree_maximum_path_sum

import (
	. "github.com/CNife/leetcode/go/types"
	"math/bits"
)

func MaxPathSum(root *TreeNode) int {
	maxSum := (1 << bits.UintSize) / -2
	maxGain(root, &maxSum)
	return maxSum
}

func maxGain(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	leftGain := max(maxGain(node.Left, maxSum), 0)
	rightGain := max(maxGain(node.Right, maxSum), 0)
	newPathGain := node.Val + leftGain + rightGain
	*maxSum = max(*maxSum, newPathGain)
	return node.Val + max(leftGain, rightGain)
}

func max(lhs, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}
