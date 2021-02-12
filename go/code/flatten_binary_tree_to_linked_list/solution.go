package flatten_binary_tree_to_linked_list

import (
	. "github.com/CNife/leetcode/go/types"
)

func Flatten(root *TreeNode) {
	if root != nil {
		doFlatten(root)
	}
}

func doFlatten(root *TreeNode) *TreeNode {
	left, right := root.Left, root.Right
	var leftEnd, rightEnd *TreeNode
	if left != nil {
		leftEnd = doFlatten(left)
	}
	if right != nil {
		rightEnd = doFlatten(right)
	}
	if leftEnd != nil {
		root.Left = nil
		root.Right = left
		root = leftEnd
	}
	if rightEnd != nil {
		root.Left = nil
		root.Right = right
		root = rightEnd
	}
	return root
}
