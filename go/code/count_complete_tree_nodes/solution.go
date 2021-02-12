package count_complete_tree_nodes

import . "github.com/CNife/leetcode/go/types"

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	count := 0
	for {
		if queue[0].Left == nil {
			break
		}

		levelCount := len(queue)
		for i := 0; i < levelCount; i++ {
			node := queue[0]
			queue = queue[1:]
			count++
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return count + len(queue)
}
