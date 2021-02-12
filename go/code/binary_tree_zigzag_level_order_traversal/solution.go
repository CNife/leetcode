package binary_tree_zigzag_level_order_traversal

import . "github.com/CNife/leetcode/go/types"

func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var result [][]int

	for flag, size := true, 1; size > 0; flag, size = !flag, len(queue) {
		values := make([]int, size)
		if flag {
			for i := range values {
				values[i] = queue[i].Val
			}
		} else {
			for i := range values {
				values[i] = queue[size-i-1].Val
			}
		}
		result = append(result, values)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
