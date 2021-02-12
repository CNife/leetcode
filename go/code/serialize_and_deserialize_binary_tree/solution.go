package serialize_and_deserialize_binary_tree

import (
	"strconv"
	"strings"

	. "github.com/CNife/leetcode/go/types"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(tree *TreeNode) string {
	var result []string
	queue := []*TreeNode{tree}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, "")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}
	for len(result) > 0 && len(result[len(result)-1]) == 0 {
		result = result[:len(result)-1]
	}
	return strings.Join(result, ",")
}

func (c *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	if len(nodes[0]) == 0 {
		return nil
	}

	root := parseNode(nodes[0])
	queue := []*TreeNode{root}
	for i := 1; i < len(nodes); {
		node := queue[0]
		queue = queue[1:]

		node.Left = parseNode(nodes[i])
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nodes) {
			node.Right = parseNode(nodes[i])
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}

func parseNode(value string) *TreeNode {
	if len(value) == 0 {
		return nil
	} else {
		nodeVal, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return &TreeNode{Val: nodeVal}
	}
}
