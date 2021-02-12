package binary_tree_preorder_traversal

import . "github.com/CNife/leetcode/go/types"

func PreorderTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	result = append(result, PreorderTraversalRecursion(root.Left)...)
	result = append(result, PreorderTraversalRecursion(root.Right)...)
	return result
}

func PreorderTraversalIteration(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{root}
	for n := 1; n > 0; n = len(stack) {
		node := stack[n-1]
		stack = stack[:n-1]
		if node != nil {
			result = append(result, node.Val)
			stack = append(stack, node.Right, node.Left)
		}
	}
	return result
}
