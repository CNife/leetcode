package unique_binary_search_trees_2

import (
	. "github.com/CNife/leetcode/go/types"
)

type pair struct {
	start, end int
}

func GenerateTrees(n int) []*TreeNode {
	if n <= 0 {
		return nil
	}
	memo := make(map[pair][]*TreeNode)
	return generate(1, n, memo)
}

func generate(start, end int, memo map[pair][]*TreeNode) []*TreeNode {
	var result []*TreeNode
	if start > end {
		result = append(result, nil)
		return result
	}

	p := pair{start, end}
	if cached, exists := memo[p]; exists {
		return deepClone(cached)
	}

	for i := start; i <= end; i++ {
		leftTrees := generate(start, i-1, memo)
		rightTrees := generate(i+1, end, memo)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				node := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				result = append(result, node)
			}
		}
	}
	memo[p] = result
	return result
}

func deepClone(src []*TreeNode) []*TreeNode {
	cloned := make([]*TreeNode, 0, len(src))
	for _, tree := range src {
		cloned = append(cloned, deepCloneTree(tree))
	}
	return cloned
}

func deepCloneTree(tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	} else {
		return &TreeNode{
			Val:   tree.Val,
			Left:  deepCloneTree(tree.Left),
			Right: deepCloneTree(tree.Right),
		}
	}
}
