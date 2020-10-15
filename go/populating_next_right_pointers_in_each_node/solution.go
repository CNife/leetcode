package populating_next_right_pointers_in_each_node

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	for level := []*Node{root}; len(level) > 0; {
		var prev *Node = nil
		var tmp []*Node
		for _, node := range level {
			if prev != nil {
				prev.Next = node
			}
			prev = node
			if node.Left != nil && node.Right != nil {
				tmp = append(tmp, node.Left, node.Right)
			}
		}
		level = tmp
	}
	return root
}
