package recover_a_tree_from_preorder_traversal

import (
	"strconv"
	"strings"

	. "github.com/CNife/leetcode/go/types"
)

func RecoverFromPreorder(s string) *TreeNode {
	nodes := parse(s)
	return buildTree(nodes)
}

type node struct {
	val   int
	level int
}

func parse(s string) []node {
	var result []node
	var numBuilder strings.Builder
	level := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '-' {
			if numBuilder.Len() > 0 {
				val, _ := strconv.Atoi(numBuilder.String())
				result = append(result, node{val, level})
				numBuilder.Reset()
				level = 0
			}
			level++
		} else {
			numBuilder.WriteByte(ch)
		}
	}
	lastVal, _ := strconv.Atoi(numBuilder.String())
	result = append(result, node{lastVal, level})
	return result
}

func buildTree(nodes []node) *TreeNode {
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: nodes[0].val}
	default:
		root := &TreeNode{Val: nodes[0].val}
		i := 2
		for ; i < len(nodes) && nodes[i].level > nodes[0].level+1; i++ {
		}
		root.Left = buildTree(nodes[1:i])
		root.Right = buildTree(nodes[i:])
		return root
	}
}
