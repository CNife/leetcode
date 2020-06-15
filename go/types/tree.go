package types

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(values ...int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}

	for i := 1; i < len(values); {
		node := queue[0]
		queue = queue[1:]

		leftValue := values[i]
		i += 1
		if leftValue >= 0 {
			node.Left = &TreeNode{Val: leftValue}
			queue = append(queue, node.Left)
		}

		if i < len(values) {
			rightValue := values[i]
			i += 1
			if rightValue >= 0 {
				node.Right = &TreeNode{Val: rightValue}
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}